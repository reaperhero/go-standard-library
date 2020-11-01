package stand

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func HttpPost(url, data string) (status int, result string) {
	body := bytes.NewBufferString(data)
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes)
}

func HttpPostWithToken(url, token, data string) (status int, result string) {
	body := bytes.NewBufferString(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes)
}

func HttpPOSTWithCustomToken(url, token, data string) (status int, result string) {
	body := bytes.NewBufferString(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)

	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes)
}

func HttpGet(url string) (status int, result string) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(bodyBytes)
}

func HttpGetWithToken(url, token string) (status int, result string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes)
}

func HttpGetCustomToken(url, token string) (status int, result string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes)
}

func HttpGetCookieWithToken(url, token string) (status int, result string, cookies []*http.Cookie) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		return 0, "", nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes), resp.Cookies() //resp header的set-cookie
}

func HttpGetWithCookie(url string, cookies []*http.Cookie) (status int, result string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	for i := 0; i < len(cookies); i++ {
		req.AddCookie(&http.Cookie{
			Name:    cookies[i].Name,
			Value:   cookies[i].Value,
			Path:    cookies[i].Path,
			Expires: cookies[i].Expires,
			Domain:  cookies[i].Domain,
		})
	}
	//fmt.Println(req)
	resp, err := client.Do(req)

	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(bodyBytes)
}

func HttpPostWithCookie(url string, data string, cookies []*http.Cookie) (status int, result string) {
	body := bytes.NewBufferString(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")

	for i := 0; i < len(cookies); i++ {
		req.AddCookie(&http.Cookie{
			Name:    cookies[i].Name,
			Value:   cookies[i].Value,
			Path:    cookies[i].Path,
			Expires: cookies[i].Expires,
			Domain:  cookies[i].Domain,
		})
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(bodyBytes)
}

func NewPutFileRequest(url, token, filePath, partNum string) (status int, result string) {

	client := &http.Client{}

	file, err := os.Open(filePath)
	if err != nil {
		return 1, ""
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return 2, ""
	}

	fi, err := file.Stat()
	if err != nil {
		return 3, ""
	}

	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)
	partWriter, err := multipartWriter.CreateFormFile("file", fi.Name())
	if err != nil {
		return 4, ""
	}
	partWriter.Write(fileContents)

	fieldWriter, err := multipartWriter.CreateFormField("part_number")
	if err != nil {
		return 5, ""
	}

	fieldWriter.Write([]byte(partNum))

	if err = multipartWriter.Close(); err != nil {
		return 6, ""
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return 7, ""
	}

	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, string(bodyBytes)
}
