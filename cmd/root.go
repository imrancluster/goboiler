package cmd

import (
	"log"
	"os"

	"github.com/imrancluster/goboiler/config"
	"github.com/imrancluster/goboiler/conn"
	"github.com/spf13/cobra"
)

// RootCmd ..
var RootCmd = &cobra.Command{
	Use:   "goboiler",
	Short: "Goboiler API Server",
	Long:  "Goboiler API Server",
}

// Execute : Execute function to run the server
func Execute() {
	config.Init()
	conn.ConnectDB()
	conn.ConnectRedis()
	conn.ConnectSentry()

	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("cannot run server. reason: %s", err.Error())
		os.Exit(1)
	}
}
