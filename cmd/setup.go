package cmd

import (
	"github.com/spf13/cobra"
)

var ApiCfg = &cobra.Command{
	Use:   "api_key_gmn",
	Short: "Gemini Api Key config command",
	Long:  "Api Key config command it insert a new api Key for Gemini",
	Run:   func(cmd *cobra.Command, args []string) {},
}
