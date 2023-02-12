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
package main

import (
	"github.com/dinolupo/camunda-utility/cmd"
	"log"
	//"github.com/rivo/tview"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("Log initialized")
}

func main() {
	cmd.Execute()
	// app := tview.NewApplication()
	// form := tview.NewForm().
	// 	AddInputField("Label:", "", 20, nil, nil).
	// 	AddButton("Go", func() {
	// 		app.Stop()
	// 	})
	// if err := app.SetRoot(form, true).Run(); err != nil {
	// 	panic(err)
	// }
}

