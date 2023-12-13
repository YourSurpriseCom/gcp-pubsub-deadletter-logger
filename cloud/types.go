package cloud

import "encoding/json"

type Deadletter struct {
	Message Message `json:"message"`
}

type Message struct {
	Data        []byte `json:"data"`
	PublishTime string `json:"publishTime"`
	MessageID   string `json:"messageId"`
	Attributes  MessageAttributes
}

type MessageAttributes struct {
	SubscriptionDeliveryCount string `json:"CloudPubSubDeadLetterSourceDeliveryCount"`
	SubscriptionError         string `json:"CloudPubSubDeadLetterSourceDeliveryErrorMessage"`
	SubscriptionProjectID     string `json:"CloudPubSubDeadLetterSourceSubscriptionProject"`
	Subscription              string `json:"CloudPubSubDeadLetterSourceSubscription"`
	TopicPublishTime          string `json:"CloudPubSubDeadLetterSourceTopicPublishTime"`
	TopicSchemaEncoding       string `json:"googclient_schemaencoding"`
	TopicSchemaName           string `json:"googclient_schemaname"`
	TopicSchemaRevision       string `json:"googclient_schemarevisionid"`
}

type LogMessage struct {
	Message      json.RawMessage        `json:"message"`
	Subscription LogMessageSubscription `json:"subscription"`
	Topic        LogMessageTopic        `json:"topic"`
}

type LogMessageSubscription struct {
	DeliveryCount string `json:"deliveryCount"`
	Error         string `json:"error"`
	ProjectID     string `json:"projectID"`
	Name          string `json:"name"`
}

type LogMessageTopic struct {
	PublishTime string                `json:"messagePublishTime"`
	Schema      LogMessageTopicSchema `json:"schema"`
}

type LogMessageTopicSchema struct {
	Name     string `json:"name"`
	Encoding string `json:"encoding"`
	Revision string `json:"revision"`
}
