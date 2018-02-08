package actions

import (
	"context"
	"os"
	"testing"
)

var (
	username = os.Getenv("axisevents-username")
	password = os.Getenv("axisevents-password")
	address  = os.Getenv("axisevents-address")
)

func TestGetActionConfigurationsLongRunning(t *testing.T) {
	verifyEnvVariables(t)
	handler := NewActionHandler(context.TODO(), username, password, address)
	resp, err := handler.GetActionConfigurations()
	if err != nil {
		t.Fatal(err)
	}
	numberOfActions := len(resp)
	if numberOfActions == 0 {
		t.Error("Should get more than 0 actionConfigurations")
	}
}

func TestGetActionTemplatesLongRunning(t *testing.T) {
	verifyEnvVariables(t)
	handler := NewActionHandler(context.TODO(), username, password, address)
	resp, err := handler.GetActionTemplates()
	if err != nil {
		t.Fatal(err)
	}
	numberOfActions := len(resp)
	if numberOfActions == 0 {
		t.Error("Should get more than 0 actionTemplates")
	}

	if len(resp[0].TemplateToken) == 0 {
		t.Error("Should get TemplateToken.")
	}

	if len(resp[0].Parameters[0].Name) == 0 {
		t.Error("Should get parameter name.")
	}
}

func TestGetRecipientTemplatesLongRunning(t *testing.T) {
	verifyEnvVariables(t)
	handler := NewActionHandler(context.TODO(), username, password, address)
	resp, err := handler.GetRecipientTemplates()
	if err != nil {
		t.Fatal(err)
	}
	numberOfRecipients := len(resp)
	if numberOfRecipients == 0 {
		t.Error("Should get more than 0 recipentTemplates")
	}

	if len(resp[0].Parameters) == 0 {
		t.Error("Should get more than 0 parameters.")
	}

	if len(resp[0].TemplateToken) == 0 {
		t.Error("Should get TemplateToken.")
	}

	if len(resp[0].Parameters[0].Name) == 0 {
		t.Error("Should get parameter name.")
	}
}

func TestGetActionRulesLongRunning(t *testing.T) {
	verifyEnvVariables(t)
	handler := NewActionHandler(context.TODO(), username, password, address)
	resp, err := handler.GetActionRules()
	if err != nil {
		t.Fatal(err)
	}
	numberOfActions := len(resp)
	if numberOfActions == 0 {
		t.Error("Should get more than 0 actionRules")
	}
}

func TestAddGetRemoveActions(t *testing.T) {
	verifyEnvVariables(t)
	// config := newTestAddConfiguration("HubbaIsTesting")
	// SoapDo(username, password, address, "http://www.axis.com/vapix/ws/action1/AddActionConfiguration", config)
	configName := "hubbatest"
	templateToken := "com.axis.action.fixed.recording.storage"
	ruleName := "bubbatest"

	handler := NewActionHandler(context.TODO(), username, password, address)

	configID, err := handler.AddActionConfigurations(configName, templateToken, getRecordingParameters())
	if err != nil {
		t.Fatalf("Failed to created ActionConfiguration, err: %s", err.Error())
	}

	assertActionConfiguration(t, handler, configID, configName, templateToken, true)

	topixExpression := NewTopicExpression("tns1:RuleEngine/tnsaxis:VideoMotionDetection/motion")
	messageContent := NewMessageContent(`boolean(//SimpleItem[@Name="active" and @Value="1"])`)
	startEvent := NewStartEvent(topixExpression, messageContent)

	ruleID, err := handler.AddActionRule(ruleName, false, startEvent, configID)
	if err != nil {
		t.Fatalf("Failed to created ActionRule, err: %s", err.Error())
	}

	assertActionRule(t, handler, ruleID, ruleName, configID, true)

	err = handler.RemoveActionRule(ruleID)
	if err != nil {
		t.Fatalf("Failed to remove ActionRules, err: %s", err.Error())
	}

	assertActionRule(t, handler, ruleID, ruleName, configID, false)

	err = handler.RemoveActionConfigurations(configID)
	if err != nil {
		t.Fatalf("Failed to remove ActionConfiguration, err: %s", err.Error())
	}

	assertActionConfiguration(t, handler, configID, configName, templateToken, false)
}

func assertActionConfiguration(t *testing.T, handler ActionHandler, configID int, configName string, templateToken string, expectedToBeFound bool) {
	configs, err := handler.GetActionConfigurations()
	if err != nil {
		t.Fatalf("Failed to list ActionConfigurations, err: %s", err.Error())
	}
	foundConfig := false
	for _, c := range configs {
		if c.Name == configName && c.TemplateToken == templateToken && c.ConfigurationID == configID {
			foundConfig = true
		}
	}

	if foundConfig != expectedToBeFound {
		t.Errorf("Should find actionconfiguration %b but was %b %+v", foundConfig, expectedToBeFound, configs)
	}
}

func assertActionRule(t *testing.T, handler ActionHandler, ruleID int, ruleName string, configID int, expectedToBeFound bool) {
	rules, err := handler.GetActionRules()
	if err != nil {
		t.Fatalf("Failed to list ActionRules, err: %s", err.Error())
	}
	foundRule := false
	for _, r := range rules {
		if r.Name == ruleName && r.Enabled == false && r.RuleID == ruleID && r.PrimaryAction == configID {
			foundRule = true
		}
	}
	if foundRule != expectedToBeFound {
		t.Errorf("Should find actionrule %b but was %b %+v", foundRule, expectedToBeFound, rules)
	}
}

func getRecordingParameters() []Parameter {
	parameter1 := Parameter{
		Name:  "stream_options",
		Value: "streamprofile=Quality;camera=1",
	}
	parameter2 := Parameter{
		Name:  "pre_duration",
		Value: "1000",
	}
	parameter3 := Parameter{
		Name:  "post_duration",
		Value: "1000",
	}
	parameter4 := Parameter{
		Name:  "storage_id",
		Value: "SD_DISK",
	}
	return []Parameter{parameter1, parameter2, parameter3, parameter4}
}

func newTestAddActionRule(name string, actionConfigurationID int) SoapEnvelope {
	topixExpression := NewTopicExpression("tns1:RuleEngine/tnsaxis:VideoMotionDetection/motion")
	messageContent := NewMessageContent("boolean(//SimpleItem[@Name='active' and @Value='1'])")
	startEvent := NewStartEvent(topixExpression, messageContent)

	actionRule := NewActionRule(name, true, startEvent.TopicExpression, startEvent.MessageContent, actionConfigurationID)
	addActionRule := NewAddActionRule(actionRule)
	body := NewSoapBody(addActionRule)
	return NewSoapEnvelope(body)
}

func verifyEnvVariables(t *testing.T) {
	if len(address) == 0 {
		t.Fatal("You need to set env variable axisevents-address")
	}

	if len(username) == 0 {
		t.Fatal("You need to set env variable axisevents-username")
	}

	if len(password) == 0 {
		t.Fatal("You need to set env variable axisevents-password")
	}
}
