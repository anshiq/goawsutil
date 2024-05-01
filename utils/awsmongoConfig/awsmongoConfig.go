package awsmongoconfig

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"time"
)

type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FileName   string             `bson:"fileName"`
	FileSize   int64              `bson:"fileSize"`
	LocalPath  string             `bson:"localpath"`
	UploadedAt time.Time          `bson:"uploadedAt"`
	AWSUrl     string             `bson:"awsurl"`
}

type HandleUploadCreds struct {
	fileFolderName string
	collection     *mongo.Collection
	s3Ins          *session.Session
	awsPathKey     string
	name           string
	size           int64
}

func HandleUploadAws(creds HandleUploadCreds) {
	file, err := os.Open(creds.fileFolderName)
	if err != nil {
		log.Fatal("err", err)
	}
	defer file.Close()
	s3 := s3manager.NewUploader(creds.s3Ins)
	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String("my-goawsutil"),
		Key:    aws.String(creds.awsPathKey),
		Body:   bufio.NewReader(file),
	}
	others, errs3 := s3.UploadWithContext(aws.BackgroundContext(), uploadInput)
	if errs3 != nil {
		log.Fatal("err", errs3)
		return
	}

	testFile := File{
		FileName:   creds.name,
		FileSize:   creds.size,
		UploadedAt: time.Now(),
		LocalPath:  creds.fileFolderName,
		AWSUrl:     others.Location,
	}
	if err != nil {
		log.Panic("error while creating a mongo db instance", err)
	}
	resmongo, errInsert := creds.collection.InsertOne(context.Background(), testFile)
	if err != nil {
		log.Fatal("Error inserting test data:", errInsert)
	}
	fmt.Println(resmongo.InsertedID, "uploaded and stored in mongo")
}
