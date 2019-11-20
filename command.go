package axisevents

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AndersMagnusson/axisevents/actions"
)

// Command executes the request.
type Command interface {
	// With specify your Do function.
	With(do func(req *http.Request)) Command
	// ExecuteOn the event on the specified device.
	ExecuteOn(ctx context.Context, device Device) error
}

type command struct {
	collector collector
	do        func(req *http.Request)
}

// With specify your Do function.
func (h *command) With(do func(req *http.Request)) Command {
	h.do = do
	return h
}

// ExecuteOn the event on the specified device.
func (h *command) ExecuteOn(ctx context.Context, device Device) error {
	handler := actions.NewActionHandler(ctx, device.Username, device.Password, device.Address)

	existingRule, found, err := handler.GetActionRule(h.collector.name)
	if found {
		err = handler.RemoveActionRule(existingRule.RuleID)
		if err != nil {
			return fmt.Errorf("Failed to remove existing ActionRule, %s", err.Error())
		}
	}

	result, found, err := handler.GetActionConfiguration(h.collector.name)
	if found {
		err = handler.RemoveActionConfigurations(result.ConfigurationID)
		if err != nil {
			return fmt.Errorf("Failed to remove existing ActionConfiguration, %s", err.Error())
		}
	}

	configID, err := handler.AddActionConfigurations(h.collector.name, h.collector.templateToken, getParameters(h.collector.properties))
	if err != nil {
		return fmt.Errorf("Failed to add ActionConfiguration, %s", err.Error())
	}

	messageContent := actions.NewMessageContent(h.collector.messageContent)

	if h.collector.version < 4 {
		topixExpression := actions.NewTopicExpression(h.collector.topicExpression)
		startEvent := actions.NewStartEvent(topixExpression, messageContent)

		_, err = handler.AddActionRule(h.collector.name, h.collector.enabled, startEvent, configID)
		if err != nil {
			return fmt.Errorf("Failed to add ActionRule, %s", err.Error())
		}
	} else {
		topixExpression := actions.NewTopicOasisTopicExpression(h.collector.topicExpression)
		c1 := actions.ConditionRequest{
			TopicExpression: topixExpression,
			MessageContent:  messageContent,
		}

		_, err = handler.AddActionConditionRule(
			h.collector.name, h.collector.enabled, []actions.ConditionRequest{c1}, configID)
		if err != nil {
			return fmt.Errorf("Failed to add ActionRule, %s", err.Error())
		}

		// c1 := actions.NewActionConditionRule(name string, enabled bool, conditions []actions.ConditionRequest, primaryActionID int)
	}

	return nil
}

func getParameters(parameters map[string]interface{}) []actions.Parameter {
	result := make([]actions.Parameter, 0)
	for k, v := range parameters {
		result = append(result, actions.Parameter{Name: k, Value: v})
	}
	return result
}
