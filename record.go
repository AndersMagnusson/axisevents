package axisevents

// RecordVideoHandler settings for the video.
type RecordVideoHandler interface {
	Video() SdCardHandler
	VideoCustom(prebuffer int, postbuffer int, streamprofile string) SdCardHandler
}

type recordVideoHandler struct {
	collector collector
}

// RecordVideo will use default values for pre, postbuffer and streamprofile.
func (h *recordVideoHandler) Video() SdCardHandler {
	return h.VideoCustom(2000, 30000, "streamprofile=Quality")
}

// RecordVideoCustom should be used when you want to override the default values.
func (h *recordVideoHandler) VideoCustom(prebuffer int, postbuffer int, streamprofile string) SdCardHandler {
	h.collector.properties["pre_duration"] = prebuffer
	h.collector.properties["post_duration"] = postbuffer
	h.collector.properties["stream_options"] = streamprofile
	h.collector.templateToken = "com.axis.action.fixed.recording.storage"
	nextHandler := &sdCardHandler{
		collector: h.collector,
	}

	return nextHandler
}
