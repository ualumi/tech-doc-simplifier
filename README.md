# tech-doc-simplifier
Преобразование технической документации в пользовательские инструкции

docker-compose up --build
docker restart auth-service

Добавление нового топика:
1.
docker exec -it broker /bin/bash

2.просмотр всех топиков:
/opt/kafka/bin/kafka-topics.sh \
  --list \
  --bootstrap-server broker:29092

3.
/opt/kafka/bin/kafka-topics.sh \
  --create \
  --topic user_requests \
  --bootstrap-server broker:29092 \
  --partitions 1 \
  --replication-factor 1


новій топик на ковой кафке
docker exec -it broker bash


kafka-topics --create \
   --bootstrap-server localhost:9092 \
   --replication-factor 1 \
   --partitions 1 \
   --topic model_requests


docker logs auth-service

kafka-consumer-groups.sh --bootstrap-server broker:29092 --describe --group auth-service-group

docker exec -it 34ffe95dd0c555103e651dd6164574d3b0d6808d81ccd0a1d93bcc35b2a87f17 kafka-consumer-groups.sh --bootstrap-server broker:29092 --describe --group auth-service-group

34ffe95dd0c555103e651dd6164574d3b0d6808d81ccd0a1d93bcc35b2a87f17
