package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/commerce-cli/sdk"
	"github.com/spf13/cobra"
)

var chargeId string

var chargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Interact with the charges endpoint",
	Long:  ChargesLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if chargeId != "" {
			charge, err := sdk.Client.GetCharge(ctx, chargeId)
			if err != nil {
				log.Fatalf("Error obtaining charge: %s - error: %s", chargeId, err)
			}
			ChargeToJSON(charge)
			return
		}

		log.Fatal("Please specify an action `create` or provide an --id to retrieve a charge")

	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().StringVar(&chargeId, "id", "", "Retrieve a charge by its id")
}
