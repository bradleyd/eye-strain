package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	usage = `usage:
countdown 25s
countdown 1m50s
countdown 2h45m50s
`
	tick      = time.Second
	sleepTime = 20 * time.Second
)

var (
	timer     *time.Timer
	ticker    *time.Ticker
	startDone bool
	debug     bool
	seconds   string
	quiet     bool
)

func format(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	if h < 1 {
		return fmt.Sprintf("%02d:%02d", m, s)
	}
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func start(d time.Duration) {
	timer = time.NewTimer(d)
	ticker = time.NewTicker(tick)
}

func stop() {
	timer.Stop()
	ticker.Stop()
}

func countdown(left time.Duration) {
	var exitCode int

	start(left)

loop:
	for {
		select {
		case <-ticker.C:
			left -= time.Duration(tick)
			if debug {
				fmt.Println(left)
			}
		case <-timer.C:
			break loop
		}
	}

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func notify() {
	err := beeep.Notify("eye-strain", "Please look away for 20 seconds at something 20 feet away.", "")
	if err != nil {
		panic(err)
	}
}

func beepWhenDone() {
	if !quiet {
		err := beeep.Beep(beeep.DefaultFreq, 10)
		if err != nil {
			fmt.Printf("There was an error when trying to beep %v\n", err)
		}
	}
}

func main() {

	flag.StringVar(&seconds, "s", "20m", "Interval in seconds")
	flag.BoolVar(&debug, "d", false, "Turn on debug (default false)")
	flag.BoolVar(&quiet, "q", false, "Dont beep or notify (default false)")
	flag.Parse()

	duration, err := time.ParseDuration(seconds)
	if err != nil {
		fmt.Printf("error: invalid duration: %v\n", seconds)
		os.Exit(2)
	}
	left := duration
	for {
		countdown(left)
		if debug {
			fmt.Print("Look away at 20 feet for 20 seconds...")
		}
		beepWhenDone()
		notify()
		time.Sleep(sleepTime)
		beepWhenDone()
	}
}
