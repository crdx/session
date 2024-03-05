package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
)

var s *session.Store

type Config struct {
	Table        string        // The table to store the session data in.
	CookieSecure bool          // Whether the cookie should be HTTPS-only.
	Expiration   time.Duration // How long the session cookie should last.
}

// Init initialises the session.
func Init(sessionConfig *Config, dsn string) {
	s = session.New(session.Config{
		Storage: mysql.New(mysql.Config{
			ConnectionURI: dsn,
			Table:         sessionConfig.Table,
		}),
		KeyLookup:      "cookie:session",
		CookieSecure:   sessionConfig.CookieSecure,
		CookieHTTPOnly: true,
		Expiration:     sessionConfig.Expiration,
	})
}
