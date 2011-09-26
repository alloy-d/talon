package main

import (
	"github.com/alloy-d/go140"
	"flag"
	"fmt"
	"strings"
	"time"
	"os"
)

func padLine(line string, length int) string {
	if len(line) < length {
		return line + strings.Repeat(" ", length - len(line))
	}
	return line
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Please specify a username!\n")
		os.Exit(-1)
	}

	api := new(go140.API)
	api.Root = "https://api.twitter.com"

	user, err := api.User(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	line := "Screen name: " + user.ScreenName
	line = padLine(line, 40)
	line += "In real life: " + user.Name
	fmt.Println(line)

	line = "Location: " + user.Location
	line = padLine(line, 40)
	line += "Homepage: " + user.URL
	fmt.Println(line)

	if user.Status == nil {
		os.Exit(0)
	}

	tweetTime, err := time.Parse(time.RubyDate, user.Status.Date)
	if err != nil {
		fmt.Println("Error parsing tweet time: ", err)
		os.Exit(1)
	}
	tweetTime = time.SecondsToLocalTime(tweetTime.Seconds())
	line = "Last tweet " + tweetTime.Format("3:04 PM, Jan _2, 2006")
	fmt.Println(line)
	fmt.Println(user.Status.Text)
}
