package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Kianrad/spekeclient/models"
)

func RequestKeys(endpoint string, requestPayload models.CPIXRequest, requestHeaders map[string][]string) (*models.CPIXResponse, error) {

	requestBody, err := xml.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshaling request payload:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, err
	}

	req.Header = requestHeaders

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 status code:", resp.StatusCode)
		fmt.Println("Response body:", string(responseBody))
		return nil, err
	}

	var spekeResponse models.CPIXResponse
	err = xml.Unmarshal(responseBody, &spekeResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response XML:", err)
		return nil, err
	}
	return &spekeResponse, nil

}
