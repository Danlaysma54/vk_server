package ad

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
	"vk_server/internal/repository"
)

type Request struct {
	AdName      string `json:"adName"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	AuthorId    string `json:"authorId"`
}
type Response struct {
	AdID        string `json:"adId"`
	AdName      string `json:"adName"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	AuthorId    string `json:"authorId"`
}
type ControllerAd struct {
	repo repository.IRepoAd
}

func NewControllerAd(repo repository.IRepoAd) *ControllerAd {
	return &ControllerAd{repo: repo}
}

func (s *ControllerAd) New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.ad.adController.New"
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Fatal("error decoding body", err)
		}
		s.repo.SaveAd(req.AdName, req.Description, req.Price, req.AuthorId)
		render.JSON(w, r, Response{
			AdID:        "bruh",
			AdName:      req.AdName,
			Description: req.Description,
			Price:       req.Price,
			AuthorId:    req.AuthorId,
		})
	}
}
