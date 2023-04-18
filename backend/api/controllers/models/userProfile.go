package models

type Userprofile struct {
	Name string
	Email string
	Age int
	Password string  `json:"password"`
}

