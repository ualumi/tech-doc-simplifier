import os
from kafka import KafkaProducer
import json

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "broker:9092")
RESPONSE_TOPIC = os.getenv("KAFKA_RESPONSE_TOPIC", "model_response")

producer = KafkaProducer(
    bootstrap_servers=KAFKA_BROKER,
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)

def send_result_to_kafka(result: dict, correlation_id: str):
    from kafka import KafkaProducer
    import json

    producer = KafkaProducer(
        bootstrap_servers=KAFKA_BROKER,
        # без value_serializer
    )

    print(f"Sending model response to topic: key={correlation_id}, value={result}")
    producer.send(
        RESPONSE_TOPIC,
        key=correlation_id.encode('utf-8'),
        value=json.dumps(result).encode('utf-8')
    )


