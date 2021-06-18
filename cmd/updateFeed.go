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

// updateFeedCmd represents the updateFeed command
var updateFeedCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a feed",
	Long: `Feeds are the core of the Adafruit IO system. The feed holds metadata about the data you push to Adafruit IO. This includes settings for whether the data is public or private, what license the stored data falls under, and a general description of the data. The feed also contains the sensor data values that get pushed to Adafruit IO from your device.

	You will need to create one feed for each unique source of data you send to Adafruit IO.
	
	You can create, read, update, or delete feeds. Every CREATE, UPDATE, or DELETE action on a feed record counts against your rate limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aio.NewClient(viper.GetString("IOKEY"), viper.GetString("IOUSER"))
		feedFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		f := aio.Feed{}
		err = json.Unmarshal([]byte(feedFile), &f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newFeed, _, err := client.Feed.Update(feedID, &f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		jsonResult, err := json.MarshalIndent(newFeed, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(jsonResult))

	},
}

func init() {
	feedsCmd.AddCommand(updateFeedCmd)
	updateFeedCmd.Flags().StringVarP(&feedID, "feed", "i", "", "The feed id")
	updateFeedCmd.Flags().StringVarP(&fileName, "file", "f", "", "The JSON file to use")
}
