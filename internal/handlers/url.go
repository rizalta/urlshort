package handlers

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/rizalta/urlshort/internal/repo"
	"github.com/rizalta/urlshort/utils"
)

type URLHandler struct {
	repo *repo.RedisRepo
}

func NewURLHandler(repo *repo.RedisRepo) *URLHandler {
	return &URLHandler{repo}
}

func (h *URLHandler) AddURL(w http.ResponseWriter, r *http.Request) {
	URL := r.FormValue("url")
	parsedURL, err := url.Parse(URL)
	if err != nil {
		http.Error(w, "error parsing url", http.StatusBadRequest)
		return
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}
	URL = parsedURL.String()
	if utils.IsDeadLink(URL) {
		http.Error(w, "Invalid Link", http.StatusBadRequest)
		return
	}
	shortStr := utils.GenerateRandomString(len(URL)%4 + 4)
	err = h.repo.AddURL(r.Context(), URL, shortStr)
	if err != nil {
		http.Error(w, "Invalid Link", http.StatusInternalServerError)
		return
	}

	baseURL := utils.CurrentBaseURL(r)
	shortURL, err := url.JoinPath(baseURL, shortStr)
	if err != nil {
		http.Error(w, "error joining url", http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles("web/templates/shorten.html"))
	data := struct {
		ShortURL string
	}{
		ShortURL: shortURL,
	}
	tmpl.Execute(w, data)
}

func (h *URLHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
}

func (h *URLHandler) RedirectToURL(w http.ResponseWriter, r *http.Request) {
	short := r.PathValue("id")
	url, err := h.repo.GetURL(r.Context(), short)
	if err != nil {
		http.Error(w, "Invalid link", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
