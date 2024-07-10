package controllers

import (
	"fmt"
	"net/http"

	"github.com/akshay0074700747/my-sandbox/enums"
	"github.com/akshay0074700747/my-sandbox/model"
)

func ExecuteCode(w http.ResponseWriter, r *http.Request) {

	var req model.CodeExecutionRequest

	e, rr, err := r.FormFile("SourceCode")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	req.FileName = rr.Filename
	req.FileSize = rr.Size

	file := make([]byte, req.FileSize)

	if _, err = e.Read(file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	req.SourceCode = file

	req.Language = r.FormValue("Language")

	switch enums.Langage(req.Language) {
		
	case enums.Languages.GOLANG:

	case enums.Languages.RUST:

	case enums.Languages.JAVA:

	case enums.Languages.PYTHON:

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the specified langauge is currently not Available"))
		return

	}

	fmt.Println(string(req.SourceCode))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("here is your result"))
}
