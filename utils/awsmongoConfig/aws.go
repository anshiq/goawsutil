package awsmongoconfig

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func AwsS3Instance(awsaccessKey, awssecretKey, awsregion string) (*session.Session, error) {
	awsCreds := credentials.NewStaticCredentials(awsaccessKey, awssecretKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsregion),
		Credentials: awsCreds,
	})
	if err != nil {
		log.Panic("err in aws session creation", err)
		return nil, err
	}
	return sess, nil
}
