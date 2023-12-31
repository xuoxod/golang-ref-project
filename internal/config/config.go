package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

// App Configuration
type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Session      *scs.SessionManager
	DBConnection string
}
