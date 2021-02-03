package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/jefersonfdasilv/codepix-go/application/grpc"
	"github.com/jefersonfdasilv/codepix-go/infrastructure/db"
)

var portNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Server port")
}
