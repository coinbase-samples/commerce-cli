package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-cli/sdk"
	"github.com/spf13/cobra"
)

var eventId string
var all bool

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Interact with the events endpoint",
	Long:  EventsLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if eventId != "" {

			event, err := sdk.Client.ShowEvent(ctx, eventId)
			if err != nil {
				log.Fatalf("error retrieving event %s error: %s\n", eventId, err)
			}
			EventToJSON(event)
			return
		}

		if all {
			events, err := sdk.Client.ListEvents(ctx)
			if err != nil {
				log.Fatalf("error retrieving events %s", err)
			}
			EventsToJSON(events)
			return
		}
		log.Fatal("Please provide an eventId to retrieve: `events --get <eventId>`")
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.Flags().StringVarP(&eventId, "get", "g", "", "Retrieves an event by its id")
	eventsCmd.Flags().BoolVar(&all, "all", false, "Retrieve all events")

}
