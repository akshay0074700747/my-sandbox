package controllers

import (
	"net/http"

	"github.com/akshay0074700747/my-sandbox/docker"
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
		req.Command = enums.Commands.GOLANG
		req.Bind = string(enums.Binds.GOLANG)
		req.Container = string(enums.Containers.GOLANG)
		req.Extension = string(enums.Extensions.GOLANG)

	case enums.Languages.RUST:
		req.Command = enums.Commands.RUST
		req.Bind = string(enums.Binds.RUST)
		req.Container = string(enums.Containers.RUST)
		req.Extension = string(enums.Extensions.RUST)

	case enums.Languages.JAVA:
		req.Command = enums.Commands.JAVA
		req.Bind = string(enums.Binds.JAVA)
		req.Container = string(enums.Containers.JAVA)
		req.Extension = string(enums.Extensions.JAVA)

	case enums.Languages.PYTHON:
		req.Command = enums.Commands.PYTHON
		req.Bind = string(enums.Binds.PYTHON)
		req.Container = string(enums.Containers.PYTHON)
		req.Extension = string(enums.Extensions.PYTHON)

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the specified langauge is currently not Available"))
		return

	}

	logs, err := docker.ExecuteOnDocker(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(logs))
}
