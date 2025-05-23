# tech-doc-simplifier
Запуск контейнера:
1. docker-compose up --build

2. docker restart api-gateway auth-service text-service model-service result-writer
(необходимо перезапустить после запуска брокера сообщений)

3. docker exec -it broker kafka-topics --create --topic result_response --bootstrap-server broker:9092 --partitions 1 --replication-factor 1

Преобразование технической документации в пользовательские инструкции

Модель на Hugginface: https://huggingface.co/disemenova/proekt_sum/tree/main




