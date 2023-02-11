package cmd

import (
	"log"
	"github.com/dinolupo/camunda-utility/pkg/camunda/client"
	"github.com/spf13/cobra"
	"os"
)

// deleteDefinitionCmd represents the deleteDefinition command
var deleteDefinitionCmd = &cobra.Command{
	Use:   "deleteDefinition",
	Short: "Delete one or all Camunda definitions along with instances",
	Long: `Use this command to delete one or all definitions and its instances passing, for example:

camunda-utility deleteDefinition --key @all
camunda-utility deleteDefinition --key <process-definition-key>`,
	Run: func(cmd *cobra.Command, args []string) {
		if key == "" {
			log.Println("The parameter key must have a value.")
			os.Exit(1)
		}
		query := make(map[string]string)
		if key != "@all" {
			query["key"] = key
		}
		pd := client.ProcessDefinition{Client: Camunda}
		result, err := pd.GetList(query)
		if err != nil {
			log.Printf("ERROR: %+v\n", err.Error())
			os.Exit(1)
		}

		if len(result) == 0 {
			log.Printf("No process definitions found with key=%+v\n", key)
			os.Exit(0)
		}

		for _, s := range result {
			log.Printf("Deleting Process Definition Cascade: %+v\n", *s)
			err := pd.Delete(client.QueryProcessDefinitionBy{Id: &s.Id}, map[string]string{"cascade": "true"})
			if err != nil {
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteDefinitionCmd)
	deleteDefinitionCmd.Flags().StringVarP(&key, "key", "k", "", "select @all for all instances, or process-definition-key (required)")
	_ = deleteDefinitionCmd.MarkFlagRequired("key")
}
