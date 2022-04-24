package main

import (
	"bufio"
	"fmt"
	"os"
	"sbg-api/models"
	"strconv"
)

func main() {
	var odd bool
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number you want to bet on: ")
	num, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(num)

	fmt.Print("Enter the colour you want to bet on: ")
	c, _ := reader.ReadString('\n')

	fmt.Print("Enter the parity you want to bet on: ")
	p, _ := reader.ReadString('\n')
	if p == "1" {
		odd = true
	}
	b := models.Bets{
		Number: n,
		Colour: c,
		IsOdd:  odd,
	}

	result := models.NewResult()
	wins := getWins(b, result)
	fmt.Println(wins.NumberWin + wins.ColourWin + wins.ParityWin)
}

func getWins(b models.Bets, r models.Result) models.Winning {
	var w models.Winning
	if b.Number == r.Number {
		w.NumberWin = "Your number won!"
	}
	if b.Colour == r.Colour {
		w.ColourWin = "Your colour won!"
	}
	if b.IsOdd == r.IsOdd {
		w.ParityWin = "Your parity won!"
	}
	return w
}
