package axisevents

import (
	"context"
	"testing"
)

var (
	username = ""
	password = ""
	address  = "192.168.1.217"
)

var device = Device{
	Address:  address,
	Username: username,
	Password: password,
}

func TestMotionDetectionRecordingOnSdCard(t *testing.T) {
	err := NewMotionDetectionHandler("newtest", true).Record().Video().SdCard().ExecuteOn(context.TODO(), device)

	if err != nil {
		t.Error("Failed to toggle motion detection on device: ", err)
	}
}

func TestHttpNotification(t *testing.T) {
	// url will be http://192.168.1.106?Message=testing&source=livingroom
	err := NewMotionDetectionHandler("httpNotifications", true).HttpNotification().Notify("testing", "source=livingroom").HTTP("http://192.168.1.106", "", "", "", "", "", "", "").ExecuteOn(context.TODO(), device)

	if err != nil {
		t.Error("Failed to set http notification: ", err)
	}
}
