package awsmongoconfig

import (
	"fmt"
	"github.com/anshiq/goawsutil/confighandle"
	"log"
)

func HandleUploadAws(fileFolderName string) {
	x, err := confighandle.GetConfigStruct()
	if err != nil {
		log.Panic("errr occur", err)
	}
	fmt.Print(x)
}
