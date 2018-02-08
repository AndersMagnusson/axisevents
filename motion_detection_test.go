package axisevents

import (
	"context"
	"os"
	"testing"
)

var (
	username = os.Getenv("axisevents-username")
	password = os.Getenv("axisevents-password")
	address  = os.Getenv("axisevents-address")
)

var device = Device{
	Address:  address,
	Username: username,
	Password: password,
}

func TestMotionDetectionRecordingOnSdCardLongRunning(t *testing.T) {
	verifyEnvVariables(t)
	err := NewMotionDetectionHandler("newtest", true).Record().Video().SdCard().ExecuteOn(context.TODO(), device)

	if err != nil {
		t.Error("Failed to toggle motion detection on device: ", err)
	}
}

func TestHttpNotificationLongRunning(t *testing.T) {
	verifyEnvVariables(t)
	// url will be http://192.168.1.106?Message=testing&source=livingroom
	err := NewMotionDetectionHandler("httpNotifications", true).HttpNotification().Notify("testing", "source=livingroom").HTTP("http://192.168.1.106", "", "", "", "", "", "", "").ExecuteOn(context.TODO(), device)

	if err != nil {
		t.Error("Failed to set http notification: ", err)
	}
}

func verifyEnvVariables(t *testing.T) {
	if len(device.Address) == 0 {
		t.Fatal("You need to set env variable axisevents-address")
	}

	if len(device.Username) == 0 {
		t.Fatal("You need to set env variable axisevents-username")
	}

	if len(device.Password) == 0 {
		t.Fatal("You need to set env variable axisevents-password")
	}
}
