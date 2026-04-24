package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando base quando chamado sem subcomandos
var rootCmd = &cobra.Command{
	Use:   "glask",
	Short: "Glask - AI no seu terminal",
	Long:  `Glask é uma ferramenta de linha de comando para integrar IA com seu terminal.`,
}

// Execute adiciona todos os comandos filhos ao comando raiz e define os sinalizadores apropriadamente.
// Isso é chamado por main.main(). Só precisa acontecer uma vez para o rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Adicione seus sub-comandos aqui
	rootCmd.AddCommand(ApiCfg)
}
