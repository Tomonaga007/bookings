package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Tomonaga007/bookings/internal/config"
	"github.com/Tomonaga007/bookings/internal/models"
	"github.com/Tomonaga007/bookings/internal/render"
	"log"
	"net/http"
)


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	remoteIP := r.RemoteAddr
	log.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request)  {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r,"about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST form"))
}

func (m *Repository) Apartment_1(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"apartment-1.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Apartment_2(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "apartment-2.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "search-availability.page.tmpl", &models.TemplateData{})
}

type responseJson struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	var resp  = responseJson{true, "message"}
	out, err := json.MarshalIndent(resp,"","     ")
	if err != nil{
		log.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(out)
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start, end := r.Form.Get("start_date"), r.Form.Get("end_date")
	w.Write([]byte(fmt.Sprintf("Start date is %s, end date is %s", start, end)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "contact.page.tmpl", &models.TemplateData{})
}