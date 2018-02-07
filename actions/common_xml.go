package actions

import "encoding/xml"

type SoapBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Body    interface{}
}

func NewSoapEnvelope(body SoapBody) SoapEnvelope {
	env := SoapEnvelope{
		Body:    body,
		XSI:     "http://www.w3.org/2001/XMLSchema-instance",
		XSD:     "http://www.w3.org/2001/XMLSchema",
		AA:      "http://www.axis.com/vapix/ws/action1",
		WSNT:    "http://docs.oasis-open.org/wsn/b-2",
		TNS1:    "http://www.onvif.org/ver10/topics",
		TNSAXIS: "http://www.axis.com/2009/event/topics",
		SOAP:    "http://www.w3.org/2003/05/soap-envelope",
	}
	return env
}

func NewSoapBody(body interface{}) SoapBody {
	return SoapBody{
		Body: body,
	}
}

type SoapEnvelope struct {
	Body    SoapBody
	XMLName xml.Name `xml:"soap:Envelope"`
	XSI     string   `xml:"xmlns:xsi,attr"`
	XSD     string   `xml:"xmlns:xsd,attr"`
	AA      string   `xml:"xmlns:aa,attr"`
	WSNT    string   `xml:"xmlns:wsnt,attr"`
	TNS1    string   `xml:"xmlns:tns1,attr"`
	TNSAXIS string   `xml:"xmlns:tnsaxis,attr"`
	SOAP    string   `xml:"xmlns:soap,attr"`
}

// type SoapBodyResponse struct {
// 	XMLName xml.Name `xml:"soap:Body"`
// 	Body    interface{}
// }

// type SoapEnvelopeResponse struct {
// 	Body    SoapBodyResonse
// 	XMLName xml.Name `xml:"Envelope"`
// }
