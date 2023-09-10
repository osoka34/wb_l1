package main

func main() {
	audioPlayer := NewAudioPlayer()

	audioPlayer.play("mp3", "song.mp3")
	audioPlayer.play("vlc", "video.vlc")
	audioPlayer.play("mp4", "movie.mp4")
	audioPlayer.play("avi", "video.avi")
}
