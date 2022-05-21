package handlers

import (
	"net/http"

	"github.com/vedantyaduvanshi/bookings/pkg/config"
	"github.com/vedantyaduvanshi/bookings/pkg/models"
	"github.com/vedantyaduvanshi/bookings/pkg/render"
)

//The Repository used by the HAndlers
var Repo *Repository

//Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo Creates the new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewhANDLERS sets the repositoryt for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the about Home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About is the about Page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//PERFORM SOME logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	//Send data to the Template

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
