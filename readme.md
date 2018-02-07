A go library to simplify configuration of events such as record video when motion is detected. The library only targets [Axis Communication](http://www.axis.com) cameras.

## Table of contents

- [How to use it](#How to use it)
- [Installation](#installation)
- [Prerequisites](#prerequisites)


---

## How to use it
```go
var device = Device{
	Address:  address,
	Username: username,
	Password: password,
}

// Record video to the cameras sd-card when motion is detected.
err := NewMotionDetectionHandler("yourname", true).Record().Video().SdCard().ExecuteOn(context.TODO(), device)

// and/or

// Send a http get request when motion is detected
err := NewMotionDetectionHandler("yourname", true).HttpNotification().Notify("testing", "source=livingroom").HTTP("http://192.168.1.106", "", "", "", "", "", "", "").ExecuteOn(context.TODO(), device)

```

## Installation
TODO

## Prerequisites

The library support cameras with the following characteristics:

- __Property__: `Properties.API.HTTP.Version=3`
- __Firmware__: 5.50 and later