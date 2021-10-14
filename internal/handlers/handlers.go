package handlers

import (
	"github.com/kasuki2/bookings/internal/render"
	"github.com/kasuki2/bookings/internal/config"
	"github.com/kasuki2/bookings/internal/models"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)



// Repo teh repository used by the handlers 
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
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)



	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again from bookings."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
// renders the generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}
// PostAvailability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request){
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}
type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}
// PostAvailabilityJSON handlers request for availability and send json response
func (m *Repository) PostAvailabilityJSON(w http.ResponseWriter, r *http.Request){
	resp := jsonResponse{
		OK: true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

