package create_user

import "time"

type Request struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	NickName  string `json:"nick_name"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

type Response struct {
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}
