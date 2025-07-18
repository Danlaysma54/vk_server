package handler

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"vk_server/internal/repository/ad"
)

const DEFAULT_ORDER = "DESC"
const DEFAULT_LIMIT = 20
const DEFAULT_OFFSET = 0
const DEFAULT_MIN_PRICE = 1
const DEFAULT_MAX_PRICE = 10000000

type QueryParams struct {
	DateSort  string `form:"date_sort"`
	PriceSort string `form:"price_sort"`
	MinPrice  int    `form:"min_price"`
	MaxPrice  int    `form:"max_price"`
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
}
type Request struct {
	AdName      string `json:"adName" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"max=1000"`
	ImageUrl    string `json:"imageUrl"`
	Price       int    `json:"price" validate:"required,min=1,max=10000000"`
	AuthorId    string `json:"authorId" validate:"required"`
}
type Response struct {
	AdID        string `json:"adId"`
	AdName      string `json:"adName"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Price       int    `json:"price"`
	AuthorId    string `json:"authorId"`
}
type errResponse struct {
	Message string `json:"message"`
}
type ControllerAd struct {
	repo ad.IRepoAd
}

var validate = validator.New()

func (r *Request) Validate() error {
	return validate.Struct(r)
}
func (r *Request) ValidateImage() error {
	u, err := url.ParseRequestURI(r.ImageUrl)
	if err != nil {
		return fmt.Errorf("invalid image url")
	}
	ext := strings.ToLower(filepath.Ext(u.Path))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	if !allowedExts[ext] {
		return fmt.Errorf("invalid image format")
	}
	resp, err := http.Head(r.ImageUrl)
	if err != nil {
		return fmt.Errorf("failed to check image %v", err)
	}
	defer resp.Body.Close()
	maxSize := int64(5 << 20)
	if resp.ContentLength >= maxSize {
		return fmt.Errorf("image is too large,max size is 5MB")
	}
	return nil
}
func sanitizeSortOrder(s, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	s = strings.ToUpper(s)
	if s != "ASC" && s != "DESC" {
		return defaultValue
	}
	return s
}
func parseInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return val
}
func NewControllerAd(repo ad.IRepoAd) *ControllerAd {
	return &ControllerAd{repo: repo}
}

func (s *ControllerAd) New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.ad.adController.New"
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err := req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, errResponse{Message: err.Error()})
			return
		}
		if err := req.ValidateImage(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, errResponse{Message: err.Error()})
			return
		}
		if err != nil {
			log.Fatal("error decoding body", err)
		}
		adId := s.repo.SaveAd(req.AdName, req.Description, req.ImageUrl, req.Price, req.AuthorId)
		adEnt := s.repo.GetAd(adId)
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, Response{
			AdID:        adEnt.AdId,
			AdName:      adEnt.Name,
			Description: adEnt.Description,
			ImageUrl:    adEnt.ImageUrl,
			Price:       adEnt.Price,
			AuthorId:    adEnt.AuthorId,
		})
		return
	}
}

func (s *ControllerAd) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		params := &QueryParams{
			DateSort:  sanitizeSortOrder(q.Get("date_sort"), DEFAULT_ORDER),
			PriceSort: sanitizeSortOrder(q.Get("date_sort"), DEFAULT_ORDER),
			MinPrice:  parseInt(q.Get("min_price"), DEFAULT_MIN_PRICE),
			MaxPrice:  parseInt(q.Get("max_price"), DEFAULT_MAX_PRICE),
			Limit:     parseInt(q.Get("limit"), DEFAULT_LIMIT),
			Offset:    parseInt(q.Get("offset"), DEFAULT_OFFSET),
		}

		res := s.repo.GetAllAds(params.DateSort, params.PriceSort, params.MinPrice, params.MaxPrice, params.Limit, params.Offset)
		render.JSON(w, r, res)
		return
	}
}
