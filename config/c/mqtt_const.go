package c

import mqtt "github.com/eclipse/paho.mqtt.golang"

const (
	Broker   = "tcp://raspberrypi:1883"
	ClientId = "go_mqtt_client"

	TopicAuth       = "topic/auth"
	TopicAuthReturn = "topic/auth-return"
	TopicOperate    = "topic/operate"
)

var (
	MqttClient mqtt.Client
)
