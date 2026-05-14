package cmd

import (
	"Glask/internal/config"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type apiKeyAnswers struct {
	ApiKey string `survey:"ApiKey"`
}

var ApiCfg = &cobra.Command{
	Use:   "api_key_gmn",
	Short: "Gemini Api Key config command",
	Long:  "Api Key config command it insert a new api Key for Gemini",
	Run: func(cmd *cobra.Command, args []string) {
		answers := apiKeyAnswers{}

		err := survey.Ask(apiKeyInsert, &answers)
		if err != nil {
			fmt.Println("An error has occurred: ", err.Error())
			return
		}
		if answers.ApiKey == "" {
			fmt.Println("No api key has been inserted")
			return
		}

		err = config.SaveAPIKey(answers.ApiKey)
		if err != nil {
			fmt.Println("Error while saving api key: ", err.Error())
			return
		}
		fmt.Println("✅ Gemini API key stored with success")
	},
}

var apiKeyInsert = []*survey.Question{
	{
		Name:     "ApiKey",
		Prompt:   &survey.Password{Message: "Por favor, insira sua Gemini API Key: "},
		Validate: survey.Required,
	},
}

var apiKeyDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Gemini Api Key",
	Long:  "Delete Gemini Api Key from keyring",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.DeleteAPIKey()
		if err != nil {
			fmt.Println("Error while deleting api key: ", err.Error())
			return
		}
		fmt.Println("✅ Gemini API key deleted with success")
	},
}
var getApiKey = &cobra.Command{
	Use:   "getApiKey",
	Short: "Get Gemini Api Key",
	Long:  "Get Gemini Api Key from keyring",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := config.GetAPIKey()
		if err != nil {
			fmt.Println("Error while getting api key: ", err.Error())
			return
		}
		fmt.Println(result)

	},
}

func init() {
	ApiCfg.AddCommand(apiKeyDelete, getApiKey)

}
