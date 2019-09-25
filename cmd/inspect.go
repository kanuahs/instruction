// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io/ioutil"
	"log"

	instruction "github.com/kanuahs/instruction/pkg"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
	v1 "k8s.io/api/apps/v1"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("inspect called")
		var deployment v1.Deployment
		// Using a go-client deployment as a sample struct
		deploymentFile, err := cmd.Flags().GetString("filename")
		if err != nil {
			log.Fatal(err)
		}
		if deploymentFile != "" {
			file, _ := ioutil.ReadFile(deploymentFile)
			yaml.Unmarshal(file, &deployment)
		} else {
			deployment = v1.Deployment{}
		}
		log.Println(deployment)
		instruction.InspectStruct("v1.Deployment", deployment)

	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	inspectCmd.Flags().String("filename", "", "File containing a struct to be unmarshalled")
}
