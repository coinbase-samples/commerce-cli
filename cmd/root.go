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
	"log"
	"net/http"
	"os"

	"github.com/coinbase-samples/commerce-sdk-go"
	"github.com/spf13/cobra"
)

var Client *commerce.Client

const (
	COMMERCE_API_KEY = "COMMERCE_API_KEY"
)

var rootCmd = &cobra.Command{
	Use:   "commerce",
	Short: "Commerce is a CLI for managing e-commerce operations.",
	Long:  `A fast and flexible CLI built with Cobra for managing e-commerce operations.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return InitClient()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(eventsCmd)
	rootCmd.AddCommand(chargesCmd)
}

func InitClient() error {
	creds, err := commerce.ReadEnvCredentials(COMMERCE_API_KEY)
	if err != nil {
		log.Fatalf("error reading environmental variable: %s", err)
	}

	Client = commerce.NewClient(creds, http.Client{})
	return nil
}
