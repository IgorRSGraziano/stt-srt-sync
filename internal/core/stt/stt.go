package stt

type Text struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Text  string `json:"text"`
}

type STTService interface {
	TranscribeAudio(audioPath string) ([]Text, error)
}
