package awsuploadhandler

import (
	"fmt"
	"os"

	awsmongoconfig "github.com/anshiq/goawsutil/utils/awsmongoConfig"
	// "github.com/anshiq/goawsutil/awsmongoConfig"
)

func UploadFile(filePath string) {
	fmt.Print(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	}
	awsmongoconfig.HandleUploadAws(filePath)

}
