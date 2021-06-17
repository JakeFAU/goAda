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
	"github.com/spf13/cobra"
)

var dashboardID string

// blocksCmd represents the blocks command
var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "Get the blocks for a dashboard",
	Long:  `Blocks are objects which can be placed on an Adafruit IO Dasboard for a user. Blocks IO range from input blocks (sliders and buttons) to output blocks (such as maps or other visual displays).`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("blocks called")
	// },
}

func init() {
	rootCmd.AddCommand(blocksCmd)
	blocksCmd.PersistentFlags().StringVarP(&dashboardID, "dashboard", "d", "", "The dashboard the blocks are on")

}
