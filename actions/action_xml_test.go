package actions

import (
	"encoding/xml"
	"testing"
)

// TestGetConfigurations checks that a GetActionConfigurations cam be marshalled as expected.
func TestGetConfigurations(t *testing.T) {
	var response GetActionConfigurationsSoapEnvelope
	err := xml.Unmarshal([]byte(GetConfigurationResponseExpected), &response)
	if err != nil {
		t.Fatal("Failed to umarshal GetActionConfigurationsSoapEnvelope", err)
	}

	numberOfConfigurations := len(response.GetActionConfigurationsSoapBody.GetActionConfigurationsResponse.ActionConfigurationsResponse)
	if numberOfConfigurations != 5 {
		t.Fatalf("Should have 5 ActionConfigurations but got %d", numberOfConfigurations)
	}

	action1 := response.GetActionConfigurationsSoapBody.GetActionConfigurationsResponse.ActionConfigurationsResponse[1]
	if action1.ConfigurationID != 2 {
		t.Errorf("ConfigurationID should be %v but was %v", 2, action1.ConfigurationID)
	}

	if action1.Name != "ACC_MotionAction" {
		t.Errorf("Name should be %v but was %v", 2, action1.Name)
	}

	if action1.TemplateToken != "com.axis.action.unlimited.recording.storage" {
		t.Errorf("TemplateToken should be %v but was %v", "com.axis.action.unlimited.recording.storage", action1.TemplateToken)
	}

	if len(action1.Parameters) != 4 {
		t.Fatalf("Should have 4 parameters but got %d", len(action1.Parameters))
	}

	parameter4 := action1.Parameters[3]

	if parameter4.Name != "stream_options" {
		t.Errorf("stream_options should be %v but was %v", "stream_options", parameter4.Name)
	}

	if parameter4.Value != "streamprofile=ACC_High" {
		t.Errorf("stream_options should be %v but was %v", "streamprofile=ACC_High", parameter4.Value)
	}
}

func TestAddConfiguration(t *testing.T) {
	expectedConfigurationID := 5
	var response AddActionConfigurationSoapEnvelope
	err := xml.Unmarshal([]byte(AddConfigurationResponseExpected), &response)
	if err != nil {
		t.Fatal("Failed to umarshal AddActionConfigurationSoapEnvelope", err)
	}
	actualConfigID := response.AddActionConfigurationSoapBody.AddActionConfigurationResponse.ConfigurationID
	if actualConfigID != expectedConfigurationID {
		t.Errorf("Should configurationId %d but got %v", expectedConfigurationID, actualConfigID)
	}
}

func TestAddActionRule(t *testing.T) {
	expectedRuleID := 5
	var response AddActionRuleSoapEnvelope
	err := xml.Unmarshal([]byte(AddActionRuleResponseExpected), &response)
	if err != nil {
		t.Fatal("Failed to umarshal AddActionRuleSoapEnvelope", err)
	}
	actualRuleID := response.AddActionRulenSoapBody.AddActionRuleResponse.RuleID
	if actualRuleID != expectedRuleID {
		t.Errorf("Should get ruleID %d but got %v", expectedRuleID, actualRuleID)
	}
}

var (
	AddActionRuleResponseExpected = `<SOAP-ENV:Envelope 
    xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" 
    xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:c14n="http://www.w3.org/2001/10/xml-exc-c14n#" 
    xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd" 
    xmlns:ds="http://www.w3.org/2000/09/xmldsig#" 
    xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" 
    xmlns:wsa5="http://www.w3.org/2005/08/addressing" 
    xmlns:xmime="http://tempuri.org/xmime.xsd" 
    xmlns:xop="http://www.w3.org/2004/08/xop/include" 
    xmlns:wsrfbf="http://docs.oasis-open.org/wsrf/bf-2" 
    xmlns:wstop="http://docs.oasis-open.org/wsn/t-1" 
    xmlns:tt="http://www.onvif.org/ver10/schema" 
    xmlns:acert="http://www.axis.com/vapix/ws/cert" 
    xmlns:wsrfr="http://docs.oasis-open.org/wsrf/r-2" 
    xmlns:aa="http://www.axis.com/vapix/ws/action1" 
    xmlns:acertificates="http://www.axis.com/vapix/ws/certificates" 
    xmlns:aev="http://www.axis.com/vapix/ws/event1" 
    xmlns:ali1="http://www.axis.com/vapix/ws/light/CommonBinding" 
    xmlns:ali2="http://www.axis.com/vapix/ws/light/IntensityBinding" 
    xmlns:ali3="http://www.axis.com/vapix/ws/light/AngleOfIlluminationBinding" 
    xmlns:ali4="http://www.axis.com/vapix/ws/light/DayNightSynchronizeBinding" 
    xmlns:ali="http://www.axis.com/vapix/ws/light" 
    xmlns:apc="http://www.axis.com/vapix/ws/panopsiscalibration1" 
    xmlns:arth="http://www.axis.com/vapix/ws/recordedtour1" 
    xmlns:aweb="http://www.axis.com/vapix/ws/webserver" 
    xmlns:tan1="http://www.onvif.org/ver20/analytics/wsdl/RuleEngineBinding" 
    xmlns:tan2="http://www.onvif.org/ver20/analytics/wsdl/AnalyticsEngineBinding" 
    xmlns:tan="http://www.onvif.org/ver20/analytics/wsdl" 
    xmlns:tds="http://www.onvif.org/ver10/device/wsdl" 
    xmlns:tev1="http://www.onvif.org/ver10/events/wsdl/NotificationProducerBinding" 
    xmlns:tev2="http://www.onvif.org/ver10/events/wsdl/EventBinding" 
    xmlns:tev3="http://www.onvif.org/ver10/events/wsdl/SubscriptionManagerBinding" 
    xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2" 
    xmlns:tev4="http://www.onvif.org/ver10/events/wsdl/PullPointSubscriptionBinding" 
    xmlns:tev="http://www.onvif.org/ver10/events/wsdl" 
    xmlns:timg="http://www.onvif.org/ver20/imaging/wsdl" 
    xmlns:tptz="http://www.onvif.org/ver20/ptz/wsdl" 
    xmlns:trt="http://www.onvif.org/ver10/media/wsdl" 
    xmlns:ter="http://www.onvif.org/ver10/error" 
    xmlns:tns1="http://www.onvif.org/ver10/topics" 
    xmlns:tnsaxis="http://www.axis.com/2009/event/topics">
    <SOAP-ENV:Body>
        <aa:AddActionRuleResponse>
            <aa:RuleID>5</aa:RuleID>
        </aa:AddActionRuleResponse>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
	AddConfigurationResponseExpected = `<SOAP-ENV:Envelope 
    xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" 
    xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:c14n="http://www.w3.org/2001/10/xml-exc-c14n#" 
    xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd" 
    xmlns:ds="http://www.w3.org/2000/09/xmldsig#" 
    xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" 
    xmlns:wsa5="http://www.w3.org/2005/08/addressing" 
    xmlns:xmime="http://tempuri.org/xmime.xsd" 
    xmlns:xop="http://www.w3.org/2004/08/xop/include" 
    xmlns:wsrfbf="http://docs.oasis-open.org/wsrf/bf-2" 
    xmlns:wstop="http://docs.oasis-open.org/wsn/t-1" 
    xmlns:tt="http://www.onvif.org/ver10/schema" 
    xmlns:acert="http://www.axis.com/vapix/ws/cert" 
    xmlns:wsrfr="http://docs.oasis-open.org/wsrf/r-2" 
    xmlns:aa="http://www.axis.com/vapix/ws/action1" 
    xmlns:acertificates="http://www.axis.com/vapix/ws/certificates" 
    xmlns:aev="http://www.axis.com/vapix/ws/event1" 
    xmlns:ali1="http://www.axis.com/vapix/ws/light/CommonBinding" 
    xmlns:ali2="http://www.axis.com/vapix/ws/light/IntensityBinding" 
    xmlns:ali3="http://www.axis.com/vapix/ws/light/AngleOfIlluminationBinding" 
    xmlns:ali4="http://www.axis.com/vapix/ws/light/DayNightSynchronizeBinding" 
    xmlns:ali="http://www.axis.com/vapix/ws/light" 
    xmlns:apc="http://www.axis.com/vapix/ws/panopsiscalibration1" 
    xmlns:arth="http://www.axis.com/vapix/ws/recordedtour1" 
    xmlns:aweb="http://www.axis.com/vapix/ws/webserver" 
    xmlns:tan1="http://www.onvif.org/ver20/analytics/wsdl/RuleEngineBinding" 
    xmlns:tan2="http://www.onvif.org/ver20/analytics/wsdl/AnalyticsEngineBinding" 
    xmlns:tan="http://www.onvif.org/ver20/analytics/wsdl" 
    xmlns:tds="http://www.onvif.org/ver10/device/wsdl" 
    xmlns:tev1="http://www.onvif.org/ver10/events/wsdl/NotificationProducerBinding" 
    xmlns:tev2="http://www.onvif.org/ver10/events/wsdl/EventBinding" 
    xmlns:tev3="http://www.onvif.org/ver10/events/wsdl/SubscriptionManagerBinding" 
    xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2" 
    xmlns:tev4="http://www.onvif.org/ver10/events/wsdl/PullPointSubscriptionBinding" 
    xmlns:tev="http://www.onvif.org/ver10/events/wsdl" 
    xmlns:timg="http://www.onvif.org/ver20/imaging/wsdl" 
    xmlns:tptz="http://www.onvif.org/ver20/ptz/wsdl" 
    xmlns:trt="http://www.onvif.org/ver10/media/wsdl" 
    xmlns:ter="http://www.onvif.org/ver10/error" 
    xmlns:tns1="http://www.onvif.org/ver10/topics" 
    xmlns:tnsaxis="http://www.axis.com/2009/event/topics">
    <SOAP-ENV:Body>
        <aa:AddActionConfigurationResponse>
            <aa:ConfigurationID>5</aa:ConfigurationID>
        </aa:AddActionConfigurationResponse>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

	GetConfigurationResponseExpected = `<SOAP-ENV:Envelope 
    xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" 
    xmlns:SOAP-ENC="http://www.w3.org/2003/05/soap-encoding" 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
    xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
    xmlns:c14n="http://www.w3.org/2001/10/xml-exc-c14n#" 
    xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd" 
    xmlns:ds="http://www.w3.org/2000/09/xmldsig#" 
    xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" 
    xmlns:wsa5="http://www.w3.org/2005/08/addressing" 
    xmlns:xmime="http://tempuri.org/xmime.xsd" 
    xmlns:xop="http://www.w3.org/2004/08/xop/include" 
    xmlns:wsrfbf="http://docs.oasis-open.org/wsrf/bf-2" 
    xmlns:wstop="http://docs.oasis-open.org/wsn/t-1" 
    xmlns:tt="http://www.onvif.org/ver10/schema" 
    xmlns:acert="http://www.axis.com/vapix/ws/cert" 
    xmlns:wsrfr="http://docs.oasis-open.org/wsrf/r-2" 
    xmlns:aa="http://www.axis.com/vapix/ws/action1" 
    xmlns:acertificates="http://www.axis.com/vapix/ws/certificates" 
    xmlns:aev="http://www.axis.com/vapix/ws/event1" 
    xmlns:ali1="http://www.axis.com/vapix/ws/light/CommonBinding" 
    xmlns:ali2="http://www.axis.com/vapix/ws/light/IntensityBinding" 
    xmlns:ali3="http://www.axis.com/vapix/ws/light/AngleOfIlluminationBinding" 
    xmlns:ali4="http://www.axis.com/vapix/ws/light/DayNightSynchronizeBinding" 
    xmlns:ali="http://www.axis.com/vapix/ws/light" 
    xmlns:apc="http://www.axis.com/vapix/ws/panopsiscalibration1" 
    xmlns:arth="http://www.axis.com/vapix/ws/recordedtour1" 
    xmlns:aweb="http://www.axis.com/vapix/ws/webserver" 
    xmlns:tan1="http://www.onvif.org/ver20/analytics/wsdl/RuleEngineBinding" 
    xmlns:tan2="http://www.onvif.org/ver20/analytics/wsdl/AnalyticsEngineBinding" 
    xmlns:tan="http://www.onvif.org/ver20/analytics/wsdl" 
    xmlns:tds="http://www.onvif.org/ver10/device/wsdl" 
    xmlns:tev1="http://www.onvif.org/ver10/events/wsdl/NotificationProducerBinding" 
    xmlns:tev2="http://www.onvif.org/ver10/events/wsdl/EventBinding" 
    xmlns:tev3="http://www.onvif.org/ver10/events/wsdl/SubscriptionManagerBinding" 
    xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2" 
    xmlns:tev4="http://www.onvif.org/ver10/events/wsdl/PullPointSubscriptionBinding" 
    xmlns:tev="http://www.onvif.org/ver10/events/wsdl" 
    xmlns:timg="http://www.onvif.org/ver20/imaging/wsdl" 
    xmlns:tptz="http://www.onvif.org/ver20/ptz/wsdl" 
    xmlns:trt="http://www.onvif.org/ver10/media/wsdl" 
    xmlns:ter="http://www.onvif.org/ver10/error" 
    xmlns:tns1="http://www.onvif.org/ver10/topics" 
    xmlns:tnsaxis="http://www.axis.com/2009/event/topics">
    <SOAP-ENV:Body>
        <aa:GetActionConfigurationsResponse>
            <aa:ActionConfigurations>
                <aa:ActionConfiguration>
                    <aa:ConfigurationID>1</aa:ConfigurationID>
                    <aa:Name>Activate LED</aa:Name>
                    <aa:TemplateToken>com.axis.action.unlimited.ledcontrol</aa:TemplateToken>
                    <aa:Parameters>
                        <aa:Parameter Value="250" Name="interval"></aa:Parameter>
                        <aa:Parameter Value="statusled" Name="led"></aa:Parameter>
                        <aa:Parameter Value="amber,none" Name="color"></aa:Parameter>
                    </aa:Parameters>
                </aa:ActionConfiguration>
                <aa:ActionConfiguration>
                    <aa:ConfigurationID>2</aa:ConfigurationID>
                    <aa:Name>ACC_MotionAction</aa:Name>
                    <aa:TemplateToken>com.axis.action.unlimited.recording.storage</aa:TemplateToken>
                    <aa:Parameters>
                        <aa:Parameter Value="30000" Name="post_duration"></aa:Parameter>
                        <aa:Parameter Value="5000" Name="pre_duration"></aa:Parameter>
                        <aa:Parameter Value="SD_DISK" Name="storage_id"></aa:Parameter>
                        <aa:Parameter Value="streamprofile=ACC_High" Name="stream_options"></aa:Parameter>
                    </aa:Parameters>
                </aa:ActionConfiguration>
                <aa:ActionConfiguration>
                    <aa:ConfigurationID>3</aa:ConfigurationID>
                    <aa:Name>ACC_ContinuousAction</aa:Name>
                    <aa:TemplateToken>com.axis.action.unlimited.recording.storage</aa:TemplateToken>
                    <aa:Parameters>
                        <aa:Parameter Value="0" Name="post_duration"></aa:Parameter>
                        <aa:Parameter Value="0" Name="pre_duration"></aa:Parameter>
                        <aa:Parameter Value="SD_DISK" Name="storage_id"></aa:Parameter>
                        <aa:Parameter Value="streamprofile=ACC_Low" Name="stream_options"></aa:Parameter>
                    </aa:Parameters>
                </aa:ActionConfiguration>
                <aa:ActionConfiguration>
                    <aa:ConfigurationID>4</aa:ConfigurationID>
                    <aa:Name>ACC_LowAction</aa:Name>
                    <aa:TemplateToken>com.axis.action.unlimited.recording.storage</aa:TemplateToken>
                    <aa:Parameters>
                        <aa:Parameter Value="30000" Name="post_duration"></aa:Parameter>
                        <aa:Parameter Value="5000" Name="pre_duration"></aa:Parameter>
                        <aa:Parameter Value="SD_DISK" Name="storage_id"></aa:Parameter>
                        <aa:Parameter Value="streamprofile=ACC_Low" Name="stream_options"></aa:Parameter>
                    </aa:Parameters>
                </aa:ActionConfiguration>
                <aa:ActionConfiguration>
                    <aa:ConfigurationID>5</aa:ConfigurationID>
                    <aa:Name>Record Video</aa:Name>
                    <aa:TemplateToken>com.axis.action.fixed.recording.storage</aa:TemplateToken>
                    <aa:Parameters>
                        <aa:Parameter Value="1000" Name="post_duration"></aa:Parameter>
                        <aa:Parameter Value="1000" Name="pre_duration"></aa:Parameter>
                        <aa:Parameter Value="SD_DISK" Name="storage_id"></aa:Parameter>
                        <aa:Parameter Value="&amp;streamprofile=Quality&amp;camera=1" Name="stream_options"></aa:Parameter>
                    </aa:Parameters>
                </aa:ActionConfiguration>
            </aa:ActionConfigurations>
        </aa:GetActionConfigurationsResponse>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
)
