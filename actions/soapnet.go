package actions

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var client = http.Client{}
var segment = "/vapix/services"

// SoapDo does the soap request.
func SoapDo(ctx context.Context, username string, password string, address string, soapAction string, soapEnvelopeRequest interface{}, soapEnvelopeResponse interface{}) error {
	data, err := xml.Marshal(soapEnvelopeRequest)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s%s", address, segment)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Add("Content-Type", "text/xml")
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyFromCamera, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed with %s, statusCode: %d, body: %s", url, resp.StatusCode, string(bodyFromCamera))
	}

	if soapEnvelopeResponse != nil {
		err = xml.Unmarshal(bodyFromCamera, soapEnvelopeResponse)
		if err != nil {
			return err
		}
	}

	return nil
}
