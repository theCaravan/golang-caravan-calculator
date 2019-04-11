package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// When called, this prints the Sezzle logo on the screen
func sezzlelogo() {

	fmt.Printf("\n                           WNW                            \n" +
		"                          NK0N                            \n" +
		"                        N000KXW                           \n" +
		"                       N000000XW                          \n" +
		"                      WK00000O0XW                         \n" +
		"                WWW   NKO0000000XW                        \n" +
		"               W0OX   N0000000000KN                       \n" +
		"              NOxx0W  WK00000000000XW                     \n" +
		"             NOxxxkX   N0000000000000XW                   \n" +
		"            W0xxxxxON  WX0000000000000KNW                 \n" +
		"            XkxxxxxxOX   NK0000000000000KNW               \n" +
		"     WK0N   XkxxxxxxxkKW  WX00000000000000KN              \n" +
		"    NOodK   XkxxxxxxxxxOXW  WNX0000000000000XN            \n" +
		"   NkoookN  W0xxxxxxxxxxk0N    NX0000000000000XW          \n" +
		"  WOoooooOW  NOxxxxxxxxxxxkKW    NX000000000000KN         \n" +
		"  XdooooooOW  N0xxxxxxxxxxxxO0KW   NK000000000000XW       \n" +
		"  KdoooooookX  WKkxxxxxxxxxxxxxOXW   NK00000000000KN      \n" +
		"  Kdooooooood0W  NKkxxxxxxxxxxxxxOXW   NK00000000000XW    \n" +
		"  XxooooooooookXW  N0kxxxxxxxxxxxxk0N   WNK0000000000XW   \n" +
		"  W0doooooooooodON   N0kxxxxxxxxxxxxkKW   WX0000000000N   \n" +
		"   W0doooooooooood0N   N0kxxxxxxxxxxxxOXW   NK00000000KW  \n" +
		"    WKxoooooooooooox0W  WN0kxxxxxxxxxxxkKW   WX000000OKW  \n" +
		"      N0dooooooooooooxKW  WX0kxxxxxxxxxxxON   WX000000KW  \n" +
		"        NOdooooooooooookXW  WWXOxxxxxxxxxxON   WX0000KN   \n" +
		"          NOdooooooooooodOX    WKkxxxxxxxxxON   WX00KN    \n" +
		"            XOdoooooooooood0W    N0xxxxxxxxxKW   NKXW     \n" +
		"              XOdoooooooooook0X   WKkxxxxxxxON    W       \n" +
		"               WXOdooooooooooodKW   NOxxxxxxOW            \n" +
		"                 WXkoooooooooood0W   NOxxxxkX             \n" +
		"                   WKxooooooooooo0W   XkxxkXW             \n" +
		"                     N0dooooooooodK    KkON               \n" +
		"                      WXkoooooooooOW   WXW                \n" +
		"                       NOooooooookW                       \n" +
		"                        W0ooooooo0                        \n" +
		"                         WOoooooOW                        \n" +
		"                          Nkood0W                         \n" +
		"                           XOOX                           \n" +
		"                            W                             \n")
}

// Graceful Shutdown
func end() {

	sezzlelogo()
	fmt.Printf("\nGoodbye! :-)\n")
}

func main() {

	// This allows me to press Ctrl-C (aka SIGINT, which can lead to SIGTERM) and exit gracefully rather than abruptly
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		end()
		os.Exit(1)
	}()

	var version string = "Build 0.22"
	fmt.Printf("\nCaravan Calculator by The Caravan -- " + version)
	fmt.Printf("\nBuilt from Python - Ported to Golang\n")

	var prevInput = ""    // This holds the previous value of whatever I type into the system after "CARAVAN-CALCULATOR >>>"
	var inputString = ""  // This holds the current  value of whatever I type into the system after "CARAVAN-CALCULATOR >>>"
	var storedAnswer = "" // This holds the answer to the previous entry. Allows use of "Ans" that is present in many calculators

	// This was a capy-paste from my Python code. Supposedly shows what operations this calculator supports
	var supportedOps = [4]string{"FLOAT + FLOAT", "FLOAT - FLOAT", "FLOAT * FLOAT", "FLOAT / FLOAT"}

	for true {

		fmt.Printf("\nCARAVAN-CALCULATOR >>> ")
		fmt.Scanf("%s", &inputString)

		// If this is empty, we will use the previous entry. If I typed "2 + 2", and then "", then the calculator will run "2 + 2" again
		if inputString == "" {
			inputString = prevInput

			// This allows me to explain what this calculator can do
		} else if inputString == "help" || inputString == "?" {

			var i int = 0

			fmt.Printf("\nList of supported operations. Note that FLOAT and INT are placeholders for the accepted value:\n")

			for i < len(supportedOps) {
				var op string = supportedOps[i]
				fmt.Printf("- \n" + op)
				i++
			}
			continue

			// Easter Egg
		} else if inputString == "sezzle" {

			sezzlelogo()
			fmt.Printf("\nIt's Sezzle Time!\n")

			// If this has "ans", does not contain an "=" (meaning it would be a logic check - aka if x = y --> 0 or 1), or if the string is not just "ans"
		} else if strings.Contains(inputString, "ans") && !strings.Contains(inputString, "=") && inputString != "ans" {

			fmt.Printf("Hi I am here - Contains ans")

			// Type "exit" to exit
		} else if inputString == "exit" {
			break
		}

	}

	end()
}
