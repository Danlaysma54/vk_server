package handler

import (
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"vk_server/internal/model"
	"vk_server/internal/repository/user"
	"vk_server/internal/utils"
)

type AuthHandler struct {
	repo                user.IRepoUser
	jwtSecret           []byte
	tokenExpireDuration time.Duration
}

func NewAuthHandler(repo user.IRepoUser, jwtSecret []byte) *AuthHandler {
	return &AuthHandler{
		repo:                repo,
		jwtSecret:           jwtSecret,
		tokenExpireDuration: time.Hour * 24,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserRequest
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}
	if h.repo.IsExist(user.Username) {
		w.WriteHeader(http.StatusConflict)
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err)
	}
	h.repo.InsertUser(user.Username, hashedPassword)
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, "User created")
	return

}
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userReq model.UserRequest
	if err := render.DecodeJSON(r.Body, &userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}
	user := h.repo.GetUserByUsername(userReq.Username)
	if !utils.CheckPasswordHash(userReq.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, "Wrong password")
	}
	claims := jwt.MapClaims{
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(h.tokenExpireDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, tokenString)
}
