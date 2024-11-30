package dto

type CreateUserData struct {
	Name	 string `json:"name"`
	IconURL	 string `json:"icon_url"`
	GoogleId string `json:"google_id"`

}