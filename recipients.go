package axisevents

type RecipientsHandler struct {
	collector collector
}

// HTTP get request.
func (h *RecipientsHandler) HTTP(
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

// HTTP get request.
func (h *RecipientsHandler) HTTPS(
	url string,
	validateServerCert bool,
	username string,
	password string,
	proxy string,
	proxyPort string,
	proxyUsername string,
	proxyPassword string,
	qos string) Command {

	cert := 0
	if validateServerCert {
		cert = 1
	}
	h.collector.properties["upload_url"] = url
	h.collector.properties["validate_server_cert"] = cert
	h.collector.properties["login"] = username
	h.collector.properties["password"] = password
	h.collector.properties["proxy_host"] = proxy
	h.collector.properties["proxy_port"] = proxyPort
	h.collector.properties["proxy_login"] = proxyUsername
	h.collector.properties["proxy_password"] = proxyPassword
	h.collector.properties["qos"] = qos
	return &command{collector: h.collector}
}
