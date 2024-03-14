package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-cli/sdk"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Retrieve Commerce events data",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		events, err := sdk.Client.ListEvents(ctx)
		if err != nil {
			log.Fatalf("error retrieving events: %s", err)
		}
		EventsToJSON(events)
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

		event, err := sdk.Client.ShowEvent(ctx, eventId)
		if err != nil {
			log.Fatalf("error retrieving event %s: %s", eventId, err)
		}
		EventToJSON(event)
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.AddCommand(getEventCmd)

	getEventCmd.Flags().String("id", "", "ID of the event to retrieve")
	getEventCmd.MarkFlagRequired("id")
}
