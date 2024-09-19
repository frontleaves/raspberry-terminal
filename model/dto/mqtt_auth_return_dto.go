package dto

type MqttAuthReturnDTO struct {
	Device     string `json:"device"`
	Authorized bool   `json:"authorized"`
}
