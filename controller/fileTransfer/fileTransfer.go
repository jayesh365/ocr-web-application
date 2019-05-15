package fileTransfer

import (
	"fmt"
	"io/ioutil"
	"jayesh/ocr-web-application/lib/flight"
	"net/http"

	"github.com/blue-jay/core/router"
)

func Load() {
	router.Post("/file-transfer", FileTransfer)
}

func FileTransfer(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("choose-file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./uploads/form.pdf", fileBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
	// redirect to homepage
	c.Redirect("/")
}
