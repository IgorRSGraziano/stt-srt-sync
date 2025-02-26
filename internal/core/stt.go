package core

type STTService interface {
	GenerateSRT(audioPath string, lyric *string) (*string, error)
}

func NewSTTService(service STTService) STTService {
	return service
}
