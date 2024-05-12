package awsuploadhandler

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

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
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	}
	if stat, err := os.Stat(filePath); !os.IsNotExist(err) {
		if stat.IsDir() {
			var allFilesArray []string
			traverseDir(&allFilesArray, filePath)
			var wg sync.WaitGroup
			allFilesArrayLen := len(allFilesArray)
			for fileIndex := 0; fileIndex < allFilesArrayLen; fileIndex += 3 {
				endIndex := fileIndex + 3
				if endIndex > allFilesArrayLen {
					endIndex = allFilesArrayLen
					noOfWg := allFilesArrayLen - fileIndex
					wg.Add(noOfWg)
				} else {
					wg.Add(3)
				}
				for i := fileIndex; i < endIndex; i++ {
					x := allFilesArray[i]
					go func(filePathName string) {
						fmt.Println("Uploading files: ", filePathName, " concurrenctly")
						defer wg.Done()
						fileDetails, _ := os.Stat(filePathName)
						fileObj := awsmongoconfig.HandleUploadCredsInstance(
							filePathName,
							collection,
							s3Ins,
							getRelativePath(filePathName),
							fileDetails.Name(),
							fileDetails.Size(),
						)
						err = awsmongoconfig.HandleUploadAws(*(fileObj))
						if err != nil {
							fmt.Print(err)
						}
					}(x)
				}
				wg.Wait()
			}
			wg.Wait()
		} else {
			fileDetails, _ := os.Stat(filePath)
			fileObj := awsmongoconfig.HandleUploadCredsInstance(
				filePath,
				collection,
				s3Ins,
				getRelativePath(filePath),
				fileDetails.Name(),
				fileDetails.Size(),
			)
			fmt.Println("Uploading single file: ", filePath)
			err = awsmongoconfig.HandleUploadAws(*(fileObj))
			if err != nil {
				fmt.Print(err)
			}
		}

	}

}
func getRelativePath(filePath string) string {
	homeDir, _ := os.UserHomeDir()
	index := strings.Index(filePath, homeDir)
	if index != -1 {
		return filePath[index+len(homeDir):]
	}
	return filePath
}
