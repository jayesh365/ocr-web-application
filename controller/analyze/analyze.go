package analyze

import (
	"fmt"
	"jayesh/ocr-web-application/lib/flight"
	"net/http"
	"os/exec"

	"github.com/blue-jay/core/router"
)

func Load() {
	router.Post("/ocr-omr", Analyze)
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	msg, err := exec.Command("cmd", "/C", "python", "./python/main.py").Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	c.Redirect("/")
}
