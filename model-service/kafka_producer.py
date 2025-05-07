import os
from kafka import KafkaProducer
import json

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "localhost:9092")
RESPONSE_TOPIC = os.getenv("KAFKA_RESPONSE_TOPIC", "model_response")

producer = KafkaProducer(
    bootstrap_servers=KAFKA_BROKER,
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)

def send_result_to_kafka(result, correlation_id):
    print(f"Sending to '{RESPONSE_TOPIC}': key={correlation_id}, value={result}")
    producer.send(RESPONSE_TOPIC, key=correlation_id.encode('utf-8'), value=result)
    producer.flush()
