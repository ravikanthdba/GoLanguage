package main

import (
	"fmt"
)

type TapePlayer struct {
	Battery string
}

type TapeRecorder struct {
	Microphone string
}

func (t TapePlayer) start(song string) {
	fmt.Println("Playing song: ", song)
}

func (t TapePlayer) stop() {
	fmt.Println("Stopped")
}

func (t TapeRecorder) start(song string) {
	fmt.Println("Playing song: ", song)
}

func (t TapeRecorder) stop() {
	fmt.Println("Stopped")
}

func (t TapeRecorder) record() {
	fmt.Println("Recording..")
}

type Player interface {
	start(song string)
	stop()
}

func playSongs(player Player, songs []string) {
	for _, song := range songs {
		player.start(song)
	}
	player.stop()
}

func tryOut(player Player) {
	player.start("Test Track")
	player.stop()
	recorder := player.(TapeRecorder)
	recorder.record()
}

func main() {
	var tapeplayer TapePlayer
	var taperecorder TapeRecorder

	var mixSongs = []string{"Song1", "Song2", "Song3"}
	playSongs(tapeplayer, mixSongs)
	playSongs(taperecorder, mixSongs)

	tryOut(taperecorder)
}
