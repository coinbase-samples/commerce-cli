/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coinbase-samples/commerce-sdk-go"
	"github.com/spf13/cobra"
)

var ChargesLongDescription = `Interact with the Coinbase Commerce charges endpoint to create and view charges.

Examples:
- Create a new charge: 'commerce charges create --type "no_price"'
- Retrieve a specific charge: 'commerce charges --id <charge_id>'

Charges are presented in JSON format. All errors are returned in a standard error format.
`

var EventsLongDescription = `Interact with the Coinbase Commerce events endpoint to view event details. Use the --all flag to retrieve all events associated with your account. The --get flag requires an event_id and retrieves details of a specific event.

Examples:
- Retrieve all events: 'commerce events'
- Retrieve a specific event: 'commerce events --id <event_id>'

Events are displayed in JSON format.
`

var CreateLongDescription = `

Examples:
- create a fixed price charge: 'commerce charges create --type fixed_price --amount 5.00'
- create a donation charge: 'commerce charges create --type no_price'
- create a charge with a formatted response : 'commerce charges create --type no_price --format true'
`

func ResponseToJson(cmd *cobra.Command, response interface{}) (string, error) {
	formatBool, err := CheckFormatFlag(cmd)
	if err != nil {
		return "", err
	}
	resp, err := MarshalJSON(response, formatBool)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func CheckFormatFlag(cmd *cobra.Command) (bool, error) {
	formatFlagValue, err := cmd.Flags().GetString("format")
	if err != nil {
		return false, fmt.Errorf("cannot read format flag: %w", err)
	}
	return formatFlagValue == "true", nil
}

func MarshalJSON(data interface{}, format bool) ([]byte, error) {
	if format {
		return json.MarshalIndent(data, "", " ")
	}
	return json.Marshal(data)
}

func BuildCharge(chargeType, amount, currency, redirect string) *commerce.ChargeRequest {

	if chargeType != "fixed_price" && chargeType != "no_price" {
		log.Fatalf("cannot create charge of type %s. Please use 'fixed_price' or 'no_price'", chargeType)
	}

	charge := &commerce.ChargeRequest{
		PricingType: chargeType,
		LocalPrice: &commerce.LocalPrice{
			Amount:   amount,
			Currency: currency,
		},
		RedirectUrl: redirect,
	}
	return charge
}
