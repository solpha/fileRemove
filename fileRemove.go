package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func fileRemove(filename []string) {
	for i := 0; i < len(filename); i++ {
		fmt.Printf("Removed file: %s\n", filename[i])
		err := os.Remove("/Volumes/Time Machine/iTunes/iTunes Media/Home Videos/" + filename[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	ls := exec.Command("ls", "/Volumes/Time Machine/iTunes/iTunes Media/Home Videos")
	lsOut, err := ls.Output()
	if err != nil {
		log.Fatal(err)
	}

	xmlFiles := regexp.MustCompile("(.*).xml")
	htmlFiles := regexp.MustCompile("(.*).html")
	jpegFiles := regexp.MustCompile("(.*).jpeg")

	findxmlFiles := xmlFiles.FindAllString(string(lsOut), -1)
	fileRemove(findxmlFiles)
	findhtmlFiles := htmlFiles.FindAllString(string(lsOut), -1)
	fileRemove(findhtmlFiles)
	findjpegFiles := jpegFiles.FindAllString(string(lsOut), -1)
	fileRemove(findjpegFiles)

}
