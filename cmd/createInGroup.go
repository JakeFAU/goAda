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

// createInGroupCmd represents the createInGroup command
var createInGroupCmd = &cobra.Command{
	Use:   "createInGroup",
	Short: "create a feed in a group",
	Long: `Groups are a set of Feeds. They're used for publishing and reading to multiple feeds at a time. For example, if you are building a weather station, you would add feeds for humidity and temperature to a new weatherstation group.

	You can create, read, update, or delete group records. Every CREATE, UPDATE, or DELETE action on a group record counts against your rate limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		feedFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client := aio.NewClient(viper.GetString("IOKEY"), viper.GetString("IOUSER"))
		f := aio.Feed{}
		err = json.Unmarshal([]byte(feedFile), &f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newFeed, _, err := client.Group.CreateFeedInGroup(groupID, &f)
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
	groupsCmd.AddCommand(createInGroupCmd)
	createInGroupCmd.Flags().StringVarP(&fileName, "file", "f", "", "The JSON file to use")
	createInGroupCmd.Flags().StringVarP(&groupID, "group", "g", "", "The group id")
}
