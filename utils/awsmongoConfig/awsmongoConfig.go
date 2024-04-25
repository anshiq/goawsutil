package awsmongoconfig

import (
	// "bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anshiq/goawsutil/confighandle"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FileName   string             `bson:"fileName"`
	FileSize   int64              `bson:"fileSize"`
	UploadedAt time.Time          `bson:"uploadedAt"`
	AWSUrl     string             `bson:"awsurl"`
	// Add any other fields as needed
}

func HandleUploadAws(fileFolderName string) {
	fileName, _ := os.Stat(fileFolderName)
	if fileName.IsDir() {
		fmt.Print("your given path is a dir not a file")
		return
	}
	// name := fileName.Name()
	// size := fileName.Size()
	// fileName.Sys()
	// awsPath := name + fmt.Sprint(size)
	testFile := File{
		FileName:   "test.txt",
		FileSize:   1024,
		UploadedAt: time.Now(),
	}
	configs, _ := confighandle.GetConfigStruct()
	mongoDBInstance, err := NewMongoDBInstance(configs.MongoURI, configs.DBname)
	if err != nil {
		log.Panic("error while creating a mongo db instance", err)
	}
	collection := mongoDBInstance.Database.Collection("allfiles")
	resmongo, errInsert := collection.InsertOne(context.Background(), testFile)
	if err != nil {
		log.Fatal("Error inserting test data:", errInsert)
	}
	fmt.Print(resmongo)
	// s3Ins, erraws := AwsS3Instance(configs.AWSAccessKey, configs.AWSSecretKey, configs.AWSRegion)
	// if erraws != nil {
	// 	fmt.Print(erraws)
	// 	return
	// }
	// file, err := os.Open(fileFolderName)
	// if err != nil {
	// 	log.Fatal("err", err)
	// 	// Handle error
	// }
	// defer file.Close()
	// s3 := s3manager.NewUploader(s3Ins)
	// uploadInput := &s3manager.UploadInput{
	// 	Bucket: aws.String("my-goawsutil"),
	// 	Key:    aws.String(awsPath),
	// 	Body:   bufio.NewReader(file),
	// }
	// others, errs3 := s3.UploadWithContext(aws.BackgroundContext(), uploadInput)
	// if errs3 != nil {
	// 	log.Fatal("err", errs3)
	// 	return
	// }
	// fmt.Print("file upload succ", others)
	// s3.Upload()

}
