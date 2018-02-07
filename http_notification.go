package axisevents

// HTTPNotificationHandler configures the querystring e.g. http://192.168.1.106?Message=testing&source=livingroom
type HTTPNotificationHandler interface {
	// NotifyEmptyQueryString will set an empty query string
	NotifyEmptyQueryString() RecipientsHandler
	// Notify message = testing, parameters = source=livingroom will give the result Message=testing&source=livingroom
	Notify(message string, parameters string) RecipientsHandler
}

type httpNotificationHandler struct {
	collector collector
}

func (h *httpNotificationHandler) NotifyEmptyQueryString() RecipientsHandler {
	return h.Notify("", "")
}

func (h *httpNotificationHandler) Notify(message string, parameters string) RecipientsHandler {
	h.collector.templateToken = "com.axis.action.fixed.notification.http"
	h.collector.properties["parameters"] = parameters
	h.collector.properties["message"] = message
	return &recipientsHandler{collector: h.collector}
}
