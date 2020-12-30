package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct{}

func (r *Request) Get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := parseBody(response)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return body, nil
}

func (r *Request) Post(url string, body interface{}) (string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	parsedBody, err := parseBody(response)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return parsedBody, nil
}

func (r *Request) Delete(url string) error {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func parseBody(response *http.Response) (string, error) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func NewRequest() *Request {
	return &Request{}
}
