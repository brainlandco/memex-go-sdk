package memex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Spaces object
type Spaces struct {
	verbose    bool
	appToken   string
	userToken  string
	serverURL  string
	httpClient *http.Client
}

// NewSpaces creates new spaces object
func NewSpaces() (*Spaces, error) {
	spaces := &Spaces{
		verbose:    true,
		appToken:   "",
		userToken:  "",
		serverURL:  serverURL(Production),
		httpClient: &http.Client{},
	}
	return spaces, nil
}

// SetAppToken sets app/client token
func (spaces *Spaces) SetAppToken(token string) {
	spaces.appToken = token
}

// SetUserToken sets user token
func (spaces *Spaces) SetUserToken(token string) {
	spaces.userToken = token
}

// SetEnvironment sets environment of service Production/Stage/Sandbox/Local
func (spaces *Spaces) SetEnvironment(environment Environment) {
	spaces.serverURL = serverURL(environment)
}

func serverURL(environment Environment) string {
	switch environment {
	case Production:
		return "https://mmx-spaces-api-prod.herokuapp.com"
	case Stage:
		return "https://mmx-spaces-api-stage.herokuapp.com"
	case Local:
		return "http://localhost:5000"
	}
	return ""
}

func (spaces *Spaces) perform(method string, path string, body []byte, responseObject interface{}) (*http.Response, error) {
	if spaces.appToken == "" {
		return nil, fmt.Errorf("Missing appToken, call SetAppToken(\"<Your app token>\")")
	}
	endpointURL := fmt.Sprintf("%v%v", spaces.serverURL, path)
	bodyReader := bytes.NewBuffer(body)
	if spaces.verbose {
		logrus.Print("==== REQUEST ====")
		logrus.Print(endpointURL)
		logrus.Print(string(body))
	}
	request, requestCreationError := http.NewRequest(method, endpointURL, bodyReader)
	if requestCreationError != nil {
		return nil, fmt.Errorf("Unable to create request: %v", requestCreationError.Error())
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-App-Token", spaces.appToken)
	if spaces.userToken != "" {
		request.Header.Add("X-User-Token", spaces.userToken)
	}
	response, fetchError := spaces.httpClient.Do(request)
	if fetchError != nil {
		return nil, fmt.Errorf("Unable to fetch url: %v", fetchError.Error())
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("Invalid response code: %v", response.StatusCode)
	}
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, fmt.Errorf("Cant read data")
	}
	parseError := json.Unmarshal(body, responseObject)
	if parseError != nil {
		s := string(body)
		fmt.Printf(s)
		return nil, fmt.Errorf("Cant parse response")
	}
	return response, nil
}
