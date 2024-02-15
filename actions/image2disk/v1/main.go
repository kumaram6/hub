package main

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/tinkerbell/hub/actions/image2disk/v1/pkg/image"
)

func main() {
	fmt.Printf("IMAGE2DISK - Cloud image streamer\n------------------------\n")
	disk := os.Getenv("DEST_DISK")
	// Check if a string is empty
	if len(disk) == 0 {
		// Get a list of drives
		drives, err := image.GetDrives()
		if err != nil {
			log.Error(err)
			return
		}
		detectedDisk, err := image.DriveDetection(drives)
		if err != nil {
			log.Error(err)
			return
		}
		log.Infof("Detected drive: [%s] ", detectedDisk)
		disk = detectedDisk
	} else {
		log.Infof("Drive provided by the user: [%s] ", disk)
	}

	img := os.Getenv("IMG_URL")
	compressedEnv := os.Getenv("COMPRESSED")

	// We can ignore the error and default compressed to false.
	cmp, _ := strconv.ParseBool(compressedEnv)

	// Write the image to disk
	err := image.Write(img, disk, cmp)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Successfully written [%s] to [%s]", img, disk)
}
