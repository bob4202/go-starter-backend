package user

import (
	"errors"
	"go-starter-backend/internal/config"
	jwtPkg "go-starter-backend/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
	cfg  *config.Config
}

func NewService(repo *Repository, cfg *config.Config) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}

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

func (s *Service) Register(name, email, password string) (*User, error) {
	existingUser, err := s.repo.FindByEmail(email)

	if err == nil && existingUser != nil {
		return nil, errors.New(string(UserAlreadyExists))
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, errors.New(string(ErrorHashingPassword))
	}

	newUser := &User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	}
	err = s.repo.CreateUser(newUser)
	if err != nil {
		return nil, errors.New(string(ErrorCreatingUser))
	}

	return newUser, nil
}

func (s *Service) GetUserByID(id string) (*User, error) {
	existingUser, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New(string(ErrorUserNotFound))
	}

	return existingUser, nil
}

func (s *Service) UpdateUser(id, name, email string) (*User, error) {

	user, err := s.repo.FindById(id)

	if err != nil {
		return nil, errors.New(string(ErrorUserNotFound))
	}

	user.Name = name
	user.Email = email

	err = s.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ChangePassword(id, password string) error {
	_, err := s.repo.FindById(id)

	if err != nil {
		return errors.New(string(ErrorUserNotFound))
	}

	passwordHash, err := HashPassword(password)
	if err != nil {
		return errors.New(string(ErrorHashingPassword))
	}

	err = s.repo.ChangePassword(id, passwordHash)
	if err != nil {
		return errors.New(string(ErrorChangingPassword))
	}

	return nil

}

func (s *Service) DeleteUser(id string) error {
	_, err := s.repo.FindById(id)

	if err != nil {
		return errors.New(string(ErrorUserNotFound))
	}

	err = s.repo.DeleteUser(id)

	if err != nil {
		return errors.New(string(ErrorDeletingUser))
	}

	return nil

}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return "", errors.New(string(ErrorUserNotFound))
	}

	if !CheckPassword(password, user.PasswordHash) {
		return "", errors.New(string(ErrorInvalidCredentials))
	}

	token, err := jwtPkg.GenerateToken(user.ID, s.cfg.JWTSecret)

	if err != nil {
		return "", errors.New(string(ErrorGeneratingJWT))
	}

	return token, nil

}

func (s *Service) Me(userID string) (*User, error) {
	user, err := s.repo.FindById(userID)
	if err != nil {
		return nil, errors.New(string(ErrorUserNotFound))
	}

	return user, nil
}

///////// Helper Functions ///////////

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
