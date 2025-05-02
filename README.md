# tech-doc-simplifier
Преобразование технической документации в пользовательские инструкции

docker-compose up --build
docker restart auth-service

Добавление нового топика:
1.
docker exec -it broker /bin/bash

2.просмотр всех топиков:
/opt/kafka/bin/kafka-topics.sh \
>   --list \
>   --bootstrap-server broker:29092

3.
/opt/kafka/bin/kafka-topics.sh \
>   --create \
>   --topic user_requests \
>   --bootstrap-server broker:29092 \
>   --partitions 1 \
>   --replication-factor 1