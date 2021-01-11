package main

import (
	"fmt"
	"os"

	"github.com/nbutton23/zxcvbn-go"
)

func checkPassword(userpassword string) error {

	passwordStenght := zxcvbn.PasswordStrength(userpassword, nil)
	if passwordStenght.Score < 4 {
		return fmt.Errorf("Password is too simple, please make more difficult one\nPassword score    (0-4): %d\nEstimated entropy (bit): %f\nEstimated time to crack: %s\n\n",
			passwordStenght.Score,
			passwordStenght.Entropy,
			passwordStenght.CrackTimeDisplay,
		)
	}
	return nil
}

func main() {
	var userpassword = os.Getenv("DB_PASSWORD")
	fmt.Println("checking DB_PASSWORD")
	var passworderr = checkPassword(userpassword)
	if passworderr != nil {
		fmt.Println(passworderr)
		// fail, insecure
		os.Exit(1)
	} else {
		// good
		os.Exit(0)
	}
}
