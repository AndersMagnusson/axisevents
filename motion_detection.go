// Package axisevents simplifies the configuration of events such as record video to sd-card when the device detects motion detection
package axisevents

// Device holds the connection info for the device
type Device struct {
	Username string
	Password string
	Address  string // Address must be in format of http://ip/ or http://hostname/ or if the device is configured for ssl https://ip/
}

// NewMotionDetectionHandler creates a new MotionDetectionHandler which is used
// to define what happens when a device detects motion.
//
// Supported actions:
// Record video to the device SD-card.
func NewMotionDetectionHandler(name string, enabled bool) MotionDetectionHandler {
	collector := collector{
		name:            name,
		enabled:         enabled,
		topicExpression: "tns1:RuleEngine/tnsaxis:VideoMotionDetection/motion",
		messageContent:  `boolean(//SimpleItem[@Name="active" and @Value="1"])`,
		properties:      make(map[string]interface{}),
	}
	return &motionDetectionHandler{
		collector: collector,
	}
}

// MotionDetectionHandler defines operations to use when a device detects motion.
type MotionDetectionHandler interface {
	// Send http notifications when motion is detected.
	HttpNotification() HTTPNotificationHandler
	// Record motion.
	Record() RecordVideoHandler
}

type motionDetectionHandler struct {
	collector collector
}

// HttpNotification sends http notifications when motion is detected.
func (h *motionDetectionHandler) HttpNotification() HTTPNotificationHandler {
	return &httpNotificationHandler{
		collector: h.collector,
	}
}

// Record motion.
func (h *motionDetectionHandler) Record() RecordVideoHandler {
	return &recordVideoHandler{
		collector: h.collector,
	}
}
