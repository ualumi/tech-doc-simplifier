package kafka

func SetupKafka() {
	InitProducer("broker:29092", "user_requests")
}
