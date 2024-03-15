package cli

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var currency string
var redirect string
var chargeType string
var amount string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new charge",
	Long:  CreateLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if amount == "" {
			log.Fatal("Please specify a price with --amount or -p when creating a charge")
		}

		chargeReq := BuildCharge(chargeType, amount, currency, redirect)
		resp, err := Client.CreateCharge(ctx, chargeReq)
		if err != nil {
			log.Fatalf("error creating charge: %s ", err)
		}
		ChargeToJSON(resp)

	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&amount, "amount", "p", "", "Set amount to charge")
	createCmd.Flags().StringVarP(&redirect, "redirect", "r", "", "URL to redirect to after charge creation")
	createCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Currency of the charge (ex: USD)")
	createCmd.Flags().StringVarP(&chargeType, "type", "t", "fixed_price", "Type of the charge: 'fixed' or 'none'")
	createCmd.MarkFlagRequired("amount")
}
