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
	"os"
	"timp/cmd/model"

	"github.com/spf13/cobra"
)

var (
	fileInfo os.FileInfo
	err      error
)

// loginCmd represents the login command
var userLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login as user",
	Long: `Login as a user, given username.
example: timp login my_username`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		if len(args) != 1 {
			fmt.Println("please specify one, and only one, username to login with")
			return
		}

		usersfile, _ := ioutil.ReadFile("cmd/resources/users.json")
		var users []model.User
		_ = json.Unmarshal([]byte(usersfile), &users)

		// Current user
		currentuserfile, _ := ioutil.ReadFile("cmd/resources/currentUser.json")

		var currentUser model.CurrentUser

		_ = json.Unmarshal([]byte(currentuserfile), &currentUser)

		if currentUser.IsLoggedIn == "true" {
			fmt.Println("already logged in as: ", currentUser.Username)
			return
		}

		var is_a_user = false
		for _, user := range users {
			if user.Username == args[0] {
				is_a_user = true
			}
		}

		if !is_a_user {
			fmt.Println("specified username is not a user. Is the username right? Is the user created?")
			return
		}

		fmt.Println("loging in as ", args[0])
		var data = model.CurrentUser{"true", args[0]}
		writefile, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile("cmd/resources/currentUser.json", writefile, 0644)
		fmt.Println("loggin succes (hopefully)")

	},
}

func init() {
	userCmd.AddCommand(userLoginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
