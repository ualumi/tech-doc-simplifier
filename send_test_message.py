from kafka import KafkaProducer
import json

producer = KafkaProducer(
    bootstrap_servers='localhost:29092',  # внешний порт брокера
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)

test_message = {
    "text": "The user interface can be customized according to user preferences."
}

producer.send("model_requests", value=test_message)
producer.flush()
print("Тестовое сообщение отправлено.")
