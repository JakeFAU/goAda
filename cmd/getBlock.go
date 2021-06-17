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
	"os"

	aio "github.com/jakefau/goAdafruit"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var blockID string

// getCmd represents the get command
var getBlockCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Block",
	Long:  `Blocks are objects which can be placed on an Adafruit IO Dasboard for a user. Blocks IO range from input blocks (sliders and buttons) to output blocks (such as maps or other visual displays).`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aio.NewClient(viper.GetString("IOKEY"), viper.GetString("IOUSER"))
		block, _, err := client.Blocks.GetBlock(dashboardID, blockID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(block)

	},
}

func init() {
	blocksCmd.AddCommand(getBlockCmd)
	getBlockCmd.Flags().StringVarP(&blockID, "id", "b", "", "The block ID")

}
