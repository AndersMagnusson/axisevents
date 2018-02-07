package axisevents

// RecipientsHandler definies recipients such as http, https etc.
type RecipientsHandler interface {
	// Http get request.
	HTTP(url string, username string, password string, proxy string, proxyPort string, proxyUsername string, proxyPassword string, qos string) Command
}

type recipientsHandler struct {
	collector collector
}

// HTTP get request.
func (h *recipientsHandler) HTTP(
	url string,
	username string,
	password string,
	proxy string,
	proxyPort string,
	proxyUsername string,
	proxyPassword string,
	qos string) Command {

	h.collector.properties["upload_url"] = url
	h.collector.properties["login"] = username
	h.collector.properties["password"] = password
	h.collector.properties["proxy_host"] = proxy
	h.collector.properties["proxy_port"] = proxyPort
	h.collector.properties["proxy_login"] = proxyUsername
	h.collector.properties["proxy_password"] = proxyPassword
	h.collector.properties["qos"] = qos
	return &command{collector: h.collector}
}
