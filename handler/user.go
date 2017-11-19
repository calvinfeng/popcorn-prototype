package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"goyak/model"
	"net/http"
)

func NewUserListHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []model.User

		if err := db.Find(&users).Error; err != nil {
			renderError(w, err.Error(), 400)
			return
		}

		res := []UserResponse{}
		for _, user := range users {
			userResponse := UserResponse{
				Name:  user.Name,
				Email: user.Email,
			}

			res = append(res, userResponse)
		}

		if bytes, err := json.Marshal(res); err != nil {
			renderError(w, err.Error(), 500)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

func NewUserRetrieveHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var responseText string

		params := mux.Vars(r)
		if params["id"] != "" {
			responseText = fmt.Sprintf("User %v detail is here!", params["id"])
		} else {
			responseText = "User id is not provided."
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseText))
	}
}

func NewUserCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			logrus.Error(err)
		}

		name, email, password := r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password")
		if len(email) == 0 || len(password) == 0 {
			renderError(w, "Please Provide email and password for user sign up", 400)
			return
		}

		hashBytes, hashErr := bcrypt.GenerateFromPassword([]byte(password), 10)
		if hashErr != nil {
			renderError(w, hashErr.Error(), 500)
			return
		}

		token, tokenErr := generateBase64Token(64)
		if tokenErr != nil {
			renderError(w, tokenErr.Error(), 500)
		}

		newUser := model.User{
			Name:           name,
			Email:          email,
			SessionToken:   token,
			PasswordDigest: hashBytes,
		}

		if err := db.Create(&newUser).Error; err != nil {
			renderError(w, err.Error(), 400)
			return
		}

		res := UserResponse{
			Name:         newUser.Name,
			Email:        newUser.Email,
			SessionToken: newUser.SessionToken,
		}

		if bytes, err := json.Marshal(res); err != nil {
			renderError(w, err.Error(), 500)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}
