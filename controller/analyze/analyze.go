package analyze

import (
	"fmt"
	"jayesh/ocr-web-application/lib/flight"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/blue-jay/core/router"
)

func Load() {
	router.Post("/ocr-omr", Analyze)
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	switch runtime.GOOS {
	case "windows":
		msg, err := exec.Command("cmd", "/C", "python", "./python/main.py").Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}
		// sdaps cannot be run on windows
	case "linux":
		msg, err := exec.Command("python", "./python/main.py").Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}

		// sdaps add tif file to latex
		// 1.tif may need to be renamed
		msg, err = exec.Command("/bin/bash", "-c", "sdaps add ./sdaps ./uploads/form.tiff --convert").Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}
		// sdaps recongize latex_test
		// runs OMR recognition on the project directory latex_test
		// requires monochrome multipage TIFFs to already by added to the project (needs to be done some point prior to this)
		msg, err = exec.Command("/bin/bash", "-c", "sdaps recognize ./sdaps").Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}
		// sdaps export csv
		msg, err = exec.Command("/bin/bash", "-c", "sdaps csv export ./sdaps").Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}
		// now parse csv
		msg, err = exec.Command("python", "./python/parseCSV.py").Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}
	default:
		break
	}
	c.Redirect("/")
}
