package axisevents

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	USERNAME = "AXISEVENTS_USERNAME"
	PASSWORD = "AXISEVENTS_PASSWORD"
	ADDRESS  = "AXISEVENTS_ADDRESS"
)

var (
	username   = os.Getenv(USERNAME)
	password   = os.Getenv(PASSWORD)
	address    = os.Getenv(ADDRESS)
	testServer = "https://172.25.125.218:58674/"
	// go test -run=TestSendVideoClipLongRunning -timeout 100s --httptest.serve=:58674
)

var device = Device{
	Address:  address,
	Username: username,
	Password: password,
}

func TestMotionDetectionRecordingOnSdCardLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("RecordingOnSdCard tests on a real camera")
	}
	verifyEnvVariables(t)
	err := NewMotionDetectionHandler("newtest", true, 4).
		Record().
		Video().
		SdCard().
		ExecuteOn(context.TODO(), device)

	if err != nil {
		t.Error("Failed to toggle motion detection on device: ", err)
	}
}

func TestHttpNotificationLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("HttpNotification tests on a real camera")
	}

	verifyEnvVariables(t)
	// url will be http://192.168.1.106?Message=testing&source=livingroom
	err := NewMotionDetectionHandler("httpNotifications", true, 4).
		HttpNotification().
		Notify("testing", "source=livingroom").
		HTTP(testServer, "", "", "", "", "", "", "").
		ExecuteOn(context.TODO(), device)

	if err != nil {
		t.Error("Failed to set http notification: ", err)
	}
}

// TestSendVideoClipLongRunning tests on a real camera
func TestSendVideoClipLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("SendVideoClip tests on a real camera")
	}

	verifyEnvVariables(t)
	ts := startTestServer(t)
	defer ts.Close()
	t.Logf("TestServer: %s", ts.URL)

	err := NewMotionDetectionHandler(
		"sendVideoClipTest", true, 4).
		VideoClip().
		Send("video-test-%y-%m-%d_%H-%M-%S-%f_#s.mkv", "hej=hopp&danke=gut", "", 1000, 1000).
		HTTPS(
			testServer, false, "", "", "", "", "", "", "").
		ExecuteOn(context.Background(), device)

	if err != nil {
		t.Error("Failed to set SendVideoClip: ", err)
	} else {
		time.Sleep(time.Second * 100)
	}
}

func startTestServer(t *testing.T) *httptest.Server {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body %s", err)
		} else {
			log.Printf("URL: %s", r.URL.String())
			log.Printf("Headers: %v", r.Header)
			_, params, err := mime.ParseMediaType(r.Header.Get("Content-Disposition"))
			if err != nil {
				log.Printf("Failed to get filename: %s", err)
			}
			filename, ok := params["filename"] // set to "foo.png"
			if !ok {
				t.Error("Expected to get filename")
			}
			now := time.Now()
			expectedPrefix := fmt.Sprintf("video-test-%d-%d-%d", now.Year(), int(now.Month()), now.Day())
			if !strings.HasPrefix(filename, expectedPrefix) {
				t.Errorf("Expected file to start with %s but was %s", expectedPrefix, filename)
			}
			n, err := ioutil.Discard.Write(b)
			if err != nil {
				t.Errorf("Failed to save file %s", err)
			}
			if n != len(b) {
				t.Errorf("Expected %d to be save but was %d", len(b), n)
			}
		}
		fmt.Fprintln(w, "Hello, client")
	}))

	return ts
}

func verifyEnvVariables(t *testing.T) {
	if len(device.Address) == 0 {
		t.Fatalf("You need to set env variable %s", ADDRESS)
	}

	if len(device.Username) == 0 {
		t.Fatalf("You need to set env variable %s", USERNAME)
	}

	if len(device.Password) == 0 {
		t.Fatalf("You need to set env variable %s", PASSWORD)
	}
}
