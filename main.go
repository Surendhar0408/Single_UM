package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	//Log file creation
	file, err := os.OpenFile("gatekeeper_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	log.SetOutput(ioMultiWriter(os.Stdout, file))

	//-------------------------------------------------------------------------------------------------
	//Arguments
	args := os.Args
	if len(args) <= 1 {
		log.Println("Provide valid arguments!!!\nUse --help for more details")
		return
	}

	log.Println("**********************************Start of Log**********************************")
	log.Println("Arguments:", args)

	switch strings.TrimSpace(strings.ToLower(args[1])) {

	case "config":

		return

	case "release":

		return

	case "down":

		return
	case "help", "h", "-h", "--h", "--help":
		Help()
		return

	default:
		log.Println("Please provide valid argument to perform.")

	}
	log.Println("**********************************End of Log**********************************")
}
func ioMultiWriter(writers ...io.Writer) io.Writer {
	return io.MultiWriter(writers...)
}
