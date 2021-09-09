package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
)

var users []string = []string{"guest", "jsef5"};
var uSelected int = 0
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var PS1 string = ">>"
func cmdLoop() {
	for {
		fmt.Printf("\nTTE@%s\n%s ", users[uSelected], PS1)
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdStr = strings.TrimSuffix(cmdStr, "\n")
		arrCmdStr := strings.Fields(cmdStr)
		switch arrCmdStr[0] {
			case "exit": os.Exit(0)
			case "setPS1": PS1 = arrCmdStr[1]
			case "usradd": users = append(users, strings.Join(arrCmdStr[1:], " "))
			case "usrlogin": {
				for i, s := range users {
					fmt.Printf("User %d: %s\n", i, s)
				}
				fmt.Print("> ")
				var uNum int
				_, err := fmt.Scanf("%d", &uNum)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Failed to read an integer")
					break;
				}
				if uNum > len(users) || uNum < 0 {
					fmt.Fprintln(os.Stderr, "User does not exist")
					break;
				}
				uSelected = uNum
			}
			default: {
				cmd := exec.Command(arrCmdStr[0], arrCmdStr[1:]...)
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout
				err = cmd.Run()
			}
		}
		if err != nil {
			fmt.Fprintln(os.Stderr , err)
		}
	}
}
func main() {
	fmt.Println("Welcome to TTE!")
	cmdLoop()
}