package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-learning/grpc/grpc-hello-world/server"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover error: %v", err)
			}
		}()

		server.Serve()
	},
}

func init() {
	fmt.Println("scmd s")
	serverCmd.Flags().StringVarP(&server.ServerPort, "port", "p", "50052", "server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertName, "cert-name", "", "grpc server name", "server's hostname")
	serverCmd.Flags().StringVarP(&server.SwaggerDir, "swagger-dir", "", "proto", "path to the directory which contains swagger definitions")
	fmt.Println("scmd e")
	rootCmd.AddCommand(serverCmd)
}
