package model

// example
type Hello struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tag  bool   `json:"tag"`
}
type HelloWorld struct {
	Id    int     `json:"id"`
	Hello []Hello `json:"hello"`
}

// authentication
type Account struct {
	Id       int    `json:"id"`
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type AuthRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
type AuthResponse struct {
	Account Account `json:"account"`
	Token   string  `json:"token"`
}
