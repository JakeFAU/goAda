/*
Copyright © 2021 Jacob Bourne

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	aio "github.com/jakefau/goAdafruit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// allCmd represents the all command
var allActivitiesCmd = &cobra.Command{
	Use:   "all",
	Short: "Return All Activities",
	Long: `Activities are Adafruit IO's list of actions you've taken to create, update, or delete objects in your Adafruit IO account. We store the last 1000 actions taken for:

	Dashboards
	Blocks
	Feeds
	Groups
	Triggers`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aio.NewClient(viper.GetString("IOKEY"), viper.GetString("IOUSER"))
		results, _, err := client.Activities.All()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, a := range results {
			jsonResult, err := json.MarshalIndent(a, "", "  ")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(string(jsonResult))
		}

	},
}

func init() {
	activitiesCmd.AddCommand(allActivitiesCmd)
}
