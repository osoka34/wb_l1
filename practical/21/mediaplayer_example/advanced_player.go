package main

// Адаптируемый класс
type AdvancedMediaPlayer interface {
	playVlc(fileName string)
	playMp4(fileName string)
}
