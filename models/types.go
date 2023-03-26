package models

type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"fullname"`
	Username   string `json:"username"`
	NPassword  string `json:"NPassword"`
	CNPassword string `json:"CNPassword"`
}
