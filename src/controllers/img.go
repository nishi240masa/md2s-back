package controllers

import (
	"encoding/base64"
	"md2s/dto"
	"md2s/utils/r2"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {

	err := r2.CreateS3Client()
	if err != nil {
		panic("Failed to initialize S3 client: " + err.Error())
	}

	var img dto.UploadImgData

	// JSONを構造体にバインド
	if err := c.ShouldBindJSON(&img); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DTOでエラー"})
		return
	}

	// base64形式の画像データをデコード
	fileData, err := base64.StdEncoding.DecodeString(img.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "画像データのデコードに失敗!"})
		return
	}

	fileName := img.Name + ".png"

	// R2にアップロード
	err = r2.UploadFile(fileName, strings.NewReader(string(fileData)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url, err := r2.GenerateURL(fileName) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, url)
}