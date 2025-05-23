import socket
from kafka_producer import init_producer
from kafka_consumer import consume

def check_kafka_connection(host, port):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        try:
            s.connect((host, port))
            print("Kafka доступен")
            return True
        except Exception as e:
            print(f"Kafka недоступен: {e}")
            return False

if __name__ == "__main__":
    if check_kafka_connection("broker", 9092):
        init_producer()
        consume()
    else:
        print("Kafka не доступен, приложение не запущено.")