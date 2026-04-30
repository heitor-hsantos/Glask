package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var ApiCfg = &cobra.Command{
	Use:   "api_key_gmn",
	Short: "Gemini Api Key config command",
	Long:  "Api Key config command it insert a new api Key for Gemini",
	Run: func(cmd *cobra.Command, args []string) {
		resposta := struct {
			ApiKey string `survey:"apikey"`
		}{}

		err := survey.Ask(apiKeyInsert, &resposta)
		if err != nil {
			fmt.Println("An error has occurred: ", err.Error())
			return
		}
		if resposta.ApiKey == "" {
			fmt.Println("No api key has been inserted")
			return
		}

		//Lógica para salvar em um arquivo e configuração
	},
}

var apiKeyInsert = []*survey.Question{
	{
		Name:     "apiKey",
		Prompt:   &survey.Password{Message: "Por favor, insira sua Gemini API Key: "},
		Validate: survey.Required,
	},
}
