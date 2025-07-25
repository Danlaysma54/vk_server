package handler

import (
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"log"
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
type TokenResponse struct {
	Token string `json:"token"`
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
	log.Println("hello")
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err})
		return
	}
	if h.repo.IsExist(user.Username) {
		w.WriteHeader(http.StatusConflict)
		render.JSON(w, r, map[string]interface{}{
			"message": "User already exists"})
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err})
	}
	h.repo.InsertUser(user.Username, hashedPassword)
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"message": "User created"})
	return

}
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userReq model.UserRequest
	if err := render.DecodeJSON(r.Body, &userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{
			"message": err})
		return
	}
	user := h.repo.GetUserByUsername(userReq.Username)
	if !utils.CheckPasswordHash(userReq.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, "Wrong password")
	}
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(h.tokenExpireDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{
			"message": err})
		return
	}
	render.JSON(w, r, TokenResponse{Token: tokenString})
}
