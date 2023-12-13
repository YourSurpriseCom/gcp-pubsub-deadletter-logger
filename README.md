# GCP Pub/Sub dead letter logger
[![Go Report Card](https://goreportcard.com/badge/github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger)](https://goreportcard.com/report/github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger) 
![workflow ci](https://github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/actions/workflows/ci.yml/badge.svg)
![workflow release ](https://github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/actions/workflows/release.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Overview
This logger can be used to log dead letter events to Google Cloud Logging. 

## Communication flow
To use this service, your topic needs a dead letter topic, which has a push subscription to a cloud run container running this software. See graph below:

```
+--------+        +--------------+            +------------+       +-------------------+       +-----------------+
| NORMAL |        | PUSH         |  (failed)  | DEADLETTER |       | PUSH Subscription |       |  Cloud Logging  |
| Topic  |  ==>   | Subscription |     ==>    | Topic      |  ==>  | Deadletter-logger |  ==>  | (event + error) |
+--------+        +--------------+            +------------+       +-------------------+       +-----------------+ 
```

## Local developing
* Run `make init` to create the `.env` file and install the required packages.
* Edit `.env` and change the values where needed.
* Run `make run` to start the application.
* Send the following message as a post to this application while running on port `8888`:

```json
{
    "message":  {
        "attributes":{
            "CloudPubSubDeadLetterSourceDeliveryCount":"1337", 
            "CloudPubSubDeadLetterSourceDeliveryErrorMessage":"Some error occured, which why this message is now a deadletter", 
            "CloudPubSubDeadLetterSourceSubscription":"example-subscription", 
            "CloudPubSubDeadLetterSourceSubscriptionProject":"google-project-id", 
            "CloudPubSubDeadLetterSourceTopicPublishTime":"2023-12-01T20:30:00.000+00:00" 
        }, 
        "data":"eyJmaWVsZDEiIDogImV4YW1wbGUgZGF0YSIsImZpZWxkMiIgOiAxMzM3LCAiZmllbGQzIiA6ICJ0aGlzIG1lc3NhZ2Ugc2hvdWxkIGJlIHJlYWRhYmxlIGluc2lkZSB0aGUgZ29vZ2xlIGNsb3VkIGxvZyJ9", 
        "messageId":"123456790", 
        "message_id":"123456790", 
        "publishTime":"2023-12-01T20:30:00.000Z", 
        "publish_time":"2023-12-01T20:30:01.000Z"
    }
}
```
