import emitter
import tkinter
import json

# key valid for 24 hours
key = "m6IK_VC94AAIqOVZZRe-x8NSED2PMfeg"
channel = "camarao-iot"

emitter = emitter.Emitter()

options = {"secure": True}
emitter.connect(options)
emitter.on("connect", lambda: print("Connected\n\n"))
emitter.on("disconnect", lambda: print("Disconnected\n\n"))
emitter.on("presence", lambda p: print("Presence message : '" + str(p) + "'\n\n"))
emitter.on("message", lambda m: print("Message received: " + m.asString() + "\n\n"))
emitter.loopStart()

emitter.publish(key, channel, json.dumps({"test": "test 123"}))

import time, sys

while True:
    try:
        time.sleep(.3)
    except KeyboardInterrupt:
        print("Shutting down...")
        emitter.unsubscribe(key, channel)
        emitter.loopStop()
        emitter.disconnect()
        sys.exit()
