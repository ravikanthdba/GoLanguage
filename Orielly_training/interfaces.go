package main

import (
	"fmt"
)

// Problem - we define a TapePlayer struct and TapeRecorder struct. We have 2 methods, which have a Play and stop method. We define a function, which takes TapePlayer  and list of songs as an input. The function plays all the songs and then stops once completed.

type TapePlayer struct {
	Batteries string
}

type TapeRecorder struct {
	Recorder string
}

// the below methods take only TapeRecorder as input, and hence the above code runs

func (t TapePlayer) start(song string) {
	fmt.Println("Playing ", song)
}

func (t TapePlayer) stop() {
	fmt.Println("TapePlayer Stopped")
}

func (t TapeRecorder) start(song string) {
	fmt.Println("Playing ", song)
}

func (t TapeRecorder) stop() {
	fmt.Println("Tape Recorder stopped")
}


func playList(device TapePlayer, songs []string) {
	for _, song := range songs {
		device.start(song)
	}
	device.stop()
}


type Player interface {
	start(song string)
	stop()
}

func PlayListInterface(device Player, songs []string) {
	for _, song := range songs {
		device.start(song)
	}
	device.stop()
}

func main() {
	var tapeplayer TapePlayer
	var mixSongs = []string{"Song # 1", "Song # 2", "Song # 3", "Song # 4"}
	playList(tapeplayer, mixSongs)

	// The below code will not run, as playList takes type TapePlayer and not TapeRecorder. Similarly with "Start" and "stop" as well
	// var taperecorder TapeRecorder
	// playList(taperecorder, mixSongs)

	var tapePlayer TapePlayer
	PlayListInterface(tapePlayer, mixSongs)

	var tapeRecorder TapeRecorder
	PlayListInterface(tapeRecorder, mixSongs)	

	// Exercise: https://play.golang.org/p/6OWXZ-686kO
}