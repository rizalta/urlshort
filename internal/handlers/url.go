package handlers

import (
	"net/http"
	"net/url"

	"github.com/rizalta/urlshort/internal/repo"
	"github.com/rizalta/urlshort/utils"
	"github.com/rizalta/urlshort/web/components"
	"github.com/rizalta/urlshort/web/pages"
)

type URLHandler struct {
	repo *repo.RedisRepo
}

func NewURLHandler(repo *repo.RedisRepo) *URLHandler {
	return &URLHandler{repo}
}

func (h *URLHandler) AddURL(w http.ResponseWriter, r *http.Request) {
	URL := r.FormValue("url")
	URL = utils.EnforceHTTP(URL)
	if !utils.IsValidURL(URL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	shortStr := utils.GenerateRandomString(len(URL)%4 + 4)
	err := h.repo.AddURL(r.Context(), URL, shortStr)
	if err != nil {
		http.Error(w, "Invalid Link", http.StatusInternalServerError)
		return
	}

	baseURL := utils.CurrentBaseURL(r)
	shortURL, err := url.JoinPath(baseURL, shortStr)
	if err != nil {
		http.Error(w, "error joining url", http.StatusInternalServerError)
	}

	err = components.ShortenURL(shortURL).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Rendering error", http.StatusInternalServerError)
	}
}

func (h *URLHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	err := pages.Home().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *URLHandler) RedirectToURL(w http.ResponseWriter, r *http.Request) {
	short := r.PathValue("id")
	url, err := h.repo.GetURL(r.Context(), short)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
