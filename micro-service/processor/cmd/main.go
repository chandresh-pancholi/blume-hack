package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"processor/api"
)

func main() {

	var command = &cobra.Command{
		Use:   "processor",
		Short: "Platform for event",
		Long:  "msg is a platform which can be used as Async communication, Data pipeline and auditing",
		RunE: func(cmd *cobra.Command, args []string) error {
			server := api.NewServer()

			server.Run()

			return nil
		},
	}

	if error := command.Execute(); error != nil {
		log.Fatalf("error in executing main command. %v", error)
		os.Exit(1)
	}
}
