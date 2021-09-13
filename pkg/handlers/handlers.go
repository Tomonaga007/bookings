package handlers

import (
	"github.com/Tomonaga007/bookings/pkg/config"
	"github.com/Tomonaga007/bookings/pkg/models"
	"github.com/Tomonaga007/bookings/pkg/render"
	"log"
	"net/http"
)


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository  {
	return &Repository{App: a}
}

func NewHandlres(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	remoteIP := r.RemoteAddr
	log.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request)  {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Apartment_1(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "apartment-1.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Apartment_2(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "apartment-2.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}