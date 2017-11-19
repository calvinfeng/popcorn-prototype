package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"goyak/model"
	"net/http"
)

func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, password := r.PostFormValue("email"), r.PostFormValue("password")
		if len(email) == 0 || len(password) == 0 {
			renderError(w, "Please Provide email and password for authentication", 400)
			return
		}

		var user model.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			renderError(w, "Email is not recognized", 400)
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(password)); err != nil {
			renderError(w, "Incorrect password", 400)
		} else {
			res := SessionResponse{
				Email:        user.Email,
				SessionToken: user.SessionToken,
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
}
