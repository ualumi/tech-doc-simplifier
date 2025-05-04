from kafka import KafkaConsumer
import json

consumer = KafkaConsumer(
    "model_response",
    bootstrap_servers="localhost:29092",
    value_deserializer=lambda m: json.loads(m.decode('utf-8')),
    auto_offset_reset='earliest',
    enable_auto_commit=True,
    group_id='test-response-reader'
)

print("Ожидание ответа...")
for msg in consumer:
    print("Результат от модели:", msg.value)
    break
