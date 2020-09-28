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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"github.com/spf13/cobra"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// mbCmd represents the mb command
var mbCmd = &cobra.Command{
	Use:   "mb",
	Short: "Create S3 bucket",
	Long: "Create S3 bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucket, _:= cmd.Flags().GetString("bucket")
   		if bucket == "" {
			cmd.Help()
			os.Exit(1)
   		}

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		// Create the S3 Bucket
		_, err = svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
		exitErrorf("Unable to create bucket %q, %v", bucket, err)
		}

		// Wait until bucket is created before finishing
		fmt.Printf("Waiting for bucket %q to be created...\n", bucket)

		err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			exitErrorf("Error occurred while waiting for bucket to be created, %v", bucket)
		}
		fmt.Printf("Bucket %q successfully created\n", bucket)
	},
}

func init() {
	awsS3Cmd.AddCommand(mbCmd)

	// Here you will define your flags and configuration settings.
	 mbCmd.Flags().StringP("bucket", "b", "", "bucket name")

	// mbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
