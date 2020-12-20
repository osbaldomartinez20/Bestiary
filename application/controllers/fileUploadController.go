package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func UploadFile(r *http.Request) string {
	log.Println("Upload File Endpoint.")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	_ = r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("beastImage")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return "Error Retrieving the File."
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return "Cannot Read The File."
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	return "Successfully Uploaded File\\n"
}