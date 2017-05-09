package memex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
)

const serverURL = "http://localhost:5000/api/v1"

// Spaces object
type Spaces struct {
	verbose            bool
	serviceAccessToken string
	httpClient         *http.Client
}

// NewSpaces creates new spaces object
func NewSpaces(serviceAccessToken string) (*Spaces, error) {
	spaces := &Spaces{
		verbose:            true,
		serviceAccessToken: serviceAccessToken,
		httpClient:         &http.Client{},
	}
	return spaces, nil
}

func (spaces *Spaces) perform(method string, path string, body []byte, responseObject interface{}) (*http.Response, error) {
	endpointURL := fmt.Sprintf("%v%v", serverURL, path)
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
	request.Header.Add("X-Service-Access-Token", spaces.serviceAccessToken)
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
