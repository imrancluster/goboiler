package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/imrancluster/goboiler/conn"
	"github.com/imrancluster/goboiler/router"

	sentry "github.com/getsentry/sentry-go"
	"github.com/imrancluster/goboiler/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting server",
	Long:  "Starting mart API server",
	Run:   serve,
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		config.Init()
		if err := conn.ConnectDB(); err != nil {
			log.Fatalf("cannot connect to the database. reason: %s", err.Error())
			return err
		}
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
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func serve(cmd *cobra.Command, args []string) {
	router := router.Router()
	cfg := config.App()

	server := http.Server{
		Addr:    ":" + strconv.FormatInt(int64(cfg.Port), 10),
		Handler: router,
	}

	go func() {
		log.Printf("server started on %d", cfg.Port)
		server.ListenAndServe()
	}()

	// creating channel

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	log.Println("Signal ", <-ch, " received")

	ctx, errCtx := context.WithTimeout(context.Background(), time.Second*5)
	if errCtx != nil {
		log.Fatalf("failed to serve: %v", errCtx)
	}
	server.Shutdown(ctx)
}
