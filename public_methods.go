package figure

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Print outputs the figure to io.Stdout.
func (fig Figure) Print() {
	for _, printRow := range fig.Slicify() {
		if fig.color != "" {
			printRow = colors[fig.color] + printRow + colors["reset"]
		}
		fmt.Println(printRow)
	}
}

// Returns the figure as a colored string.
func (fig Figure) ColorString() string {
	s := ""
	for _, printRow := range fig.Slicify() {
		if fig.color != "" {
			printRow = colors[fig.color] + printRow + colors["reset"]
		}
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

// Returns the figure as a string.
func (fig Figure) String() string {
	s := ""
	for _, printRow := range fig.Slicify() {
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

// Scrolls writes the figure and then animates it with scrolling.
// duration: total time the banner will display.
//
// stillness: duration of time the text will not move; the lower the stillness the faster
// the scroll speed.
//
// direction: can be either "right" or "left".
func (fig Figure) Scroll(duration, stillness time.Duration, direction string) {
	endTime := time.Now().Add(duration)
	fig.phrase = fig.phrase + "   "
	clearScreen()
	for time.Now().Before(endTime) {
		var shiftedPhrase string
		chars := []byte(fig.phrase)
		if strings.HasPrefix(strings.ToLower(direction), "r") {
			shiftedPhrase = string(append(chars[len(chars)-1:], chars[0:len(chars)-1]...))
		} else {
			shiftedPhrase = string(append(chars[1:], chars[0]))
		}
		fig.phrase = shiftedPhrase
		fig.Print()
		sleep(stillness)
		clearScreen()
	}
}

// Blink animates the figure with blinking.
// duration: total time the banner will display.
//
// timeOn: duration of time the text will blink on.
//
// timeOff: duration of time the text will blink off.
//
// For an even blink, set `timeOff` to -1 (same as setting `timeOff` to the value of `timeOn`).
func (fig Figure) Blink(duration, timeOn, timeOff time.Duration) {
	if timeOff < 0 {
		timeOff = timeOn
	}
	endTime := time.Now().Add(duration)
	clearScreen()
	for time.Now().Before(endTime) {
		fig.Print()
		sleep(timeOn)
		clearScreen()
		sleep(timeOff)
	}
}

// Dance writes the figure and animates it with "dancing".
//
// duration: total time the banner will display.
//
// freeze: duration of time between dance moves.
//
// NOTE: The lower the freeze the faster the dancing.
func (fig Figure) Dance(duration, freeze time.Duration) {
	endTime := time.Now().Add(duration)
	font := fig.font // TODO: change to deep copy
	font.evenLetters()
	figures := []Figure{{font: font}, {font: font}}
	clearScreen()
	for i, c := range fig.phrase {
		appenders := []string{" ", " "}
		appenders[i%2] = string(c)
		for f := range figures {
			figures[f].phrase = figures[f].phrase + appenders[f]
		}
	}
	for p := 0; time.Now().Before(endTime); p ^= 1 {
		figures[p].Print()
		figures[1-p].Print()
		sleep(freeze)
		clearScreen()
	}
}

// Write outputs the figure to an arbitrary io.Writer.
func Write(w io.Writer, fig Figure) {
	for _, printRow := range fig.Slicify() {
		fmt.Fprintf(w, "%v\n", printRow)
	}
}

// helpers
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func sleep(milliseconds time.Duration) {
	time.Sleep(milliseconds)
}
