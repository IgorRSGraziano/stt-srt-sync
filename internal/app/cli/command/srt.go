package command

import (
	"fmt"
	adapter "srtsync/internal/adapter/whisper"
	"srtsync/internal/core"

	"github.com/spf13/cobra"
)

func newSrtCommand(audioFile string, lyricFile *string, outputFile string) {
	whisperService := adapter.NewWhisperService("")

	service := core.NewSTTService(whisperService)

	// Executar a conversão
	text, err := service.GenerateSRT(audioFile, lyricFile)
	if err != nil {
		fmt.Println("Erro ao gerar SRT:", err)
		return
	}

	fmt.Println("Arquivo SRT gerado com sucesso:", outputFile)
	fmt.Println(*text)
}

func NewSRTCommand() *cobra.Command {
	var audioFile string
	var outputFile string
	var lyricFile string

	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate SRT file from audio and lyric (optional)",
		Run: func(cmd *cobra.Command, args []string) {
			// newSrtCommand(audioFile, lyricFile, outputFile)
		},
	}

	cmd.Flags().
		StringVarP(&audioFile, "audio", "a", "", "File path to audio")

	cmd.Flags().
		StringVarP(&outputFile, "output", "o", "output.srt", "Output file path")

	cmd.Flags().
		StringVarP(&lyricFile, "lyric", "l", "", "File path to lyric (simple text)")

	cmd.MarkFlagRequired("audio")

	return cmd
}
