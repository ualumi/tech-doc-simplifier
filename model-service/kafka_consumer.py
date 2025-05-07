import os
from kafka import KafkaConsumer
import json
from model.simplifier import simplify_text
from kafka_producer import send_result_to_kafka

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "localhost:9092")
REQUEST_TOPIC = os.getenv("KAFKA_REQUEST_TOPIC", "model_requests")

consumer = KafkaConsumer(
    REQUEST_TOPIC,
    bootstrap_servers=KAFKA_BROKER,
    value_deserializer=lambda m: json.loads(m.decode('utf-8')),
    auto_offset_reset='earliest',
    enable_auto_commit=True,
    group_id='model-service-group'
)

def consume():
    print(f"Listening to '{REQUEST_TOPIC}' on {KAFKA_BROKER}...")
    for message in consumer:
        request = message.value
        print(f"Received: {request}")
        
        text = request.get("text")
        token = request.get("token")  # correlation ID для ответа

        if text and token:
            simplified = simplify_text(text)
            result = {
                "original": {"text": text, "token": token},
                "simplified": {"text": simplified, "token": token}
            }
            send_result_to_kafka(result, correlation_id=token)

