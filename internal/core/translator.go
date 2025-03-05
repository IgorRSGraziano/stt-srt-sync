package core

type TranslatorService interface {
	Translate(text *string, targetLanguage string) (*string, error)
}

func NewTranslatorService(service TranslatorService) TranslatorService {
	return service
}
