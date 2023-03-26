package useraccountmanagement

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	NPassword  string `json:"NPassword"`
	CNPassword string `json:"CNPassword"`
}
