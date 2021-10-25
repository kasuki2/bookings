package render

import (
	"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	
	"bytes"
	"github.com/kasuki2/bookings/internal/config"
	"github.com/kasuki2/bookings/internal/models"
	"github.com/justinas/nosurf"
	"errors"
)

var app *config.AppConfig

var pathToTemplates = "./templates"

var functions = template.FuncMap{
	
}

// sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}
// td template data 
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the template cache from the app config 

	
	t, ok := tc[tmpl]
	if !ok {
		//log.Fatal("Could not get template from template cache")
		return errors.New("can't get template from cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map 
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err 
	}

	for _, page := range pages {
		// filepath returns the full path 
		name := filepath.Base(page)
		
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err 
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err 
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err 
			}
		}

		myCache[name] = ts 
	}

	return myCache, nil

}