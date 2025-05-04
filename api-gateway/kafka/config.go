package kafka

//func SetupKafka() {
//	InitProducer("broker:29092", "user_requests")
//}
func SetupKafka() {
	InitProducer("broker:9092", "user_requests")
	InitConsumer("broker:9092", "user_responses")
}
