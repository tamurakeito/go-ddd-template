package entity

type Account struct {
	Id       int    `json:"id"`
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type SignInRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Id     int    `json:"id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}
type SignUpRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type SignUpResponse struct {
	Id     int    `json:"id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}
