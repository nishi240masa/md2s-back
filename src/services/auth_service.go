package services

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"md2s/models"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const googleCertsURL = "https://www.googleapis.com/oauth2/v3/certs"

var (
	googleClientID  = os.Getenv("GOOGLE_CLIENT_ID")
	cachedGoogleCerts models.GoogleCerts
	certsMutex      sync.Mutex
	lastCertFetch   time.Time
)

func createRSAPublicKey(e, n string) (*rsa.PublicKey, error) {
	exp, err := base64.RawURLEncoding.DecodeString(e)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %v", err)
	}

	mod, err := base64.RawURLEncoding.DecodeString(n)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %v", err)
	}

	return &rsa.PublicKey{
		N: new(big.Int).SetBytes(mod),
		E: int(new(big.Int).SetBytes(exp).Int64()),
	}, nil
}

func fetchGoogleCerts() (models.GoogleCerts, error) {
	certsMutex.Lock()
	defer certsMutex.Unlock()

	if time.Since(lastCertFetch) < 1*time.Hour {
		return cachedGoogleCerts, nil
	}

	resp, err := http.Get(googleCertsURL)
	if err != nil {
		return models.GoogleCerts{}, fmt.Errorf("failed to fetch Google certs: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.GoogleCerts{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.GoogleCerts{}, fmt.Errorf("failed to read response body: %v", err)
	}

	if len(body) == 0 {
		return models.GoogleCerts{}, errors.New("empty response body")
	}

	var certs models.GoogleCerts
	if err := json.Unmarshal(body, &certs); err != nil {
		return models.GoogleCerts{}, fmt.Errorf("failed to decode Google certs: %v", err)
	}

	cachedGoogleCerts = certs
	lastCertFetch = time.Now()
	return certs, nil
}


func VerifyGoogleToken(idToken string) (*models.GoogleIDToken, error) {
    certs, err := fetchGoogleCerts()
    if err != nil {
        return nil, err
    }

    var claims models.GoogleIDToken
    token, err := jwt.ParseWithClaims(idToken, &claims, func(token *jwt.Token) (interface{}, error) {
        kid := token.Header["kid"].(string)
        for _, key := range certs.Keys {
            if key.Kid == kid {
                return createRSAPublicKey(key.E, key.N)
            }
        }
        return nil, errors.New("invalid key ID")
    })

    if err != nil {
        return nil, fmt.Errorf("token parsing failed: %v", err)
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    if claims.Aud != googleClientID {
        return nil, errors.New("invalid audience")
    }

    return &claims, nil
}
