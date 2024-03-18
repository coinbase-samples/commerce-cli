package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Retrieve Commerce events data",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		events, err := Client.ListEvents(ctx)
		if err != nil {
			log.Fatalf("error retrieving events: %s", err)
		}
		response, err := ResponseToJson(cmd, events)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(response)
	},
}

var getEventCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a specific event by its ID",
	Run: func(cmd *cobra.Command, args []string) {
		eventId, _ := cmd.Flags().GetString("id")

		if eventId == "" {
			log.Fatal("Please specify an event ID with --id")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		event, err := Client.ShowEvent(ctx, eventId)
		if err != nil {
			log.Fatalf("error retrieving event %s: %s", eventId, err)
		}
		response, err := ResponseToJson(cmd, event)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(response)
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")
	eventsCmd.AddCommand(getEventCmd)

	getEventCmd.Flags().String("id", "", "ID of the event to retrieve")
	getEventCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")
	getEventCmd.MarkFlagRequired("id")
}
