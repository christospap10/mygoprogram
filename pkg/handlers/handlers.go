package handlers

import (
	"net/http"

	"github.com/christospap10/mygoprogram/pkg/config"
	"github.com/christospap10/mygoprogram/pkg/models"
	"github.com/christospap10/mygoprogram/pkg/render"
)

// Reposistory Pattern
// The repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a New repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// This is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	//Render the template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// This is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform logic
	stringMap := make(map[string]string)
	stringMap["name"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
