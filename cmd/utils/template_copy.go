package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

// Initialize S3 client
var s3Client *s3.S3

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	aws_access_id := os.Getenv("AWS_ACCESS_ID")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"), // Update with your region
		Credentials: credentials.NewStaticCredentials(aws_access_id, aws_secret_key, ""),
	}))
	s3Client = s3.New(sess)
}

// CopyTemplate loads the template from S3 and copies it to the destination file
func CopyTemplate(s3Bucket, s3Key, dest string, replacements ...map[string]string) error {
	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(s3Key),
	})
	if err != nil {
		fmt.Println("Error fetching from S3:", err)
		return err
	}
	defer result.Body.Close()

	s3Content, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error reading S3 content:", err)
		return err
	}

	content := string(s3Content)

	if len(replacements) > 0 {
		for placeholder, replacement := range replacements[0] {
			content = strings.ReplaceAll(content, placeholder, replacement)
		}
	}

	err = os.WriteFile(dest, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to destination file:", err)
		return err
	}
	return nil
}
