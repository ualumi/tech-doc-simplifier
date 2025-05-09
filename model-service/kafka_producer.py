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
        # можно также задать value_serializer здесь, но сейчас оставим явное кодирование
    )

    # Используем ensure_ascii=False, чтобы сохранить кириллицу как есть
    message = json.dumps(result, ensure_ascii=False).encode('utf-8')

    print(f"Sending model response to topic: key={correlation_id}, value={json.dumps(result, ensure_ascii=False)}")
    producer.send(
        RESPONSE_TOPIC,
        key=correlation_id.encode('utf-8'),
        value=message
    )


