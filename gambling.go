package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Red string = "\x1b[31m"
const Green string = "\x1b[32m"
const Yellow string = "\x1b[33m"
const Underline string = "\x1b[4m"
const Bold string = "\x1b[1m"
const End_EscapeCode string = "\x1b[0m"

func Rolling() {
	fmt.Print("Rolling.")
	time.Sleep(1 * time.Second)
	fmt.Print("\rRolling..")
	time.Sleep(1 * time.Second)
	fmt.Print("\rRolling...")
	time.Sleep(1 * time.Second)
}

func main() {
	var mode string
	var Chips int = 10
	var Wager int

	for {

		if Chips > 0 {

			fmt.Printf("\n\n\n               %s%sMODE SELECT%s\n"+
				"|-=====================================-|\n"+
				"|[%sEasy%s]: 1 |[%sMedium%s]: 2 |[%s%sDifficult%s]: 3 |\n"+
				"|-=====================================-|\n"+
				"Chips %d | exit: 4\n"+
				"\n"+
				"---> ",
				Bold, Underline, End_EscapeCode, Green, End_EscapeCode, Yellow, End_EscapeCode, Red, Bold, End_EscapeCode, Chips)
			fmt.Scanln(&mode)

			if mode == "1" {
				Roll := rand.Intn(101)

				fmt.Print("Welcome to the roulette table\nadd your guess: ")
				_, _ = fmt.Scanln(new(string))

				fmt.Print("add your wager: ")
				fmt.Scanln(&Wager)

				if Wager <= Chips {
					if Roll >= 50 {
						Rolling()

						fmt.Printf("\ryou won %d chips", Wager*2)
						Chips += Wager * 2
						time.Sleep(2 * time.Second)

					} else {
						Rolling()

						fmt.Printf("\rSorry you lost %d chips", Wager)
						time.Sleep(2 * time.Second)
						Chips -= Wager

					}
				} else {
					fmt.Print("Sorry buddy you don't have enough chips")
					time.Sleep(2 * time.Second)

				}

			} else if mode == "2" {
				Roll := rand.Intn(101)

				fmt.Print("Welcome to the roulette table\nadd your guess: ")
				_, _ = fmt.Scanln(new(string))

				fmt.Print("add your wager: ")
				fmt.Scanln(&Wager)

				if Wager <= Chips {
					if Roll >= 60 {
						Rolling()

						fmt.Printf("\ryou won %d chips", Wager*3)
						Chips += Wager * 3
						time.Sleep(2 * time.Second)

					} else {
						Rolling()

						fmt.Printf("\rSorry you lost %d chips", Wager*2)
						time.Sleep(2 * time.Second)
						Chips -= Wager * 2

					}
				} else {
					fmt.Print("Sorry buddy you don't have enough chips")
					time.Sleep(2 * time.Second)

				}

			} else if mode == "3" {
				Roll := rand.Intn(101)

				fmt.Print("Welcome to the roulette table\nadd your guess: ")
				_, _ = fmt.Scanln(new(string))

				fmt.Print("add your wager: ")
				fmt.Scanln(&Wager)

				if Wager <= Chips {
					if Roll >= 80 {
						Rolling()

						fmt.Printf("\ryou won %d chips", Wager*4)
						Chips += Wager * 4
						time.Sleep(2 * time.Second)

					} else {
						Rolling()

						fmt.Printf("\rSorry you lost %d chips", Wager*3)
						time.Sleep(2 * time.Second)
						Chips -= Wager * 3

					}
				} else {
					fmt.Print("Sorry buddy you don't have enough chips")
					time.Sleep(2 * time.Second)

				}

			} else if mode == "4" {
				fmt.Print("Goodbye")
				time.Sleep(2 * time.Second)
				break
			} else {
				fmt.Print("err: Command not recognized")
				time.Sleep(2 * time.Second)
			}

		} else {
			fmt.Print("\n\nSorry")
			time.Sleep(1 * time.Second)
			fmt.Print("\rSorry you")
			time.Sleep(1 * time.Second)
			fmt.Print("\rSorry you lost")
			time.Sleep(1 * time.Second)
			fmt.Print("\rSorry you lost the")
			time.Sleep(1 * time.Second)
			fmt.Print("\rSorry you lost the game")
			time.Sleep(2 * time.Second)
			Chips = 10

		}
	}
}
