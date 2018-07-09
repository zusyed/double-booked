package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zusyed/doublebooked/calendar"
)

func main() {
	filename := flag.String("input", "input.txt", "the name of the file to read the input from")
	flag.Parse()

	inFile, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var events []calendar.Event
	for scanner.Scan() {
		line := scanner.Text()
		interval := strings.Fields(line)
		start, err := strconv.ParseInt(interval[0], 10, 64)
		if err != nil {
			fmt.Printf("error converting interval's start time to int: %s", err)
			os.Exit(1)
		}

		end, err := strconv.ParseInt(interval[1], 10, 64)
		if err != nil {
			fmt.Printf("error converting interval's end time to int: %s", err)
			os.Exit(1)
		}

		event := calendar.Event{
			StartTime: start,
			EndTime:   end,
		}
		events = append(events, event)
	}

	overlappingEvents, err := calendar.GetOverlappingEvents(events)
	if err != nil {
		fmt.Printf("error occured getting overlapping events: %s", err)
		os.Exit(1)
	}

	fmt.Printf("%v", overlappingEvents)
}
