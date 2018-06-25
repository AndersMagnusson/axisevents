package actions

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

var segment = "/vapix/services"

// SoapDo does the soap request.
func SoapDo(ctx context.Context, username string, password string, address string, soapAction string, soapEnvelopeRequest interface{}, soapEnvelopeResponse interface{}) error {
	data, err := xml.Marshal(soapEnvelopeRequest)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s%s", address, segment)

	dr := dac.NewRequest(username, password, "POST", url, string(data))
	resp, err := dr.Execute()

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
