package users

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}

type GetUsersResponse struct {
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
	Total int    `json:"total"`
	Users []User `json:"users"`
}

type CreateUserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
}
