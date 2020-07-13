/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Displays node in which the resources is scheduled.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be implemented.")
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	
	cf = genericclioptions.NewConfigFlags(true)

	nodeCmd.Flags().BoolP(allNamespacesFlag, "A", false, "query all objects in all API groups, both namespaced and non-namespaced")

	//AddFlags binds client configuration flags to a given flagset
	cf.AddFlags(nodeCmd.Flags())
	
	if err := flag.Set("logtostderr", "true"); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set logtostderr flag: %v\n", err)
		os.Exit(1)
	}
}
