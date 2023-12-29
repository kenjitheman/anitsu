package main

import (
	"flag"
	"fmt"
	"os"

	jikan "github.com/darenliang/jikan-go"
	"github.com/kenjitheman/anitsu/api"
)

func main() {
	filter := parseFlags()

	data, err := api.GetData(filter)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(data)
}

func parseFlags() jikan.ScheduleFilter {
	var day string

	flag.StringVar(&day, "d", "monday", "Day of the week")
	flag.Parse()

	return jikan.ScheduleFilter(day)
}
