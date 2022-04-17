package domain

type User struct {
	ID        uint64 `json:"user id"`
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Email     string `json:"email"`
}
