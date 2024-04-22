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
		} else if args[1] == "upload" && args[2] == "-file" && len(args[3]) > 0 {
			// fmt.Print(args[3])
			cwd, _ := os.Getwd()
			filePath := path.Join(cwd, args[3])
			awsuploadhandler.UploadFile(filePath)
			// handle uploading..
		}
	} else {
		mapOfArgs := make(map[string]string) // Initialize mapOfArgs
		mapOfArgs["upload"] = "goawsutil upload pathOfFile - to upload files"
		mapOfArgs["config"] = "goawsutil config - to configure aws and mongo uri"
		for key, value := range mapOfArgs {
			fmt.Println(key, "(", value, ")")
		}

	}
	//remove k[0]
	// confighandle.CreateOrCheckConfig()
}
