package main

import (
	"fmt"
	"net/http"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		fmt.Println("Error in render")
		app.errorLog.Println(err)
	}
}
