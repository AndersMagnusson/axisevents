package axisevents

// HttpNotificationHandler configures the querystring e.g. http://192.168.1.106?Message=testing&source=livingroom
type HttpNotificationHandler struct {
	collector collector
}

func (h *HttpNotificationHandler) NotifyEmptyQueryString() *RecipientsHandler {
	return h.Notify("", "")
}

func (h *HttpNotificationHandler) Notify(message string, parameters string) *RecipientsHandler {
	h.collector.templateToken = "com.axis.action.fixed.notification.http"
	h.collector.properties["parameters"] = parameters
	h.collector.properties["message"] = message
	return &RecipientsHandler{collector: h.collector}
}
