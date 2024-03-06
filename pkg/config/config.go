package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool                          // define if cache is enabled
	TemplateCache map[string]*template.Template // define template cache
	InfoLog       *log.Logger                   // define log
	InProduction  bool                          // define if application is in production
	Session       *scs.SessionManager           // define session manager
}
