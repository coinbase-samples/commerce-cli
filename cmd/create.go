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
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var currency string
var redirect string
var chargeType string
var amount string
var format string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new charge",
	Long:  CreateLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if chargeType == "fixed_price" && amount == "" {
			log.Fatal("Please specify a price with --amount or -p when creating a charge")
		}

		chargeReq := BuildCharge(chargeType, amount, currency, redirect)
		charge, err := Client.CreateCharge(ctx, chargeReq)
		if err != nil {
			log.Fatalf("error creating charge: %s ", err)
		}
		response, err := ResponseToJson(cmd, charge)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(response)

	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&amount, "amount", "p", "", "Set amount to charge")
	createCmd.Flags().StringVarP(&redirect, "redirect", "r", "", "URL to redirect to after charge creation")
	createCmd.Flags().StringVarP(&currency, "currency", "c", "USD", "Currency of the charge (ex: USD)")
	createCmd.Flags().StringVarP(&chargeType, "type", "t", "fixed_price", "Type of the charge: 'fixed_price' or 'no_price'")
	createCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")
}
