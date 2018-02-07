package actions

import "context"

func NewActionHandler(ctx context.Context, username string, password string, address string) ActionHandler {
	return &actionHandler{
		username: username,
		password: password,
		address:  address,
		ctx:      ctx,
	}
}

type ActionHandler interface {
	GetActionConfiguration(name string) (ActionConfigurationResponse, bool, error)
	GetActionConfigurations() ([]ActionConfigurationResponse, error)
	AddActionConfigurations(name string, templateToken string, parameters []Parameter) (int, error)
	RemoveActionConfigurations(configurationID int) error

	GetActionTemplates() ([]ActionTemplateResponse, error)
	GetRecipientTemplates() ([]RecipientTemplateResponse, error)

	GetActionRule(name string) (ActionRuleResponse, bool, error)
	GetActionRules() ([]ActionRuleResponse, error)
	AddActionRule(name string, enabled bool, startEvent StartEvent, primaryActionID int) (int, error)
	RemoveActionRule(ruleID int) error
}

type actionHandler struct {
	username   string
	password   string
	address    string
	url        string
	soapAction string
	ctx        context.Context
}

func (ah *actionHandler) GetActionConfigurations() ([]ActionConfigurationResponse, error) {
	request := NewGetActionConfigurationsRequest()
	body := NewSoapBody(request)
	envelope := NewSoapEnvelope(body)

	response := &GetActionConfigurationsSoapEnvelope{}
	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/GetActionConfigurations", envelope, response)
	if err != nil {
		return nil, err
	}
	return response.GetActionConfigurationsSoapBody.GetActionConfigurationsResponse.ActionConfigurationsResponse, nil
}

func (ah *actionHandler) GetActionConfiguration(name string) (ActionConfigurationResponse, bool, error) {
	existing, err := ah.GetActionConfigurations()
	if err != nil {
		return ActionConfigurationResponse{}, false, err
	}
	for _, a := range existing {
		if a.Name == name {
			return a, true, nil
		}
	}
	return ActionConfigurationResponse{}, false, nil
}

func (ah *actionHandler) AddActionConfigurations(name string, templateToken string, parameters []Parameter) (int, error) {
	request := NewAddActionConfigurationsRequest(name, templateToken, parameters)

	response := &AddActionConfigurationSoapEnvelope{}
	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/AddActionConfiguration", request, response)
	if err != nil {
		return 0, err
	}
	return response.AddActionConfigurationSoapBody.AddActionConfigurationResponse.ConfigurationID, nil
}

func (ah *actionHandler) RemoveActionConfigurations(configurationID int) error {
	request := NewRemoveActionConfigurationsRequest(configurationID)

	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/RemoveActionConfiguration", request, nil)
	if err != nil {
		return err
	}
	return nil
}

func (ah *actionHandler) GetActionTemplates() ([]ActionTemplateResponse, error) {
	request := NewGetActionTemplatesRequest()
	body := NewSoapBody(request)
	envelope := NewSoapEnvelope(body)

	response := &GetActionTemplatesSoapEnvelope{}
	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/GetActionTemplates", envelope, response)
	if err != nil {
		return nil, err
	}
	return response.GetActionTemplatesSoapBody.GetActionTemplatesResponse.ActionTemplatesResponse, nil
}

func (ah *actionHandler) GetRecipientTemplates() ([]RecipientTemplateResponse, error) {
	request := NewGetRecipientTemplatesRequest()
	body := NewSoapBody(request)
	envelope := NewSoapEnvelope(body)

	response := &GetRecipientTemplatesSoapEnvelope{}
	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/GetRecipientTemplates", envelope, response)
	if err != nil {
		return nil, err
	}
	return response.GetRecipientTemplatesSoapBody.GetRecipientTemplatesResponse.RecipientTemplatesResponse, nil
}

func (ah *actionHandler) GetActionRule(name string) (ActionRuleResponse, bool, error) {
	existing, err := ah.GetActionRules()
	if err != nil {
		return ActionRuleResponse{}, false, err
	}
	for _, a := range existing {
		if a.Name == name {
			return a, true, nil
		}
	}
	return ActionRuleResponse{}, false, nil
}

func (ah *actionHandler) GetActionRules() ([]ActionRuleResponse, error) {
	request := NewGetActionRulesRequest()
	body := NewSoapBody(request)
	envelope := NewSoapEnvelope(body)

	response := &GetActionRulesSoapEnvelope{}
	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/GetActionRules", envelope, response)
	if err != nil {
		return nil, err
	}
	return response.GetActionRulesSoapBody.GetActionRulesResponse.ActionRulesResponse, nil
}

func (ah *actionHandler) AddActionRule(name string, enabled bool, startEvent StartEvent, primaryActionID int) (int, error) {
	request := NewAddActionRule(NewActionRule(name, enabled, startEvent.TopicExpression, startEvent.MessageContent, primaryActionID))
	body := NewSoapBody(request)
	envelope := NewSoapEnvelope(body)

	response := &AddActionRuleSoapEnvelope{}
	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/AddActionRule", envelope, response)
	if err != nil {
		return 0, err
	}
	return response.AddActionRulenSoapBody.AddActionRuleResponse.RuleID, nil
}

func (ah *actionHandler) RemoveActionRule(ruleID int) error {
	request := NewRemoveActionRuleRequest(ruleID)

	err := SoapDo(ah.ctx, ah.username, ah.password, ah.address, "http://www.axis.com/vapix/ws/action1/RemoveActionRule", request, nil)
	if err != nil {
		return err
	}
	return nil
}
