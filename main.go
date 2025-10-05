package session

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3/extractors"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/mysql"
)

var s *session.Store

type Config struct {
	Table        string        // The table to store the session data in.
	CookieSecure bool          // Whether the cookie should be HTTPS-only.
	IdleTimeout  time.Duration // How long the session cookie should last.
}

type wrapper struct {
	*mysql.Storage
}

func (self *wrapper) GetWithContext(_ context.Context, key string) ([]byte, error) {
	return self.Get(key)
}

func (self *wrapper) SetWithContext(_ context.Context, key string, value []byte, expiration time.Duration) error {
	return self.Set(key, value, expiration)
}

func (self *wrapper) DeleteWithContext(_ context.Context, key string) error {
	return self.Delete(key)
}

func (self *wrapper) ResetWithContext(_ context.Context) error {
	return self.Reset()
}

// Init initialises the session.
func Init(sessionConfig *Config, dsn string) {
	s = session.NewStore(session.Config{
		Storage: &wrapper{
			Storage: mysql.New(mysql.Config{
				ConnectionURI: dsn,
				Table:         sessionConfig.Table,
			}),
		},
		Extractor:      extractors.FromCookie("session"),
		CookieSecure:   sessionConfig.CookieSecure,
		CookieHTTPOnly: true,
		IdleTimeout:    sessionConfig.IdleTimeout,
	})
}
