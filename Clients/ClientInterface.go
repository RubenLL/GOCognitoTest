package clients

type AuthClient interface {
	SignUp(name string, email string, password string) (error, bool)
	SignIn(userID string, password string) (bool, string, error)
	ConfirmSignUp(email string, code string) (error, bool)
}
