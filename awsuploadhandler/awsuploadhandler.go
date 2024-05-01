package awsuploadhandler

import (
	"fmt"
	"os"
	"path"

	"github.com/anshiq/goawsutil/confighandle"
	awsmongoconfig "github.com/anshiq/goawsutil/utils/awsmongoConfig"
	// "github.com/anshiq/goawsutil/awsmongoConfig"
)

func traverseDir(allFilesArray *[]string, pathx string) error {
	pathStat, err := os.Stat(pathx)
	if err != nil {
		return err
	}
	if pathStat.IsDir() {
		arr, err := os.ReadDir(pathx)
		if err != nil {
			return err
		}
		for _, x := range arr {
			fullPath := path.Join(pathx, x.Name())
			if x.IsDir() {
				err = traverseDir(allFilesArray, fullPath)
				if err != nil {
					return err
				}
			} else {
				*allFilesArray = append(*allFilesArray, fullPath)
			}
		}
		return nil
	}
	*allFilesArray = append(*allFilesArray, pathx)
	return nil
}
func UploadFile(filePath string) {
	configs, _ := confighandle.GetConfigStruct()
	s3Ins, erraws := awsmongoconfig.AwsS3Instance(configs.AWSAccessKey, configs.AWSSecretKey, configs.AWSRegion)
	mongoDBInstance, err := awsmongoconfig.NewMongoDBInstance(configs.MongoURI, configs.DBname)
	collection := mongoDBInstance.Database.Collection("allfiles")
	if erraws != nil || err != nil {
		fmt.Print(erraws, err)
		return
	}

	fmt.Print(s3Ins, collection)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	}
	if stat, err := os.Stat(filePath); !os.IsNotExist(err) {
		if stat.IsDir() {
			var allFilesArray []string
			traverseDir(&allFilesArray, filePath)
			for _, x := range allFilesArray {
				awsmongoconfig.HandleUploadAws(x)
			}
		} else {
			awsmongoconfig.HandleUploadAws(filePath)
		}

	}

}
