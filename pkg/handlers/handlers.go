package handlers

import (
	"net/http"

	"github.com/Denis-Andrei/bookings/pkg/config"
	"github.com/Denis-Andrei/bookings/pkg/models"
	"github.com/Denis-Andrei/bookings/pkg/render"
)


var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.Appconfig
}

// NewRepo creates a new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlres
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
// m is a receiver and now all the handlers have access to the Repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	// m.App.Session.Cookie = something - we added Session into the config so we can have access to it here

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
