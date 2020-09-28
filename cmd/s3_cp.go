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
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//func exitErrorf(msg string, args ...interface{}) {
//	fmt.Fprintf(os.Stderr, msg+"\n", args...)
//	os.Exit(1)
//}

// s3cpCmd represents the cp command
var s3cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy data to S3 buckets",
	Long: "Copy data to S3 buckets",
	Run: func(cmd *cobra.Command, args []string) {

		bucket, _:= cmd.Flags().GetString("bucket")
		folder, _:= cmd.Flags().GetString("source")
		region, _:= cmd.Flags().GetString("region")

		if (folder == "" || bucket == "") {
			fmt.Println("Null Options.....")
            cmd.Help()
            os.Exit(1)
       	}
		if (region == "") {
			region = "us-west-2"
		}

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region)},
		)
		uploader := s3manager.NewUploader(sess)

		// Read files from folder
		err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					exitErrorf("Unable to open file %q, %v", err)
				}

				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(bucket),
					Key: aws.String(info.Name()),
					Body: file,
				})
				fmt.Printf("Successfully uploaded %q to %q\n", path, bucket)
        	}
        	return nil
    	})
    	if err != nil {
			exitErrorf("Unable to upload %q to %q, %v", folder, bucket, err)
    	}
	},
}

func init() {
	awsS3Cmd.AddCommand(s3cpCmd)

	// Here you will define your flags and configuration settings.
	s3cpCmd.Flags().StringP("source", "s", "", "Source folder")
	s3cpCmd.Flags().StringP("bucket", "b", "", "Target bucket to upload content")
	s3cpCmd.Flags().StringP("region", "r", "", "aws region default: us-west-2")

}
