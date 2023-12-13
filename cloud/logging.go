package cloud

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/logging"
	"github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/config"
	log "github.com/jlentink/yaglogger"
)

var (
	ctx       context.Context
	GCPLogger *logging.Logger
)

func init() {
	ctx = context.Background()

	client, err := logging.NewClient(ctx, config.AppConfig.GoogleProjectId)
	if err != nil {
		log.Fatalf("could not create logging client: %v", err)
	}
	GCPLogger = client.Logger("deadletter-log")
}

func LogFullMessage(data []byte) bool {
	var rawJSONMessage json.RawMessage
	if err := json.Unmarshal([]byte(data), &rawJSONMessage); err != nil {
		log.Warn("pubsub event does not contain valid json:  %v", err)
		return false
	}

	err := GCPLogger.LogSync(ctx, logging.Entry{
		Payload:  rawJSONMessage,
		Severity: logging.Info,
	})

	if err != nil {
		log.Warn("could not log to gcp logger: %v", err)
		return false
	}

	log.Debug("log sent to google cloud logging!")
	return true
}

func LogError(data []byte) bool {

	// Create struct with all known data fields
	var deadletter Deadletter
	err := json.Unmarshal(data, &deadletter)
	if err != nil {
		return false
	}

	// As we do not know what is in the data object we cannot use a struct
	var rawJSONMessage json.RawMessage
	if err := json.Unmarshal([]byte(deadletter.Message.Data), &rawJSONMessage); err != nil {
		log.Warn("[pubsub event does not contain exptected format: %v", err)
		return false
	}

	logMessage := LogMessage{
		Message: rawJSONMessage,
		Subscription: LogMessageSubscription{
			DeliveryCount: deadletter.Message.Attributes.SubscriptionDeliveryCount,
			Error:         deadletter.Message.Attributes.SubscriptionError,
			ProjectID:     deadletter.Message.Attributes.SubscriptionProjectID,
			Name:          deadletter.Message.Attributes.Subscription,
		},
		Topic: LogMessageTopic{
			PublishTime: deadletter.Message.Attributes.TopicPublishTime,
			Schema: LogMessageTopicSchema{
				Name:     deadletter.Message.Attributes.TopicSchemaName,
				Encoding: deadletter.Message.Attributes.TopicSchemaEncoding,
				Revision: deadletter.Message.Attributes.TopicSchemaRevision,
			},
		},
	}
	log.Debug("object: %s", logMessage)

	//Log message blocking
	err = GCPLogger.LogSync(ctx, logging.Entry{
		Payload:  logMessage,
		Severity: logging.Error,
	})

	return err != nil
}
