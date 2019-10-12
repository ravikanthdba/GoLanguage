package main

import (
	"github.com/ravikanthdba/GoLanguage/Orielly_training/greeting"
	"github.com/ravikanthdba/GoLanguage/Orielly_training/greeting/hindi"
	"github.com/ravikanthdba/GoLanguage/Orielly_training/greeting/telugu"
)

func main() {
	greeting.Hello("Ravikanth")
	greeting.Hi()
    hindi.Namaste("Ravikanth")
	telugu.Namaskaram("Ravikanth")
}