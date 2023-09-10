package main

import "fmt"

type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) MediaAdapter {
	if audioType == "vlc" {
		return MediaAdapter{&VlcPlayer{}}
	}
	// Здесь можно добавить поддержку других форматов
	return MediaAdapter{}
}

func (m MediaAdapter) play(fileType, fileName string) {
	if fileType == "vlc" {
		m.advancedPlayer.playVlc(fileName)
	} else {
		fmt.Println("Invalid media. " + fileType + " format not supported")
	}
}
