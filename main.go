package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
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
	start(left)

loop:
	for {
		select {
		case <-ticker.C:
			left -= time.Duration(tick)
			systray.SetTitle(left.String())
			if debug {
				fmt.Println(left)
			}
		case <-timer.C:
			break loop
		}
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

func onReady() {
	go func() {
		duration, err := time.ParseDuration(seconds)
		if err != nil {
			fmt.Printf("error: invalid duration: %v\n", seconds)
			os.Exit(2)
		}
		left := duration
		systray.SetIcon(decodeIcon())

		for {
			countdown(left)
			notify()
			systray.SetTitle("Look away")
			time.Sleep(sleepTime)
			beepWhenDone()
		}
	}()
	systray.SetTooltip("EyeStrain")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}
func onExit() {
	fmt.Println("Thats all folks!")
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

func main() {

	flag.StringVar(&seconds, "s", "1m", "Interval in seconds")
	flag.BoolVar(&debug, "d", false, "Turn on debug (default false)")
	flag.BoolVar(&quiet, "q", false, "Dont beep or notify (default false)")
	flag.Parse()
	systray.Run(onReady, onExit)
}
