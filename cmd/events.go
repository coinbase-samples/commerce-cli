package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-cli/sdk"
	"github.com/spf13/cobra"
)

var setEventId string
var getAll bool

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Interact with the events endpoint",
	Long:  EventsLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if setEventId != "" {

			event, err := sdk.Client.ShowEvent(ctx, setEventId)
			if err != nil {
				log.Fatalf("error retrieving event %s error: %s\n", setEventId, err)
			}
			eventJson, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatalf("error marshalling events into JSON \n all events: %v ", event)
			}

			fmt.Printf("event %s found \n", string(eventJson))
			return
		}

		if getAll {
			allEvents, err := sdk.Client.ListEvents(ctx)
			if err != nil {
				log.Fatalf("error retrieving events %s", err)
			}
			eventsJson, err := json.MarshalIndent(allEvents, "", " ")
			if err != nil {
				log.Fatalf("error marshalling events into JSON \n all events: %v ", allEvents)
			}

			fmt.Println(string(eventsJson))
			return
		}

		log.Fatal("Please provide an eventId to retrieve: `events --get <eventId>`")

	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.Flags().StringVarP(&setEventId, "get", "g", "", "Retrieves an event by its id")
	eventsCmd.Flags().BoolVar(&getAll, "all", false, "Retrieve all events")

}
