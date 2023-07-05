package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	ErrorLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}
