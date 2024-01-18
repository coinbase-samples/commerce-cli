package cli

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
