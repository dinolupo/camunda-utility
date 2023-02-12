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
	"github.com/dinolupo/camunda-utility/pkg/camunda/client"
	"github.com/spf13/cobra"
	//"github.com/dinolupo/camunda-utility/pkg/utils"
	"os"
	"log"
)

// deleteInstancesCmd represents the deleteInstances command
var deleteInstancesCmd = &cobra.Command{
	Use:   "deleteInstances",
	Short: "Delete all instances of one or all process definitions",
	Long: `Use this option to delete all instances of a process definition key or all of them, for example:

	camunda-utility deleteInstances --key <process-definition-key>
	camunda-utility deleteInstances --key @all`,
	Run: func(cmd *cobra.Command, args []string) {
		if key == "" {
			log.Println("The parameter key must have a value.")
			os.Exit(1)
		}
		// ----------------------
		query := make(map[string]string)
		if key != "@all" {
			query["key"] = key
		}
		pd := client.ProcessDefinition{Client: Camunda}
		result, err := pd.GetList(query)
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}

		if len(result) == 0 {
			if key != "@all" {
				log.Printf("No process definitions found with key=%+v\n", key)
			} else {
				log.Printf("No process definitions found.\n")
			}
			os.Exit(0)
		}

		// Print list of process definitions
		// for _, s := range result {
		// 	res, _ := utils.PrettyStruct(*s)
		// 	log.Printf("%+v\n", res)
		// }

		for _, s := range result {
			log.Printf("Deleting all Process Instances of Definition: %+v\n", s.Id)

			pi := client.ProcessInstance{Client: Camunda}
			query["processDefinitionId"] = s.Id
			result, err := pi.GetList(query)
			if err != nil {
				log.Fatal(err.Error())
				os.Exit(1)
			}
			if len(result) == 0 {
				log.Printf("No process instances found with query=%+v\n", query)
			}

			for _, instance := range result {
				log.Printf("\tDeleting Instance: %+v query: %+v\n", instance.Id, query)
				err := pi.Delete(instance.Id, nil)
				if err != nil {
					log.Fatal(err.Error())
					os.Exit(1)
				}
			}
			if err != nil {
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteInstancesCmd)
	deleteInstancesCmd.Flags().StringVarP(&key, "key", "k", "", "select @all for all definitions, or process-definition-key (required)")
	_ = deleteInstancesCmd.MarkFlagRequired("key")
}
