package client

import (
	"bytes"
	xml "encoding/xml"

	nxml "github.com/nbio/xml"

	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/Kianrad/spekeclient/models"
	modelsv2 "github.com/Kianrad/spekeclient/modelsv2"
)

func RequestKeys(endpoint string, requestPayload models.CPIXRequest, requestHeaders map[string][]string) (*models.CPIXResponse, error) {

	requestBody, err := xml.MarshalIndent(requestPayload, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling request payload:", err)
		return nil, err
	}

	requestBody = []byte(xml.Header + string(requestBody))
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
	decoder := nxml.NewDecoder(bytes.NewReader(responseBody))

	var spekeResponse models.CPIXResponse
	err = decoder.Decode(&spekeResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response XML:", err)
		return nil, err
	}
	return &spekeResponse, nil
}

func RequestV2Keys(endpoint string, requestPayload modelsv2.CPIXRequest, requestHeaders map[string][]string) (*modelsv2.CPIXResponse, error) {

	requestBody, err := xml.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshaling request payload:", err)
		return nil, err
	}

	requestBody = []byte(xml.Header + string(requestBody))
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, err
	}

	req.Header = requestHeaders
	req.Header.Set("X-Speke-Version", "2.0")

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
	decoder := nxml.NewDecoder(bytes.NewReader(responseBody))
	var spekeResponse modelsv2.CPIXResponse
	err = decoder.Decode(&spekeResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response XML:", err)
		return nil, err
	}
	return &spekeResponse, nil

}
