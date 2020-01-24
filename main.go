package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

//
// JSON struct generated from:
//		https://mholt.github.io/json-to-go/
//
type macAddrResponse struct {
	VendorDetails struct {
		Oui            string `json:"oui"`
		IsPrivate      bool   `json:"isPrivate"`
		CompanyName    string `json:"companyName"`
		CompanyAddress string `json:"companyAddress"`
		CountryCode    string `json:"countryCode"`
	} `json:"vendorDetails"`
	BlockDetails struct {
		BlockFound          bool   `json:"blockFound"`
		BorderLeft          string `json:"borderLeft"`
		BorderRight         string `json:"borderRight"`
		BlockSize           int    `json:"blockSize"`
		AssignmentBlockSize string `json:"assignmentBlockSize"`
		DateCreated         string `json:"dateCreated"`
		DateUpdated         string `json:"dateUpdated"`
	} `json:"blockDetails"`
	MacAddressDetails struct {
		SearchTerm         string   `json:"searchTerm"`
		IsValid            bool     `json:"isValid"`
		VirtualMachine     string   `json:"virtualMachine"`
		Applications       []string `json:"applications"`
		TransmissionType   string   `json:"transmissionType"`
		AdministrationType string   `json:"administrationType"`
		WiresharkNotes     string   `json:"wiresharkNotes"`
		Comment            string   `json:"comment"`
	} `json:"macAddressDetails"`
}

type macAddrError struct {
	Error string `json:"error"`
}

// Associate this function as a method of type macAddrResponse
func (theObject macAddrResponse) parseResults(body []byte) (string, error) {
	responseText := ""
	responseErr := json.Unmarshal(body, &theObject)
	if responseErr == nil {
		responseText = "Owner Name: " + theObject.VendorDetails.CompanyName + "\n"
	}
	return responseText, responseErr
}

// Associate this function as a method of type macAddrError
func (theObject macAddrError) parseResults(body []byte) (string, error) {
	responseText := ""
	responseErr := json.Unmarshal(body, &theObject)
	if responseErr == nil {
		responseText = "Error: " + theObject.Error + "\n"
	}
	return responseText, responseErr
}

func main() {
	args := os.Args
	apiKey := ""
	baseURL := "https://api.macaddress.io/v1?output=json&search="

	if len(args) != 2 {
		fmt.Println("usage: command macAddr")
		fmt.Println("")
		//This should exit now because the command format is not correct
	}

	separator := "/"
	homeVariable := "HOME"
	if runtime.GOOS == "windows" {
		homeVariable = "HOMEPATH"
		separator = "\\"
	}
	homeDir := os.Getenv(homeVariable)
	apiKeyFile := homeDir + separator + ".macaddress"

	//apiKeyFile := "C:\\Users\\brian\\.macaddress"
	f, fOpenStatus := os.Open(apiKeyFile)
	if fOpenStatus != nil {
		panic(fOpenStatus)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		apiKey = line
	}
	//
	// Now we have the API key we can actually get the information from "macaddress.io"
	// We will need a pointer to an http Client
	//
	httpClient := &http.Client{}

	httpRequest, _ := http.NewRequest("GET", baseURL+args[1], nil)
	httpRequest.Header.Set("X-Authentication-Token", apiKey)
	httpResp, httpErr := httpClient.Do(httpRequest)

	if httpErr != nil {
		//handling the error
		log.Fatal(httpErr)
	}

	defer httpResp.Body.Close()

	body, readErr := ioutil.ReadAll(httpResp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var responseErr error
	var responseText string
	//
	// The parsing of the response is attached to the results. Since Go does not have
	// runtime typing, it is done at compile time. This results in having to know what to expect
	// and then process accordingly.
	//
	// We know the good results comes with a 200 response. We are making the assumption that
	// all other reponses result in an error. But just in case we are returning the error
	// indicator and processing a fatal error after returning instead of within the method.
	//
	if httpResp.StatusCode == 200 {
		goodResponseStruct := macAddrResponse{}
		responseText, responseErr = goodResponseStruct.parseResults(body)
	} else {
		errorResponseStruct := macAddrError{}
		responseText, responseErr = errorResponseStruct.parseResults(body)
	}

	if responseErr != nil {
		log.Fatal(responseErr)
	}

	fmt.Printf(responseText)
}
