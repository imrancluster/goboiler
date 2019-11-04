package conn

import (
	"log"
	"os"

	sentry "github.com/getsentry/sentry-go"
	"github.com/imrancluster/goboiler/config"
)

// ConnectSentry : function for sentry connection
func ConnectSentry() error {
	cfg := config.App()
	if dsn := cfg.Sentry; dsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         dsn,
			DebugWriter: os.Stderr,
		}); err != nil {
			log.Fatalf("cannot connect to sentry. reason: %s", err.Error())
			return err
		}
	}
	return nil
}
