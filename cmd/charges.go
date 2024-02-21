package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-cli/sdk"

	"github.com/spf13/cobra"
)

var amount string
var chargeId string

var chargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Interact with the charges endpoint",
	Long:  ChargesLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if sdk.Client == nil {
			log.Fatalf("client not initialized")
		}

		if amount != "" && chargeId != "" {
			log.Fatalf("cannot have both a price and charge id")
		} else if amount != "" {
			chargeReq := BuildCharge(amount)
			resp, err := sdk.Client.CreateCharge(ctx, chargeReq)
			if err != nil {
				log.Fatalf("error creating charge: %s ", err)
			}
			ChargeToJSON(resp)
		} else if chargeId != "" {
			charge, err := sdk.Client.GetCharge(ctx, chargeId)
			if err != nil {
				log.Fatalf("Error obtaining charge: %s - error: %s\n", chargeId, err)
			}
			ChargeToJSON(charge)
		} else {
			log.Fatalf("Please provide either --setPrice (-p) or --get (-g) flag.")
		}
	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().StringVarP(&amount, "setPrice", "p", "", "Set the price for a charge")
	chargesCmd.Flags().StringVarP(&chargeId, "get", "g", "", "Retrieve a charge by its code")
}
