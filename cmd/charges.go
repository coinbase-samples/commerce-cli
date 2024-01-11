package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-sdk-go"
	"github.com/hughescoin/commerce-cli/sdk"

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
			log.Fatalf("cannot have both a PriceValue and ChargeId")
		}

		if setPriceValue == "" && setChargeId == "" {
			log.Fatalf("Please provide either --setprice or --get flag.")
		}

		if setPriceValue != "" {
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
			fmt.Printf("charge created successfully: %v\n", resp.Data)

		}

		if setChargeId != "" {
			charge, err := sdk.Client.GetCharge(ctx, setChargeId)
			if err != nil {
				log.Fatalf("Error obtaining charge: %s - error: %s\n", setChargeId, err)
			}

			fmt.Printf("Charge details: %+v\n", charge)
		}

	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().StringVarP(&setPriceValue, "setPrice", "p", "", "Set the price for the charge")
	chargesCmd.Flags().StringVarP(&setChargeId, "get", "g", "", "Retrieve a charge by its code")
}
