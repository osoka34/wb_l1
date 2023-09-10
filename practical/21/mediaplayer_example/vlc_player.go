package main

import "fmt"

type VlcPlayer struct{}

func (v *VlcPlayer) playVlc(fileName string) {
	fmt.Println("Playing vlc file:", fileName)
}

func (v *VlcPlayer) playMp4(fileName string) {
	// Ничего не делаем
}
