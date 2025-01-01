package model

// example
type Hello struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tag  bool   `json:"tag"`
}
type HelloWorld struct {
	Id    int   `json:"id"`
	Hello Hello `json:"hello"`
}
