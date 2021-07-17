package handlers

import (
	"net/http"

	"github.com/selvamshan/bookingwebapp/pkg/config"
	"github.com/selvamshan/bookingwebapp/pkg/models"
	"github.com/selvamshan/bookingwebapp/pkg/render"
)

// Repo ....
var Repo *Repository

// Repository ...
type Repository struct {
	App *config.AppConfig
}

// NewRepo ...
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler set the repository for handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home handler function for home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About handler function for about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] ="Hello again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
