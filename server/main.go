package main

import (
	"fmt"
	"time"

	emitter "github.com/emitter-io/go"
)

// key valid for 24 hours
const key = "m6IK_VC94AAIqOVZZRe-x8NSED2PMfeg"
const channel = "camarao-iot"

func main() {
	// Create the options with default values
	o := emitter.NewClientOptions()

	// Set the message handler
	o.SetOnMessageHandler(func(client emitter.Emitter, msg emitter.Message) {
		fmt.Printf("Received message: %s\n", msg.Payload())
	})

	// Create a new emitter client and connect to the broker
	c := emitter.NewClient(o)
	sToken := c.Connect()
	if sToken.Wait() && sToken.Error() != nil {
		panic("Error on Client.Connect(): " + sToken.Error().Error())
	}

	// Subscribe to the presence demo channel
	c.Subscribe(key, channel)

	// Publish to the channel
	c.Publish(key, channel, "hello")

	// Ask for presence
	r := emitter.NewPresenceRequest()
	r.Key = key
	r.Channel = channel
	c.Presence(r)

	for {
		time.Sleep(time.Second)
	}
}
