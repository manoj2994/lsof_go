package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func check_error(e error) {

	if e != nil {
		panic(e)
	}
}

func getUid(p string) (string, error) {

	file, err := os.Open("/proc/" + p + "/status")
	if err != nil {
		// do nothing
	} else {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "Uid:") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					return parts[1], nil
				} else {
					return "", fmt.Errorf("No user ID mentioned")
				}

			}
		}
	}

	return "", fmt.Errorf("unable to find the uid")

}

func collect_info(s string, userarg *bool) {

	command, err := os.ReadFile("/proc/" + s + "/comm")
	commandline, err := os.ReadFile("/proc/" + s + "/cmdline")
	if err != nil {
		// do nothing
	} else {
		userid, _ := getUid(s)
		username, err := user.LookupId(userid)
		usernametoprint := ""
		if err == nil {
			usernametoprint = username.Username
		}
		if *userarg {
			fmt.Printf("%-6s %-30s %-10s %-10s\n", s, strings.TrimSpace(string(command)), usernametoprint, strings.TrimSpace(string(commandline)))
			//fmt.Printf("PID: '%s', CMD: '%s', UID: '%s'\n", s, string(command), userid)
		} else {
			fmt.Printf("%-6s %-30s %-10s\n", s, strings.TrimSpace(string(command)), strings.TrimSpace(string(commandline)))
		}
	}

}

func main() {

	d, err := os.ReadDir("/proc")
	userptr := flag.Bool("username", false, "gives the user who runs this program")
	check_error(err)

	flag.Parse()

	if *userptr {
		fmt.Printf("%-6s %-30s %-10s %-10s\n", "PID", "COMMAND", "USER", "COMMAND LINE")
	} else {
		fmt.Printf("%-6s %-30s %-10s\n", "PID", "COMMAND", "COMMAND LINE")
	}
	fmt.Println("-------------------------------------------------------------")
	for _, i := range d {

		j, err := strconv.Atoi(i.Name())

		if err != nil {
			continue
		}
		collect_info(strconv.Itoa(j), userptr)
	}
	//fmt.Println("lsof")
}
