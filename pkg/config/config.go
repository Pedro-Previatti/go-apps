package config

import (
	"log"
	"text/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool                          // define if cache is enabled
	TemplateCache map[string]*template.Template // define template cache
	InfoLog       *log.Logger                   // define log
}
