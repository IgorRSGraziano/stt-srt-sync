package cli

import (
	"fmt"
	"os"
	"srtsync/internal/app/cli/command"

	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "srt-cli",
		Short: "Audio to SRT generator",
	}

	rootCmd.AddCommand(command.NewSRTCommand())

	return rootCmd
}

func Execute() {
	if err := NewCLI().Execute(); err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
