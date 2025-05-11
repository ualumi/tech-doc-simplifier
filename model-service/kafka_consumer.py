import os
import json
from kafka import KafkaConsumer
from model.simplifier import simplify_text
from kafka_producer import send_result_to_kafka

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "broker:9092")
REQUEST_TOPIC = os.getenv("KAFKA_REQUEST_TOPIC", "model_requests")

# Не создаём consumer при импорте!
consumer = None

def init_consumer():
    global consumer
    if consumer is None:
        consumer = KafkaConsumer(
            REQUEST_TOPIC,
            bootstrap_servers=KAFKA_BROKER,
            value_deserializer=lambda m: json.loads(m.decode('utf-8')),
            auto_offset_reset='earliest',
            enable_auto_commit=True,
            group_id='model-service-group'
        )

def consume():
    init_consumer()
    print(f"Listening to '{REQUEST_TOPIC}' on {KAFKA_BROKER}...")
    for message in consumer:
        request = message.value
        correlation_id = message.key.decode('utf-8') if message.key else None
        print(f"Received: key={correlation_id}, value={request}")
        
        text = request.get("text")
        token = request.get("token")

        if text and correlation_id:
            simplified = simplify_text(text)
            result = {
                "original": {"text": text, "token": token},
                "simplified": {"text": simplified, "token": token}
            }
            send_result_to_kafka(result, correlation_id=correlation_id)
