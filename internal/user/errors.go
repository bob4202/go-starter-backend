package user

type UserErrors string

var (
	UserAlreadyExists       UserErrors = "user already exists"
	ErrorHashingPassword    UserErrors = "error generating hash"
	ErrorCreatingUser       UserErrors = "error creating user"
	ErrorUserNotFound       UserErrors = "error user not found"
	ErrorChangingPassword   UserErrors = "error changing password"
	ErrorDeletingUser       UserErrors = "error deleting user"
	ErrorInvalidCredentials UserErrors = "invalid credentials"
	ErrorGeneratingJWT      UserErrors = "error generating jwt token"
)
