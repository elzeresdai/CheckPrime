package internal

type requestStruct struct {
	Numbers []int `json:"numbers"`
}

type errorStruct struct {
	code    int
	index   string
	message string
}
