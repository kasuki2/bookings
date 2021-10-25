package config

import (
	"html/template"
	"github.com/alexedwards/scs/v2"
	"log"
)

// holds the app config 
// it is a struct, so anything (any type) can be put in it
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	ErrorLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}
