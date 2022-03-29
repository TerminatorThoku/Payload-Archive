package main

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func WriteLog(filename string, data []byte) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		fmt.Println(err)
		return
	}

	if _, err = f.WriteString("\n"); err != nil {
		fmt.Println(err)
		return
	}
}

// 2 * 3
// 4 * 6
// X   Y

// 0 < X < 10
// 0 < Y =< 6

func genMathCaptcha() (string, string, error) {

	x, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(10))
	if err != nil {
		return "", "", err
	}

	//Add
	y, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(14))
	if err != nil {
		return "", "", err
	}

	answer := big.NewInt(0).Add(x, y)

	return fmt.Sprint(y, " + ", x), fmt.Sprint(answer), nil

}
