package main

import (
	"fmt"
	"math"
)

const (
	jan = 31
	mar = 31
	apr = 30
	may = 31
	jun = 30
	jul = 31
	aug = 31
	sep = 30
	oct = 31
	nov = 30
	dec = 31
)

func main() {
	fmt.Println("Pythagorean Triple Dates: ")
	var feb int
	for i := 1900; i <= 3000; i++ {
		if determineLeapYear(i) == true {
			feb = 29
		} else {
			feb = 28
		}

		for j := 1; j <= jan; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Jan %v, %v\n", j, i)
			}
		}
		for j := 1; j <= feb; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Feb %v, %v\n", j, i)
			}
		}
		for j := 1; j <= mar; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Mar %v, %v\n", j, i)
			}
		}
		for j := 1; j <= apr; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Apr %v, %v\n", j, i)
			}
		}
		for j := 1; j <= may; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Mau %v, %v\n", j, i)
			}
		}
		for j := 1; j <= jun; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Jun %v, %v\n", j, i)
			}
		}
		for j := 1; j <= jul; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Jul %v, %v\n", j, i)
			}
		}
		for j := 1; j <= aug; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Aug %v, %v\n", j, i)
			}
		}
		for j := 1; j <= sep; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Sep %v, %v\n", j, i)
			}
		}
		for j := 1; j <= oct; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Oct %v, %v\n", j, i)
			}
		}
		for j := 1; j <= nov; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Nov %v, %v\n", j, i)
			}
		}
		for j := 1; j <= dec; j++ {
			if findPythagTriple(j, 1, i) == true {
				fmt.Printf("Dec %v, %v\n", j, i)
			}
		}
	}

}

func findPythagTriple(day int, month int, year int) bool {
	lastTwoDig := year % 1e2
	if math.Pow(float64(day), 2)+math.Pow(float64(month), 2) == math.Pow(float64(lastTwoDig), 2) {
		return true
	} else {
		return false
	}
}

func determineLeapYear(year int) bool {
	if year%4 == 0 && year%400 == 0 {
		return true
	} else if year%100 == 0 && year%400 != 0 {
		return false
	} else {
		return false
	}
}
