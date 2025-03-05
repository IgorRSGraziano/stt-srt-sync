package command

import (
	"fmt"
	gpt "srtsync/internal/adapter/gpt"
	whisper "srtsync/internal/adapter/whisper"
	"srtsync/internal/core"

	"github.com/spf13/cobra"
)

func newSrtCommand(audioFile string, lyricFile *string, outputFile string, translateTo string) {
	whisperService := whisper.NewWhisperService("")
	gpt := gpt.NewGPTService("")

	sttService := core.NewSTTService(whisperService)
	translateService := core.NewTranslatorService(gpt)

	// Executar a conversão
	text, err := sttService.GenerateSRT(audioFile, lyricFile)
	if err != nil {
		fmt.Println("error on generate srt:", err)
		return
	}

	if translateTo != "" {
		translatedText, err := translateService.Translate(text, translateTo)
		if err != nil {
			fmt.Println("error on translate:", err)
			return
		}

		text = translatedText
	}

	fmt.Println("SRT file generated at", outputFile)
	fmt.Println(*text)
}

func NewSRTCommand() *cobra.Command {
	var audioFile string
	var outputFile string
	var lyricFile string

	var translateTo string

	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate SRT file from audio and lyric (optional)",
		Run: func(cmd *cobra.Command, args []string) {
			newSrtCommand(audioFile, &lyricFile, outputFile, translateTo)
		},
	}

	cmd.Flags().
		StringVarP(&audioFile, "audio", "a", "", "File path to audio")

	cmd.Flags().
		StringVarP(&outputFile, "output", "o", "output.srt", "Output file path")

	cmd.Flags().
		StringVarP(&lyricFile, "lyric", "l", "", "File path to lyric (simple text)")

	cmd.Flags().
		StringVarP(&translateTo, "translate", "t", "", "Translate lyric to another language")

	cmd.MarkFlagRequired("audio")

	return cmd
}
