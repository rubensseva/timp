/*
Copyright © 2019 Ruben Svanåsbakken Sevaldson <r.sevaldson@gmail.com>

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
	"timp/cmd/model"

	"github.com/spf13/cobra"
)

// newUserCmd represents the newUser command
var userNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Add new user",
	Long: `Adds a new user.
Users are more like a profile, there are no passwords.
With a user, you may track score etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newUser called")
		if len(args) != 1 {
			fmt.Println("please specify one, and only one, username to create")
			return
		}

		usersfile, _ := ioutil.ReadFile("cmd/resources/users.json")
		var users []model.User
		_ = json.Unmarshal([]byte(usersfile), &users)

		var is_a_user = false
		for _, user := range users {
			if user.Username == args[0] {
				is_a_user = true
			}
		}

		if is_a_user {
			fmt.Println("specified username " + args[0] + " is already a user")
			return
		}

		fmt.Println("creating user ", args[0])
		var newUser = model.User{args[0], 0}
		users = append(users, newUser)
		writefile, _ := json.MarshalIndent(users, "", " ")
		_ = ioutil.WriteFile("cmd/resources/users.json", writefile, 0644)
		fmt.Println("create user success (hopefully)")
	},
}

func init() {
	userCmd.AddCommand(userNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
