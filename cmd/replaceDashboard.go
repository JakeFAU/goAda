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
	"io/ioutil"
	"os"

	aio "github.com/jakefau/goAdafruit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var oldDashboardID string

// replaceDashboardCmd represents the replaceDashboard command
var replaceDashboardCmd = &cobra.Command{
	Use:   "replace",
	Short: "A brief description of your command",
	Long:  `Dashboards allow you to visualize data and control Adafruit IO connected projects from any modern web browser. Blocks such as charts, sliders, and buttons are available to help you quickly get your IoT project up and running without the need for any custom code.`,
	Run: func(cmd *cobra.Command, args []string) {
		dashFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client := aio.NewClient(viper.GetString("IOKEY"), viper.GetString("IOUSER"))
		b := aio.Dashboard{}
		err = json.Unmarshal([]byte(dashFile), &b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		board, _, err := client.Dashboard.ChangeDashboard(dashboardID, b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		jsonDash, err := json.MarshalIndent(board, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(jsonDash))

	},
}

func init() {
	dashboardsCmd.AddCommand(replaceDashboardCmd)
	replaceDashboardCmd.Flags().StringVarP(&oldDashboardID, "block", "b", "", "The ID of the dashboard to replace")
	replaceDashboardCmd.Flags().StringVarP(&fileName, "file", "f", "", "The JSON file to use")

}
