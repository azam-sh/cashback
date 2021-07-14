package main

import "fmt"

const fullPercent = 100
const cashbackPercent = 15
const limit = 2_000

func calculateCashback(amount int) {
	cashback := 0
	if amount >= 5_000 {
		cashback = amount * cashbackPercent / fullPercent
	} else if amount < 5_000 {
		cashback = 0
	}
	if cashback > limit {
		cashback = limit
	}
	fmt.Println(cashback)
}

func main() {
	calculateCashback(12_000)
}

// HOMETASK
