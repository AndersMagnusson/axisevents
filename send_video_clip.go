package axisevents

type SendVideoClipper struct {
	collector collector
}

func (h *SendVideoClipper) SendDefault(
	filename string,
	customParams string,
) *RecipientsHandler {
	return h.Send(filename, customParams, "", 2000, 3000)
}

func (h *SendVideoClipper) Send(
	filename string,
	customParams string,
	folder string,
	prebuffer int,
	postbuffer int,
) *RecipientsHandler {
	h.collector.templateToken = "com.axis.action.fixed.send_videoclip.https"
	h.collector.properties["pre_duration"] = prebuffer
	h.collector.properties["post_duration"] = postbuffer
	h.collector.properties["filename"] = filename
	h.collector.properties["custom_params"] = customParams
	h.collector.properties["max_duration"] = "0"
	h.collector.properties["max_file_size"] = "0"
	h.collector.properties["create_folder"] = folder
	h.collector.properties["stream_options"] = "streamprofile=Quality"

	return &RecipientsHandler{collector: h.collector}
}
