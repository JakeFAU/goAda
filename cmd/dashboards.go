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
	"fmt"

	"github.com/spf13/cobra"
)

// dashboardsCmd represents the dashboards command
var dashboardsCmd = &cobra.Command{
	Use:   "dashboards",
	Short: "Manipulate dashboards",
	Long:  `Dashboards allow you to visualize data and control Adafruit IO connected projects from any modern web browser. Blocks such as charts, sliders, and buttons are available to help you quickly get your IoT project up and running without the need for any custom code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dashboards called")
	},
}

func init() {
	rootCmd.AddCommand(dashboardsCmd)
	dashboardsCmd.PersistentFlags().StringVarP(&dashboardID, "dashboard", "d", "", "The dashboard the blocks are on")
}
