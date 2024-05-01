package main

import (
	"fmt"
	"os"
	"path"

	"github.com/anshiq/goawsutil/awsuploadhandler"
	"github.com/anshiq/goawsutil/confighandle"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "config" {
			confighandle.CreateOrCheckConfig()
		} else if args[1] == "upload" && len(args[2]) > 0 {
			// fmt.Print(args[3])
			cwd, _ := os.Getwd()
			filePath := path.Join(cwd, args[2])
			awsuploadhandler.UploadFile(filePath)
			// handle uploading..
		} else if args[1] == "reconfig" {
			confighandle.RemoveConfigFile()
			confighandle.CreateOrCheckConfig()
		}
	} else {
		mapOfArgs := make(map[string]string) // Initialize mapOfArgs
		mapOfArgs["upload"] = "goawsutil upload-dir pathOfDir - to upload files in given dir"
		// mapOfArgs["upload"] = "goawsutil upload-file pathOfFile - to upload files"
		mapOfArgs["config"] = "goawsutil config - to configure aws and mongo uri"
		for key, value := range mapOfArgs {
			fmt.Println(key, "(", value, ")")
		}

	}
}
