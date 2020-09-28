/*
Copyright © 2020 Knyt.io

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
	"os"
	"github.com/spf13/cobra"
)

// awsS3 represents the s3 command
var awsS3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "S3 Ransomware Simulator",
	Long: "S3 Ransomware Simulator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("s3 called")
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(awsS3Cmd)
}
