package cmd

import (
	"github.com/joho/godotenv"
	"github.com/magiskboy/katalog/interfaces"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "katalog",
	Short: "Katalog commandline",
}

// Execute run main cli
func Execute() error {
	var envFilename string

	var restCmd = &cobra.Command{
		Use: "rest",
		Run: func(cmd *cobra.Command, args []string) {

			if len(envFilename) == 0 {
				envFilename = ".env"
			}

			if _, err := os.Stat(envFilename); !os.IsNotExist(err) {
				godotenv.Load(envFilename)
				log.Printf("env %s is loaded", envFilename)
			}

			log.Println("HTTP application is starting...")
			app := interfaces.GetHTTPBinding()
			app.Run()
		},
	}
	restCmd.Flags().StringVarP(&envFilename, "env", "e", "", "Environment variable filename *.env")
	rootCmd.AddCommand(restCmd)

	var grpcCmd = &cobra.Command{
		Use: "grpc",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Start grpc server")
		},
	}
	rootCmd.AddCommand(grpcCmd)

	return rootCmd.Execute()
}
