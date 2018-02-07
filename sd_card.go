package axisevents

// SdCardHandler record to the device sd-card.
type SdCardHandler interface {
	SdCard() Command
}

type sdCardHandler struct {
	collector collector
}

// SdCard record video to the device sd-card.
func (h *sdCardHandler) SdCard() Command {
	h.collector.properties["storage_id"] = "SD_DISK"
	return &command{collector: h.collector}
}
