package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

func main() {

	public_key := 10.0
	private_key_a := 14.0
	private_key_b := 16.0

	key := public_key * private_key_a * private_key_b
	sft := 0.0
	sfd := 0.0
	fmt.Println(key)

	zero := 33.0
	md := 94.0

	file, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	fmt.Print(text)
	fmt.Print(string(text))

	text2 := make([]byte, len(text))

	for i, chr := range text {
		//fmt.Println(sft)
		switch {
		case chr > 32:
			if chr == byte(zero+sft) {
				text2[i] = 32
			} else {
				sft = math.Mod(key, md)

				if sft == 0 {
					key++
					sft = math.Mod(key, md)
				}
				sfd = float64(chr) - (sft + 0)
				if sfd > (zero - 1) {
					text2[i] = byte(sfd)
				} else {
					text2[i] = byte(md + sfd)
				}
				//text2[i] = byte(math.Mod((float64(chr)-zero)+sft, md) + ,ro)
				key += math.Floor(key / sft)
			}

		case chr == 32:
			text2[i] = byte(zero + sft)
		default:
			text2[i] = chr
		}
		//fmt.Println(chr)

	}

	fmt.Print(text2)
	fmt.Print(string(text2))
}
