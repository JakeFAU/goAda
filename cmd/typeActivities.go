/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	aio "github.com/jakefau/goAdafruit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var activityType string

// allCmd represents the all command
var typeActivitiesCmd = &cobra.Command{
	Use:   "type",
	Short: "Return All Activities of a certain type",
	Long: `Activities are Adafruit IO's list of actions you've taken to create, update, or delete objects in your Adafruit IO account. We store the last 1000 actions taken for:

	Dashboards
	Blocks
	Feeds
	Groups
	Triggers`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aio.NewClient(viper.GetString("IOKEY"), viper.GetString("IOUSER"))
		results, _, err := client.Activities.ActivitiesByType(&activityType)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, a := range results {
			fmt.Println(a)
		}

	},
}

func init() {
	activitiesCmd.AddCommand(typeActivitiesCmd)
	typeActivitiesCmd.Flags().StringVarP(&activityType, "type", "t", "", "Type of Activity")
}
