package main

// Basic User Struct
type User struct {
	Name     string `json:"name" xml:"name" form:"name" query:"name"`
	Username string `json:"username" xml:"username" form:"username" query:"username"`
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}
