package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"log"
	"fmt"
//	"github.com/dinolupo/camunda-utility/pkg/utils"
)

const DefaultTimeoutSec = 60
const PackageVersion = "1.0.0"
const DefaultUserAgent = "CamundaClientGo/" + PackageVersion
const Protocol = "http://"
const DefaultHost = "localhost"
const DefaultPort = 8080
const DefaultEngineRest = "/engine-rest"
const DefaultEndpointUrl = "http://localhost:8080/engine-rest"

// ClientOptions a client options
type ClientOptions struct {
	UserAgent string
	Host      string
	Port      int
	Timeout   time.Duration
}

// Client a client for Camunda API
type Client struct {
	httpClient  *http.Client
	endpointUrl string
	userAgent   string
}

var ErrorNotFound = &Error{
	Type:    "NotFound",
	Message: "Not found",
}

// Error a custom error type
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Error error message
func (e *Error) Error() string {
	return e.Message
}

// NewClient create new instance Client
func NewClient(options ClientOptions) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: time.Second * DefaultTimeoutSec,
		},
		endpointUrl: "",
		userAgent:   DefaultUserAgent,
	}

	client.endpointUrl = fmt.Sprintf("%s%s:%d%s", Protocol, options.Host, options.Port, DefaultEngineRest)

	if options.UserAgent != "" {
		client.userAgent = options.UserAgent
	}

	if options.Timeout.Nanoseconds() != 0 {
		client.httpClient.Timeout = options.Timeout
	}

	return client
}

func (c *Client) doPostJson(path string, query map[string]string, v interface{}) (res *http.Response, err error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(v); err != nil {
		return nil, err
	}

	res, err = c.do(http.MethodPost, path, query, body, "application/json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) doPutJson(path string, query map[string]string, v interface{}) error {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(v); err != nil {
		return err
	}

	_, err := c.do(http.MethodPut, path, query, body, "application/json")
	return err
}

func (c *Client) doGet(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodGet, path, query, nil, "")
}

func (c *Client) doDelete(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodDelete, path, query, nil, "")
}

func (c *Client) doPost(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodPost, path, query, nil, "")
}

func (c *Client) doPut(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodPut, path, query, nil, "")
}

func (c *Client) do(method, path string, query map[string]string, body io.Reader, contentType string) (res *http.Response, err error) {
	url, err := c.buildUrl(path, query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	//req.SetBasicAuth(c.apiUser, c.apiPassword)

	log.Printf("%s %s\n", method, url)
	log.Println("HEADERS:")
	for k, v := range req.Header {
		log.Printf("\t%s: %s\n", k, v)
	}

	var requestBody []byte
	if body != nil {
		requestBody, err = ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}
		log.Println("BODY:")
		log.Println(string(requestBody))
	}

	res, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// log.Printf("Content-Length: %d\n", res.ContentLength)
	// responseBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// prettyResponse, err := utils.PrettyString(string(responseBody))
	// if err != nil {
	// 	return nil, err
	// }
	// log.Printf("RESPONSE BODY: %s\n", prettyResponse)

	// responseBody, err := ioutil.ReadAll(res.Body)
	// log.Printf("RESPONSE BODY: %s\n", string(responseBody))


	if err := c.checkResponse(res); err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) buildUrl(path string, query map[string]string) (string, error) {
	if len(query) == 0 {
		return c.endpointUrl + path, nil
	}
	url, err := url.Parse(c.endpointUrl + path)
	if err != nil {
		return "", err
	}

	q := url.Query()
	for k, v := range query {
		q.Set(k, v)
	}

	url.RawQuery = q.Encode()
	return url.String(), nil
}

func (c *Client) checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	defer res.Body.Close()

	if res.Header.Get("Content-Type") == "application/json" {
		if res.StatusCode == 404 {
			return ErrorNotFound
		}

		jsonErr := &Error{}
		err := json.NewDecoder(res.Body).Decode(jsonErr)
		if err != nil {
			return fmt.Errorf("response error with status code %d: failed unmarshal error response: %w", res.StatusCode, err)
		}

		return jsonErr
	}

	errText, err := ioutil.ReadAll(res.Body)
	if err == nil {
		return fmt.Errorf("response error with status code %d: %s", res.StatusCode, string(errText))
	}

	return fmt.Errorf("response error with status code %d", res.StatusCode)
}

func (c *Client) readJsonResponse(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
