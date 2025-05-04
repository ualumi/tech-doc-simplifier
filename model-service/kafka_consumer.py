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
        if text:
            simplified = simplify_text(text)
            send_result_to_kafka({"original": text, "simplified": simplified})
