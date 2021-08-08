package cmd

import (
	"fmt"
	"github.com/dinolupo/camunda-utility/pkg/camunda/client"
	"github.com/spf13/cobra"
	"os"
)

var key *string

// deleteDefinitionCmd represents the deleteDefinition command
var deleteDefinitionCmd = &cobra.Command{
	Use:   "deleteDefinition",
	Short: "Delete Camunda definition and instances for a single or all process definitions",
	Long: `Use this command to delete definition and its instances of a single process definition
	key, or delete all definitions and their associated instances, for example:

camunda-utility deleteDefinition --key @all
camunda-utility deleteDefinition --key <process-definition-key>`,
	Run: func(cmd *cobra.Command, args []string) {
		if *key == "" {
			fmt.Println("The parameter key must have a value.")
			os.Exit(1)
		}
		query := make(map[string]string)
		if *key != "@all" {
			query["key"] = *key
		}
		pd := client.ProcessDefinition{Client: Camunda}
		result, err := pd.GetList(query)
		if err != nil {
			os.Exit(1)
		}
		for _, s := range result {
			fmt.Printf("Deleting Process Definition Cascade: %+v\n", *s)
			err := pd.Delete(client.QueryProcessDefinitionBy{Id: &s.Id}, map[string]string{"cascade": "true"})
			if err != nil {
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteDefinitionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteDefinitionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	key = deleteDefinitionCmd.Flags().String("key", "", "select @all for all instances, or process-definition-key (required)")
	_ = deleteDefinitionCmd.MarkFlagRequired("key")
}
