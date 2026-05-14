package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glask",
	Short: "Glask - AI no seu terminal",
	Long:  `Glask é uma ferramenta de linha de comando para integrar IA com seu terminal.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(ApiCfg)
}
