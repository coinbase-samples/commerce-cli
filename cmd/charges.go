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
var get bool
var create bool

var chargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Interact with the charges endpoint",
	Long:  ChargesLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if create {
			if amount == "" {
				log.Fatal("Please specify a price with --price when creating a charge")
			}
			chargeReq := BuildCharge(amount)
			resp, err := sdk.Client.CreateCharge(ctx, chargeReq)
			if err != nil {
				log.Fatalf("error creating charge: %s ", err)
			}
			ChargeToJSON(resp)
			return
		}

		if get {
			if chargeId == "" {
				log.Fatal("Please specify a charge ID with --id when retrieving a charge")
			}
			charge, err := sdk.Client.GetCharge(ctx, chargeId)
			if err != nil {
				log.Fatalf("Error obtaining charge: %s - error: %s", chargeId, err)
			}
			ChargeToJSON(charge)
			return
		}
		log.Fatal("Please specify an action --create or --get")

	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().BoolVar(&create, "create", false, "Create a new charge")
	chargesCmd.Flags().BoolVar(&get, "get", false, "Retrieve an existing charge by its id")
	chargesCmd.Flags().StringVarP(&amount, "amount", "p", "", "Set the price for a charge")
	chargesCmd.Flags().StringVar(&chargeId, "id", "", "Retrieve a charge by its id")
}
