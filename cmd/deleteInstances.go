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
	"fmt"
	"github.com/dinolupo/camunda-utility/pkg/camunda/client"
	"github.com/spf13/cobra"
	"os"
)

// deleteInstancesCmd represents the deleteInstances command
var deleteInstancesCmd = &cobra.Command{
	Use:   "deleteInstances",
	Short: "Delete all instances of a process definition",
	Long: `Use this option to delete all instances of a specified process definition key, for example:

	camunda-utility deleteInstances --key <process-definition-key>`,
	Run: func(cmd *cobra.Command, args []string) {
		if key == "" {
			fmt.Println("The parameter key must have a value.")
			os.Exit(1)
		}
		query := make(map[string]string)
		query["processDefinitionKey"] = key
		pd := client.ProcessInstance{Client: Camunda}
		result, err := pd.GetList(query)
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err.Error())
			os.Exit(1)
		}
		if len(result) == 0 {
			fmt.Printf("No process instances found with processDefinitionKey=%+v\n", key)
			os.Exit(0)
		}

		for _, s := range result {
			fmt.Printf("Deleting Instance: %+v\n", s.Id)
			err := pd.Delete(s.Id, query)
			if err != nil {
				fmt.Printf("ERROR: %+v\n", err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteInstancesCmd)

	deleteInstancesCmd.Flags().StringVarP(&key, "key", "k", "", "the process-definition-key (required)")
	_ = deleteInstancesCmd.MarkFlagRequired("key")
}
