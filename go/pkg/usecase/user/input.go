package user

type GetByIDInput struct {
	UserID int
}

type SignUpInput struct {
	Username string
	Email    string
	Password string
}
