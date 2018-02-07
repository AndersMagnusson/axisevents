package actions

import "encoding/xml"

type GetActionConfigurationsSoapEnvelope struct {
	XMLName                         xml.Name `xml:"Envelope"`
	GetActionConfigurationsSoapBody GetActionConfigurationsSoapBody
}

type GetActionConfigurationsSoapBody struct {
	XMLName                         xml.Name `xml:"Body"`
	GetActionConfigurationsResponse GetActionConfigurationsResponse
}

type GetActionConfigurations struct {
	XMLNS   string   `xml:"xmlns,attr"`
	XMLName xml.Name `xml:"aa:GetActionConfigurations"`
}

func NewGetActionConfiguration() GetActionConfigurations {
	return GetActionConfigurations{
		XMLNS: "http://www.axis.com/vapix/ws/action1",
	}
}

type GetActionConfigurationsResponse struct {
	XMLName                      xml.Name                      `xml:"GetActionConfigurationsResponse"`
	ActionConfigurationsResponse []ActionConfigurationResponse `xml:"ActionConfigurations>ActionConfiguration"`
}

type ActionConfigurationResponse struct {
	XMLName         xml.Name            `xml:"ActionConfiguration"`
	ConfigurationID int                 `xml:"ConfigurationID"`
	Name            string              `xml:"Name"`
	TemplateToken   string              `xml:"TemplateToken"`
	Parameters      []ParameterResponse `xml:"Parameters>Parameter"`
}

type ParameterResponse struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

// ActionTemplates
type GetActionTemplatesSoapEnvelope struct {
	XMLName                    xml.Name `xml:"Envelope"`
	GetActionTemplatesSoapBody GetActionTemplatesSoapBody
}

type GetActionTemplatesSoapBody struct {
	XMLName                    xml.Name `xml:"Body"`
	GetActionTemplatesResponse GetActionTemplatesResponse
}

type GetActionTemplatesResponse struct {
	XMLName                 xml.Name                 `xml:"GetActionTemplatesResponse"`
	ActionTemplatesResponse []ActionTemplateResponse `xml:"ActionTemplates>ActionTemplate"`
}

type ActionTemplateResponse struct {
	XMLName           xml.Name            `xml:"ActionTemplate"`
	RecipientTemplate string              `xml:"RecipientTemplate"`
	TemplateToken     string              `xml:"TemplateToken"`
	Parameters        []ParameterResponse `xml:"Parameters>Parameter"`
}

// Recipients
type GetRecipientTemplatesSoapEnvelope struct {
	XMLName                       xml.Name `xml:"Envelope"`
	GetRecipientTemplatesSoapBody GetRecipientTemplatesSoapBody
}

type GetRecipientTemplatesSoapBody struct {
	XMLName                       xml.Name `xml:"Body"`
	GetRecipientTemplatesResponse GetRecipientTemplatesResponse
}

type GetRecipientTemplatesResponse struct {
	XMLName                    xml.Name                    `xml:"GetRecipientTemplatesResponse"`
	RecipientTemplatesResponse []RecipientTemplateResponse `xml:"RecipientTemplates>RecipientTemplate"`
}

type RecipientTemplateResponse struct {
	XMLName       xml.Name            `xml:"RecipientTemplate"`
	TemplateToken string              `xml:"TemplateToken"`
	Parameters    []ParameterResponse `xml:"Parameters>Parameter"`
}

// ActionRules
type GetActionRulesSoapEnvelope struct {
	XMLName                xml.Name `xml:"Envelope"`
	GetActionRulesSoapBody GetActionRulesSoapBody
}

type GetActionRulesSoapBody struct {
	XMLName                xml.Name `xml:"Body"`
	GetActionRulesResponse GetActionRulesResponse
}

type GetActionRulesResponse struct {
	XMLName             xml.Name             `xml:"GetActionRulesResponse"`
	ActionRulesResponse []ActionRuleResponse `xml:"ActionRules>ActionRule"`
}

type ActionRuleResponse struct {
	XMLName       xml.Name            `xml:"ActionRule"`
	Name          string              `xml:"Name"`
	RuleID        int                 `xml:"RuleID"`
	Enabled       bool                `xml:"Enabled"`
	PrimaryAction int                 `xml:"PrimaryAction"`
	Parameters    []ConditionResponse `xml:"Conditions>Condition"`
}

type ConditionResponse struct {
	TopicExpression string `xml:"Name,attr"`
	MessageContent  string `xml:"Value,attr"`
}

type AddActionConfigurationSoapEnvelope struct {
	XMLName                        xml.Name `xml:"Envelope"`
	AddActionConfigurationSoapBody AddActionConfigurationSoapBody
}

type AddActionConfigurationSoapBody struct {
	XMLName                        xml.Name `xml:"Body"`
	AddActionConfigurationResponse AddActionConfigurationResponse
}

type AddActionConfigurationResponse struct {
	XMLName         xml.Name `xml:"AddActionConfigurationResponse"`
	ConfigurationID int      `xml:"ConfigurationID"`
}

type AddActionRuleSoapEnvelope struct {
	XMLName                xml.Name `xml:"Envelope"`
	AddActionRulenSoapBody AddActionRulenSoapBody
}

type AddActionRulenSoapBody struct {
	XMLName               xml.Name `xml:"Body"`
	AddActionRuleResponse AddActionRuleResponse
}

type AddActionRuleResponse struct {
	XMLName xml.Name `xml:"AddActionRuleResponse"`
	RuleID  int      `xml:"RuleID"`
}
