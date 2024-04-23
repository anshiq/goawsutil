package awsmongoconfig

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/anshiq/goawsutil/confighandle"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func HandleUploadAws(fileFolderName string) {
	fileName, _ := os.Stat(fileFolderName)
	if fileName.IsDir() {
		fmt.Print("your given path is a dir not a file")
		return
	}
	name := fileName.Name()
	size := fileName.Size()
	// fileName.Sys()
	awsPath := name + fmt.Sprint(size)

	configs, _ := confighandle.GetConfigStruct()
	mongoDBInstance, err := NewMongoDBInstance(configs.MongoURI, configs.DBname)
	if err != nil {
		log.Panic("error while creating a mongo db instance", err)
	}
	collection := mongoDBInstance.Database.Collection("allfiles")
	fmt.Print(collection)
	s3Ins, erraws := AwsS3Instance(configs.AWSAccessKey, configs.AWSSecretKey, configs.AWSRegion)
	if erraws != nil {
		fmt.Print(erraws)
		return
	}
	file, err := os.Open("path/to/local/file.txt")
	if err != nil {
		log.Fatal("err", err)
		// Handle error
	}
	defer file.Close()
	s3 := s3manager.NewUploader(s3Ins)
	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String("my-goawsutil"),
		Key:    aws.String(awsPath),
		Body:   bufio.NewReader(file),
	}
	_, err = s3.UploadWithContext(aws.BackgroundContext(), uploadInput)
	if err != nil {
		log.Fatal("err", err)
	}
	// s3.Upload()

}
