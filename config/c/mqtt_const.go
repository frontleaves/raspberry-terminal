package c

import mqtt "github.com/eclipse/paho.mqtt.golang"

const (
	Broker   = "tcp://localhost:1883"
	ClientId = "go_mqtt_client"

	TopicAuth          = "topic/auth"
	TopicAuthReturn    = "topic/auth-return"
	TopicOperate       = "topic/operate"
	TopicOperateReturn = "topic/operate-return"
)

var (
	MqttClient mqtt.Client
)
