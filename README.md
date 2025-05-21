# tech-doc-simplifier
Преобразование технической документации в пользовательские инструкции

Модель на Hugginface: https://huggingface.co/disemenova/proekt_sum/tree/main

docker-compose up --build

docker restart api-gateway auth-service text-service model-service result-writer

docker exec -it broker kafka-topics --create --topic result_response --bootstrap-server broker:9092 --partitions 1 --replication-factor 1


