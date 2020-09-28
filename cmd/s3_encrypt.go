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
	"os"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var s3encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt S3 bucket",
	Long: "Encrypt S3 bucket",
	Run: func(cmd *cobra.Command, args []string) {
		bucket, _:= cmd.Flags().GetString("bucket")
   		if bucket == "" {
			cmd.Help()
			os.Exit(1)
   		}
		region, _:= cmd.Flags().GetString("region")
		if (region == "") {
			region = "us-west-2"
		}
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region)},
		)

		// Create S3 service client
		svc := s3.New(sess)

		// Get the list of items
		resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
		if err != nil {
			exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
		}

		for _, item := range resp.Contents {

			source_file := bucket + "/" + *item.Key
			target_bucket := bucket
			target_file := *item.Key + ".enc"

			copyObjectInput := &s3.CopyObjectInput{
				Bucket:     aws.String(target_bucket),
				CopySource: aws.String(source_file),
				Key:        aws.String(target_file),
				SSECustomerAlgorithm: aws.String("AES256"),
				SSECustomerKey: aws.String("B3DBCB8D7594F0A21D3D9E0EA3B75444"),
			}

			_,err := svc.CopyObject(copyObjectInput)
			if err != nil {
				exitErrorf("Error encrypt, %v", err)
			}

			// Delete the original item
			_, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(target_bucket), Key: aws.String(*item.Key)})
			if err != nil {
				exitErrorf("Unable to delete object %q from bucket %q, %v", *item.Key, bucket, err)
			}
			fmt.Println("Encrypted and deleting file: ", source_file)
		}
	},
}

func init() {
	awsS3Cmd.AddCommand(s3encryptCmd)

	// Here you will define your flags and configuration settings.
	s3encryptCmd.Flags().StringP("bucket", "b", "", "bucket name")
	s3encryptCmd.Flags().StringP("region", "r", "", "aws region default: us-west-2")
}
