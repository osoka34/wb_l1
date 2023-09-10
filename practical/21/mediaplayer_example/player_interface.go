package main

// Целевой интерфейс
type MediaPlayer interface {
	play(fileType, fileName string)
}
