// Package controller loads the routes for each of the controllers.
package controller

import (
	"jayesh/ocr-web-application/controller/about"
	"jayesh/ocr-web-application/controller/analyze"
	"jayesh/ocr-web-application/controller/debug"
	"jayesh/ocr-web-application/controller/fileTransfer"
	"jayesh/ocr-web-application/controller/home"
	"jayesh/ocr-web-application/controller/login"
	"jayesh/ocr-web-application/controller/notepad"
	"jayesh/ocr-web-application/controller/register"
	"jayesh/ocr-web-application/controller/static"
	"jayesh/ocr-web-application/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
	fileTransfer.Load()
	analyze.Load()
}
