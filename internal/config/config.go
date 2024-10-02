package config

import "github.com/gorilla/sessions"

type AppConfig struct {
	Session *sessions.CookieStore
}
