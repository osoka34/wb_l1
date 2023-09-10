package main

import "fmt"

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func NewAudioPlayer() AudioPlayer {
	return AudioPlayer{}
}

func (a AudioPlayer) play(fileType, fileName string) {
	if fileType == "mp3" {
		fmt.Println("Playing mp3 file:", fileName)
	} else if fileType == "vlc" || fileType == "mp4" {
		a.mediaAdapter = NewMediaAdapter(fileType)
		a.mediaAdapter.play(fileType, fileName)
	} else {
		fmt.Println("Invalid media. " + fileType + " format not supported")
	}
}
