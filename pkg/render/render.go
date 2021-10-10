package render

import (
	"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	"log"
	"bytes"
	"github.com/kasuki2/bookings/pkg/config"
	"github.com/kasuki2/bookings/pkg/models"
)

var app *config.AppConfig

var functions = template.FuncMap{
	
}

// sets the config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the template cache from the app config 

	
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map 
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
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

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err 
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err 
			}
		}

		myCache[name] = ts 
	}

	return myCache, nil

}