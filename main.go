package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error")
	}
	banner := "standard.txt"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}
	bannerfile := banner
	txt, err := loadBanner(bannerfile)
	if err != nil {
		fmt.Println("invalid")
	}
	last := os.Args[1]
	lasst := generateArt(last, txt)
	fmt.Println(lasst)
}
