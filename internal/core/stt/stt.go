package stt

type STTService interface {
	GenerateSRT(audioPath, lyric string) (*string, error)
}
