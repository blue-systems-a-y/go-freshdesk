package freshservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Client is the freshservice client
type Client struct {
	domain string
	apiKey string
	hc     *http.Client
}

//NewClient returns a new client
func NewClient(domain string, apiKey string) *Client {
	return &Client{
		domain: domain,
		apiKey: apiKey,
		hc:     &http.Client{},
	}
}

//ReadObject reads an object with GET method
func (c *Client) ReadObject(path string, obj interface{}) error {
	log.Printf("freshservice: reading object from %v", path)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}
	res, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = responseOK(res)
	if err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(obj)
}

//WriteObject writes object with the provided method
func (c *Client) WriteObject(path string, method string, in interface{}, out interface{}) error {
	log.Printf("freshservice: writing (%v) object %v to %v", method, in, path)
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := c.newRequest(method, path, data)
	if err != nil {
		return err
	}
	res, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = responseOK(res)
	if err != nil {
		return err
	}
	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}

	return nil
}

func (c *Client) newRequest(method string, path string, data []byte) (*http.Request, error) {
	url := fmt.Sprintf("https://%s.freshservice.com%s", c.domain, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.apiKey, "X")
	return req, nil
}

//responseOK returns nil if the response was ok or the error
func responseOK(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusCreated:
		return nil
	case http.StatusNoContent:
		return nil
	}

	data, err := ioutil.ReadAll(res.Body)
	return fmt.Errorf("bad status: %v. %v (read error: %v)", res.Status, string(data), err)
}

//GetCredentials of fresh service client from environment varibles
func GetCredentials() (string, string) {

	domain := os.Getenv("FS_DOMAIN")
	apiKey := os.Getenv("FS_KEY")

	return domain, apiKey
}
