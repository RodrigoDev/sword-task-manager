package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	l "github.com/RodrigoDev/sword-task-manager/internal/logging"
)

type ErrorResponse struct {
	Err string
}

type Service struct {
	repository *Repository
}

func NewUserService(repository *Repository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	ctx := r.Context()

	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not create a user: %s", err.Error()))
		err := ErrorResponse{
			Err: "password encryption failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	err = s.repository.AddUser(user)
	if err != nil {
		l.Logger(ctx).Error(fmt.Sprintf("could not create user: %s", err.Error()))
		err := ErrorResponse{
			Err: "could not create user",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := s.findOne(user.Email, user.Password)

	json.NewEncoder(w).Encode(resp)
}

func (s *Service) findOne(email, password string) map[string]interface{} {
	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return map[string]interface{}{"status": false, "message": "email address not found"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"status": false, "message": "invalid login credentials"}
	}

	claims := jwt.MapClaims{
		"user_id": *user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"exp": time.Now().Add(time.Minute * 100000).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	return map[string]interface{}{
		"status": true,
		"message": "logged in",
		"token": tokenString,
	}
}

func (s *Service) TestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	props := ctx.Value("props").(jwt.MapClaims)
	l.Logger(ctx).Info(fmt.Sprintf("%v", props["user_id"]))

	return
}
