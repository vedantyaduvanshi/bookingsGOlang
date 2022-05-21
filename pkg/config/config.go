package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig HOLDS THE APPLICATION CONFIG
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
