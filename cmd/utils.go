package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coinbase-samples/commerce-sdk-go"
)

var ChargesLongDescription = `Interact with the Coinbase Commerce charges endpoint to create and view charges. Use --setPrice to create a new charge with a specified USD amount. The --get flag requires a charge_id to retrieve a specific charge.

Examples:
- Create a new charge: 'commerce charges --setPrice 1.5'
- Retrieve a specific charge: 'commerce charges --get <charge_id>'

Charges are presented in JSON format. All errors are returned in a standard error format.
`

var EventsLongDescription = `Interact with the Coinbase Commerce events endpoint to view event details. Use the --all flag to retrieve all events associated with your account. The --get flag requires an event_id and retrieves details of a specific event.

Examples:
- Retrieve all events: 'commerce events --all'
- Retrieve a specific event: 'commerce events --get <event_id>'

Events are displayed in JSON format.
`

func EventToJSON(e *commerce.SingleEvent) {
	eventJson, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		log.Fatalf("error marshalling events into JSON \n all events: %v ", e)
	}
	fmt.Printf("event %s found \n", string(eventJson))

}

func EventsToJSON(e *commerce.EventResponse) {
	eventsJson, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		log.Fatalf("error marshalling events into JSON \n all events: %v ", e)
	}
	fmt.Println(string(eventsJson))
}

func ChargeToJSON(c *commerce.ChargeResponse) {
	chargeJson, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Fatalf("error marshalling response into JSON: %s \n. charge response: %v", err, c)
	}
	fmt.Printf("charge: %s", string(chargeJson))

}

func BuildCharge(v string) *commerce.ChargeRequest {
	c := &commerce.ChargeRequest{
		PricingType: "fixed_price",
		LocalPrice: &commerce.LocalPrice{
			Amount:   v,
			Currency: "USD",
		},
	}
	return c
}
