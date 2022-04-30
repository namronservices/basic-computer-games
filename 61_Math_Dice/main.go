package main

import (
	"fmt"
	"math/rand"
	"time"
)

func drawDice(num int) {
	fmt.Printf("#######\n")

	var s1, s2, s3 string
	switch num {
	case 3:
		s1, s2, s3 = "o", " ", " "
	case 4, 5, 6:
		s1, s2, s3 = "o", " ", "o"
	default:
		s1, s2, s3 = " ", " ", " "
	}
	fmt.Printf("# %v%v%v #\n", s1, s2, s3)

	switch num {
	case 1, 3, 5:
		s1, s2, s3 = " ", "o", " "
	case 2, 6:
		s1, s2, s3 = "o", " ", "o"
	default:
		s1, s2, s3 = " ", " ", " "

	}
	fmt.Printf("# %v%v%v #\n", s1, s2, s3)

	switch num {
	case 3:
		s1, s2, s3 = " ", " ", "o"
	case 4, 5, 6:
		s1, s2, s3 = "o", " ", "o"
	default:
		s1, s2, s3 = " ", " ", " "

	}
	fmt.Printf("# %v%v%v #\n", s1, s2, s3)

	fmt.Printf("#######\n\n")
}

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var i int
	for {
		a := r1.Intn(5) + 1
		b := r1.Intn(5) + 1
		drawDice(a)
		drawDice(b)
		for {
			_, err := fmt.Scanf("%d", &i)
			if err == nil {
				break
			}
		}

		if i == a+b {
			fmt.Printf("Good answer!\n")
		} else {
			fmt.Printf("Unfortunately it is a wrong answer! The good answer is %d\n", a+b)
		}
	}

}
