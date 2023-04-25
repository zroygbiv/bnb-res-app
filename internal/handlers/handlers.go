package handlers

import (
	"fmt"
	"net/http"

	"github.com/zroygbiv/bnb-res-app/internal/config"
	"github.com/zroygbiv/bnb-res-app/internal/forms"
	"github.com/zroygbiv/bnb-res-app/internal/models"
	"github.com/zroygbiv/bnb-res-app/internal/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}


// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


// Kitchen renders kitchen page; displays form
func (m *Repository) Kitchen(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "kitchen.page.tmpl", &models.TemplateData{})
}


// Common renders common page; displays form
func (m *Repository) Common(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"common.page.tmpl", &models.TemplateData{})
}


// Bedroom renders bedroom page; displays form
func (m *Repository) Bedroom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "bedroom.page.tmpl", &models.TemplateData{})
}


// Availability renders search availability page; displays form
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}
// PostAvailability renders search availability page; displays form
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}


// Reservation renders reservation page; displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}
// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

}


// Reservation renders reservation page; displays form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}