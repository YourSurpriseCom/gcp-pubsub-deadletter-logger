# GCP Pub/Sub dead-letter logger
[![Go Report Card](https://goreportcard.com/badge/github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger)](https://goreportcard.com/report/github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger) 
![workflow ci](https://github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/actions/workflows/ci.yml/badge.svg)
![workflow release ](https://github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/actions/workflows/release.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview
This application is used to log dead-letter events to Google Cloud Logging. 

### Example usage
The BigQuery subscription can sent every event directly to a BigQuery table. 
When an event cannot be added to the BigQuery table, the error is added to the event message and sent to the dead-letter topic.
Setting this application as an Cloud-Run push subscription on the dead-letter topic, the error and event will be logged to Google Cloud Logging.
This way its possible to see the error, which occurred on the BigQuery subscription directly in Google Cloud Logging.

## Installation
This application should be installed as a dead-letter topic push subscription to a Cloud-Run instance running this application.

### Configuration
Setup a Cloud-Run instance and use the docker container `yoursurprise/gcp-pubsub-deadletter-logger` with the following environment variables:

* `GOOGLE_PROJECT_ID` - holds the Google project id where it should write to.
* `LOG_LEVEL` - Enable optional debug information (`info` or `debug`)


## Communication flow
To use this service, your topic needs a dead letter topic, which has a push subscription to a cloud run container running this software. See graph below:

```
+--------+        +--------------+            +------------+       +-------------------+       +-----------------+
| NORMAL |        | PUSH         |  (failed)  | DEADLETTER |       | PUSH Subscription |       |  Cloud Logging  |
| Topic  |  ==>   | Subscription |     ==>    | Topic      |  ==>  | Deadletter-logger |  ==>  | (event + error) |
+--------+        +--------------+            +------------+       +-------------------+       +-----------------+ 
```

## Usage
This application is available as a docker container and can be started as following:

```shell
foo@bar:~$ docker run --name gcp-pubsub-deadletter-logger \
    -e GOOGLE_PROJECT_ID="<google-project-id>" \
    -e LOG_LEVEL="debug" \
    yoursurprise/gcp-pubsub-deadletter-logger:latest
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
