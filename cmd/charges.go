package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-cli/sdk"
	"github.com/coinbase-samples/commerce-sdk-go"

	"github.com/spf13/cobra"
)

var setPriceValue string
var setChargeId string

var chargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Interact with the charges endpoint",
	Long:  `Interact with the Coinbase Commerce charges endpoint to create and view charges.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if sdk.Client == nil {
			log.Fatalf("client not initialized")
		}

		if setPriceValue != "" && setChargeId != "" {
			log.Fatalf("cannot have both a price and charge id")
		} else if setPriceValue != "" {
			chargeReq := commerce.ChargeRequest{
				PricingType: "fixed_price",
				LocalPrice: &commerce.LocalPrice{
					Amount:   setPriceValue,
					Currency: "USD",
				},
			}
			resp, err := sdk.Client.CreateCharge(ctx, &chargeReq)
			if err != nil {
				log.Fatalf("error creating charge: %s ", err)
			}
			jsonResponse, err := json.MarshalIndent(resp, "", " ")
			if err != nil {
				log.Fatalf("error marshalling response into JSON: %s \n. charge response: %v", err, resp.Data)
			}
			fmt.Printf("charge created: \n %s", string(jsonResponse))

		} else if setChargeId != "" {
			charge, err := sdk.Client.GetCharge(ctx, setChargeId)
			if err != nil {
				log.Fatalf("Error obtaining charge: %s - error: %s\n", setChargeId, err)
			}
			chargeJson, err := json.MarshalIndent(charge, "", " ")
			if err != nil {
				log.Fatalf("error marshalling response into JSON: %s \n. charge response: %v", err, charge)
			}
			fmt.Printf("charge %s retreived: \n %s", setChargeId, string(chargeJson))
		} else {
			log.Fatalf("Please provide either --setPrice (-p) or --get (-g) flag.")
		}
	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().StringVarP(&setPriceValue, "setPrice", "p", "", "Set the price for a charge")
	chargesCmd.Flags().StringVarP(&setChargeId, "get", "g", "", "Retrieve a charge by its code")
}
