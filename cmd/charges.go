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

var chargeId string

var chargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Interact with the charges endpoint",
	Long:  ChargesLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if chargeId != "" {
			charge, err := Client.GetCharge(ctx, chargeId)
			if err != nil {
				log.Fatalf("Error obtaining charge: %s - error: %s", chargeId, err)
			}
			resp, err := ResponseToJson(cmd, charge)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Print(resp)
		}

		log.Fatal("Please specify an action `create` or provide an --id to retrieve a charge")
	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")
	chargesCmd.Flags().StringVar(&chargeId, "id", "", "Retrieve a charge by its id")
}
