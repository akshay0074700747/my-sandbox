package model

type CodeExecutionRequest struct {
	SourceCode []byte
	Language   string
	FileSize   int64
	FileName   string
	Container  string
	Bind       string
	Extension  string
	Command    []string
}

type CodeExecutionResponce struct {
	Logs          string `json:"Logs"`
	ExecutionTime string `json:"ExecutionTime"`
	Memmory       string `json:"Memmory"`
}
