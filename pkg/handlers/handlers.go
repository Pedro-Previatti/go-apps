package handlers

import (
	"net/http"

	"github.com/Pedro-Previatti/go-apps/pkg/config"
	"github.com/Pedro-Previatti/go-apps/pkg/models"
	"github.com/Pedro-Previatti/go-apps/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
