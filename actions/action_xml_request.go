package actions

import "encoding/xml"

func NewGetActionConfigurationsRequest() GetActionConfigurationsRequest {
	return GetActionConfigurationsRequest{
		XMLNS: "http://www.axis.com/vapix/ws/action1",
	}
}

func NewAddActionConfigurationsRequest(name string, templateToken string, parameters []Parameter) SoapEnvelope {
	newcConfig := NewActionConfiguration{
		Name:          name,
		TemplateToken: templateToken,
		Parameters:    parameters,
	}

	config := AddActionConfiguration{
		XMLNS:                  "http://www.axis.com/vapix/ws/action1",
		NewActionConfiguration: newcConfig,
	}

	return NewSoapEnvelope(NewSoapBody(config))
}

func NewRemoveActionConfigurationsRequest(configurationID int) SoapEnvelope {
	config := RemoveActionConfiguration{ConfigurationID: configurationID, XMLNS: "http://www.axis.com/vapix/ws/action1"}
	return NewSoapEnvelope(NewSoapBody(config))
}

type GetActionConfigurationsRequest struct {
	XMLName xml.Name `xml:"aa:GetActionConfigurations"`
	XMLNS   string   `xml:"xmlns,attr"`
}

func NewGetActionTemplatesRequest() GetActionTemplatesRequest {
	return GetActionTemplatesRequest{
		XMLNS: "http://www.axis.com/vapix/ws/action1",
	}
}

type GetActionTemplatesRequest struct {
	XMLName xml.Name `xml:"aa:GetActionTemplates"`
	XMLNS   string   `xml:"xmlns,attr"`
}

func NewGetRecipientTemplatesRequest() GetRecipientTemplatesRequest {
	return GetRecipientTemplatesRequest{
		XMLNS: "http://www.axis.com/vapix/ws/action1",
	}
}

type GetRecipientTemplatesRequest struct {
	XMLName xml.Name `xml:"aa:GetRecipientTemplates"`
	XMLNS   string   `xml:"xmlns,attr"`
}

func NewGetActionRulesRequest() GetActionRulesRequest {
	return GetActionRulesRequest{
		XMLNS: "http://www.axis.com/vapix/ws/action1",
	}
}

type GetActionRulesRequest struct {
	XMLName xml.Name `xml:"aa:GetActionRules"`
	XMLNS   string   `xml:"xmlns,attr"`
}

type NewActionConfiguration struct {
	TemplateToken string      `xml:"TemplateToken`
	Name          string      `xml:"Name`
	Parameters    []Parameter `xml:"Parameters>Parameter"`
}

type Parameter struct {
	Name  string      `xml:"Name,attr"`
	Value interface{} `xml:"Value,attr"`
}

type AddActionConfiguration struct {
	XMLNS                  string                 `xml:"xmlns,attr"`
	NewActionConfiguration NewActionConfiguration `xml:"NewActionConfiguration"`
	XMLName                xml.Name               `xml:"aa:AddActionConfiguration"`
}

type RemoveActionConfiguration struct {
	XMLNS           string   `xml:"xmlns,attr"`
	ConfigurationID int      `xml:"ConfigurationID"`
	XMLName         xml.Name `xml:"aa:RemoveActionConfiguration"`
}

func NewParameter(name string, value string) Parameter {
	return Parameter{Name: name, Value: value}
}

func NewAddActionRule(actionRule ActionRule) AddActionRule {
	return AddActionRule{
		ActionRule: actionRule,
		XmlNS:      "http://www.axis.com/vapix/ws/action1",
	}
}

func NewAddActionConditionRule(actionRule ActionConditionRule) AddActionConditionRule {
	return AddActionConditionRule{
		ActionRule: actionRule,
		XmlNS:      "http://www.axis.com/vapix/ws/action1",
	}
}

type AddActionRule struct {
	XmlNS      string     `xml:"xmlns,attr"`
	ActionRule ActionRule `xml:"NewActionRule"`
	XMLName    xml.Name   `xml:"aa:AddActionRule"`
}

type AddActionConditionRule struct {
	XmlNS      string              `xml:"xmlns,attr"`
	ActionRule ActionConditionRule `xml:"NewActionRule"`
	XMLName    xml.Name            `xml:"aa:AddActionRule"`
}

func NewStartEvent(topicExpression TopicExpression, messageContent MessageContent) StartEvent {
	return StartEvent{
		TopicExpression: topicExpression,
		MessageContent:  messageContent,
	}
}

type StartEvent struct {
	TopicExpression TopicExpression `xml:""`
	MessageContent  MessageContent  `xml:""`
}

func NewTopicExpression(value string) TopicExpression {
	return TopicExpression{
		Expression: value,
		Dialect:    "http://www.onvif.org/ver10/tev/topicExpression/ConcreteSet",
	}
}

func NewTopicOasisTopicExpression(value string) TopicExpression {
	return TopicExpression{
		Expression: value,
		Dialect:    "http://docs.oasis-open.org/wsn/t-1/TopicExpression/Concrete",
	}
}

type ConditionRequest struct {
	TopicExpression TopicExpression `xml:"wsnt:TopicExpression"`
	MessageContent  MessageContent  `xml:"wsnt:MessageContent"`
}

type TopicExpression struct {
	Expression string `xml:",chardata"`
	Dialect    string `xml:"Dialect,attr"`
}

func NewMessageContent(value string) MessageContent {
	return MessageContent{
		Content: value,
		Dialect: "http://www.onvif.org/ver10/tev/messageContentFilter/ItemFilter",
	}
}

type MessageContent struct {
	Content string `xml:",chardata"`
	Dialect string `xml:"Dialect,attr"`
}

func NewActionRule(name string, enabled bool, topic TopicExpression, content MessageContent, primaryActionID int) ActionRule {
	return ActionRule{
		Name:            name,
		Enabled:         enabled,
		TopicExpression: topic,
		MessageContent:  content,
		//StartEvent:    startEvent,
		PrimaryAction: primaryActionID,
	}
}

type ActionRule struct {
	Name            string          `xml:"Name`
	Enabled         bool            `xml:"Enabled`
	TopicExpression TopicExpression `xml:"StartEvent>wsnt:TopicExpression"`
	MessageContent  MessageContent  `xml:"StartEvent>wsnt:MessageContent"`
	PrimaryAction   int             `xml:"PrimaryAction`
}

func NewActionConditionRule(
	name string, enabled bool, conditions []ConditionRequest, primaryActionID int) ActionConditionRule {
	return ActionConditionRule{
		Name:          name,
		Enabled:       enabled,
		Conditions:    conditions,
		PrimaryAction: primaryActionID,
	}
}

type ActionConditionRule struct {
	Name          string             `xml:"Name`
	Enabled       bool               `xml:"Enabled`
	Conditions    []ConditionRequest `xml:"Conditions>Condition"`
	PrimaryAction int                `xml:"PrimaryAction`
}

type RemoveActionRule struct {
	XMLNS   string   `xml:"xmlns,attr"`
	RuleID  int      `xml:"RuleID"`
	XMLName xml.Name `xml:"aa:RemoveActionRule"`
}

func NewRemoveActionRuleRequest(ruleID int) SoapEnvelope {
	config := RemoveActionRule{RuleID: ruleID, XMLNS: "http://www.axis.com/vapix/ws/action1"}
	return NewSoapEnvelope(NewSoapBody(config))
}
