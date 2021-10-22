// tmp.link CLI uploader
// Language: go
// Path: main.go

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var (
	h        bool
	v        bool
	token    string
	model    string
	filepath string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&token, "k", "N/A", "your account token")
	flag.StringVar(&model, "m", "0", "upload model, 0 : only 24 hours, 1,2,3: 24,72,")
	flag.StringVar(&filepath, "f", "", "upload file path")
}

func main() {
	var upload_api = "https://connect.tmp.link/api_v2/cli_uploader"
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if v {
		fmt.Println("tmp.link CLI uploader")
		fmt.Println("Version: 1")
	}

	if token == "N/A" {
		fmt.Println("Please input your token:")
		fmt.Scanln(&token)
	}

	checkFilePath()

	f, err := os.Open(filepath)
	exitIfErr(err)
	defer f.Close()

	fields := map[string]string{
		"filename": filepath,
		"token":    token,
		"model":    model,
	}
	res, err := multipartUpload(upload_api, f, fields)
	exitIfErr(err)
	fmt.Println("res: ", res)
}

// check file path,if not exist, select one and check again.
func checkFilePath() {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("file not exist, please select one:")
		fmt.Scanln(&filepath)
		checkFilePath()
	}
}

func exitIfErr(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func multipartUpload(destURL string, f io.Reader, fields map[string]string) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", fields["filename"])
	if err != nil {
		return nil, fmt.Errorf("CreateFormFile %v", err)
	}

	_, err = io.Copy(fw, f)
	if err != nil {
		return nil, fmt.Errorf("copying fileWriter %v", err)
	}

	for k, v := range fields {
		_ = writer.WriteField(k, v)
	}

	err = writer.Close() // close writer before POST request
	if err != nil {
		return nil, fmt.Errorf("writerClose: %v", err)
	}

	resp, err := http.Post(destURL, writer.FormDataContentType(), body)
	if err != nil {
		return nil, err
	}

	return resp, nil

	// req, err := http.NewRequest("POST", destURL, body)
	// if err != nil {
	//  return nil, err
	// }

	// req.Header.Set("Content-Type", writer.FormDataContentType())

	// if req.Close && req.Body != nil {
	//  defer req.Body.Close()
	// }

	// return http.DefaultClient.Do(req)
}
