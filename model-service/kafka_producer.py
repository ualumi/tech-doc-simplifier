'''import os
from kafka import KafkaProducer
import json

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "broker:9092")
RESPONSE_TOPIC = os.getenv("KAFKA_RESPONSE_TOPIC", "model_response")

# Убираем создание producer при импорте!
producer = None

def init_producer():
    global producer
    if producer is None:
        producer = KafkaProducer(
            bootstrap_servers=KAFKA_BROKER,
            value_serializer=lambda v: json.dumps(v, ensure_ascii=False).encode('utf-8')
        )

def send_result_to_kafka(result: dict, correlation_id: str):
    if producer is None:
        raise RuntimeError("Producer not initialized. Call init_producer() first.")
    
    # Преобразуем все значения bytes → str (рекурсивно)
    def decode_bytes(obj):
        if isinstance(obj, bytes):
            return obj.decode('utf-8')
        elif isinstance(obj, dict):
            return {k: decode_bytes(v) for k, v in obj.items()}
        elif isinstance(obj, list):
            return [decode_bytes(item) for item in obj]
        return obj

    result = decode_bytes(result)

    print(f"Sending model response to topic: key={correlation_id}, value={json.dumps(result, ensure_ascii=False)}")
    
    producer.send(
        RESPONSE_TOPIC,
        key=correlation_id.encode('utf-8'),
        value=json.dumps(result, ensure_ascii=False).encode('utf-8')
    )'''

'''import os
from kafka import KafkaProducer
import json

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "broker:9092")
RESPONSE_TOPIC = os.getenv("KAFKA_RESPONSE_TOPIC", "model_response")

# Убираем создание producer при импорте!
producer = None

def init_producer():
    global producer
    if producer is None:
        producer = KafkaProducer(
            bootstrap_servers=KAFKA_BROKER,
            value_serializer=lambda v: json.dumps(v, ensure_ascii=False)  # Убираем кодирование в байты
        )

# Функция для декодирования любых байтов в строку
def decode_bytes(obj):
    if isinstance(obj, bytes):
        return obj.decode('utf-8')
    elif isinstance(obj, dict):
        return {k: decode_bytes(v) for k, v in obj.items()}
    elif isinstance(obj, list):
        return [decode_bytes(item) for item in obj]
    return obj

def send_result_to_kafka(result: dict, correlation_id: str):
    if producer is None:
        raise RuntimeError("Producer not initialized. Call init_producer() first.")
    
    # Логирование до декодирования
    print(f"Before decoding: {result}")

    # Декодируем результат, чтобы избежать проблемы с bytes
    result = decode_bytes(result)

    # Логирование после декодирования
    print(f"After decoding: {result}")

    # Сериализуем результат и отправляем в Kafka (без кодирования в байты в value_serializer)
    print(f"Sending model response to topic: key={correlation_id}, value={json.dumps(result, ensure_ascii=False)}")
    
    producer.send(
        RESPONSE_TOPIC,
        key=correlation_id.encode('utf-8'),
        value=json.dumps(result, ensure_ascii=False)  # Без .encode('utf-8')
    )
'''



import os
from kafka import KafkaProducer
import json

KAFKA_BROKER = os.getenv("KAFKA_BROKER", "broker:9092")
RESPONSE_TOPIC = os.getenv("KAFKA_RESPONSE_TOPIC", "model_response")

# Убираем создание producer при импорте!
producer = None

def init_producer():
    global producer
    if producer is None:
        producer = KafkaProducer(
            bootstrap_servers=KAFKA_BROKER,
        )

def send_result_to_kafka(result: dict, correlation_id: str):
    if producer is None:
        raise RuntimeError("Producer not initialized. Call init_producer() first.")
    
    # Преобразуем все значения bytes → str (рекурсивно)
    def decode_bytes(obj):
        if isinstance(obj, bytes):
            return obj.decode('utf-8')
        elif isinstance(obj, dict):
            return {k: decode_bytes(v) for k, v in obj.items()}
        elif isinstance(obj, list):
            return [decode_bytes(item) for item in obj]
        return obj

    result = decode_bytes(result)

    print(f"Sending model response to topic: key={correlation_id}, value={json.dumps(result, ensure_ascii=False)}")
    
    producer.send(
        RESPONSE_TOPIC,
        key=correlation_id.encode('utf-8'),
        value=json.dumps(result, ensure_ascii=False).encode('utf-8')
    )


