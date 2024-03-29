/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"github.com/dinolupo/camunda-utility/pkg/camunda/client"
	"github.com/dinolupo/camunda-utility/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

// listDefinitionsCmd represents the listDefinitions command
var listDefinitionsCmd = &cobra.Command{
	Use:   "listDefinitions",
	Short: "List all Process Definitions deployed",
	Long: `Use this command to list all process definitions in the specified Camunda rest endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		pd := client.ProcessDefinition{Client: Camunda}
		query := make(map[string]string)
		result, err := pd.GetList(query)
		if err != nil {
			log.Printf("ERROR: %+v\n", err.Error())
			os.Exit(1)
		}

		if len(result) == 0 {
			log.Printf("No process definition found")
		}

		for _, s := range result {
			res, _ := utils.PrettyStruct(*s)
			log.Printf("%+v,\n", res)
		}
	},
}

func init() {
	rootCmd.AddCommand(listDefinitionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDefinitionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDefinitionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
