package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber()int{
	rand.Seed(time.Now().Unix())
	return rand.Int()%11;
	}

	func getUserInput()int{
		fmt.Printf("Please enter your guess")
		var input int;

		_,err :=fmt.Scan(&input)
		if err != nil{
			fmt.Println("Failed to parse your input")
		}else{
			fmt.Println("you guessed ",input)
		}
		return input
	}



func main(){
	//create a seceret number
    secret := getRandomNumber()
	//Get user Input
	for matching :=false; !matching ;{
	guess := getUserInput()
	// fmt.Println(secret,guess)
	//make comparisons (secret vs guess)
	matching := isMatching(secret,guess)
	fmt.Println(matching)
	// }
	}
}
func isMatching(secret,guess int)bool{
	if guess > secret{
		fmt.Println("your guess is too big")
		return false
	}else if guess<secret{
	 fmt.Println("your guess is too short")
	 return false
	}else{
	fmt.Println("You have guessed it correctly")
	return true
	}
}
	//Make comparisons (seceret vs guess)
// 	matching :=isMatching(seceret,guess)
// 	return
// }
