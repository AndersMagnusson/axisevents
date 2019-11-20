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
// Supported versions:
// 2,3,4 defaults to version 2
// Supported actions:
// Record video to the device SD-card.
func NewMotionDetectionHandler(
	name string,
	enabled bool,
	version int,
) *MotionDetectionHandler {
	var topicExpression string
	switch version {
	case 2:
		topicExpression = "tns1:RuleEngine/tnsaxis:VideoMotionDetection/motion"
	case 3:
		topicExpression = ""
	case 4:
		topicExpression = "tnsaxis:CameraApplicationPlatform/VMD/Camera1ProfileANY"
	default:
		topicExpression = "tns1:RuleEngine/tnsaxis:VideoMotionDetection/motion"
	}

	collector := collector{
		name:            name,
		enabled:         enabled,
		version:         version,
		topicExpression: topicExpression,
		messageContent:  `boolean(//SimpleItem[@Name="active" and @Value="1"])`,
		properties:      make(map[string]interface{}),
	}
	return &MotionDetectionHandler{
		collector: collector,
	}
}

type MotionDetectionHandler struct {
	collector collector
}

// HttpNotification sends http notifications when motion is detected.
func (h *MotionDetectionHandler) HttpNotification() *HttpNotificationHandler {
	return &HttpNotificationHandler{
		collector: h.collector,
	}
}

// Record motion.
func (h *MotionDetectionHandler) Record() RecordVideoHandler {
	return &recordVideoHandler{
		collector: h.collector,
	}
}

// Send video clip.
func (h *MotionDetectionHandler) VideoClip() *SendVideoClipper {
	return &SendVideoClipper{
		collector: h.collector,
	}
}
