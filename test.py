#тест запроса на упрощение по http
'''import requests

# URL микросервиса (если локально запущен на 8080)
url = "http://localhost:8080/simplify"

# Заголовки с авторизацией (микросервис ожидает Authorization)
headers = {
    "Content-Type": "application/json",
    "Authorization": "Bearer test-token"
}

# Тело запроса
payload = {
    "text": "Это пример сложного технического текста"
}

# Отправка запроса
response = requests.post(url, json=payload, headers=headers)

# Вывод результата
print("Status Code:", response.status_code)
print("Response Body:", response.text)'''

#faa6d19245ab74032ecc9b95ef3a07757b837c2231492983c964b57cf5bff469
#kafka-net id

'''go mod init api-gateway
go get github.com/gorilla/mux
go get github.com/segmentio/kafka-go
go mod tidy

go mod init auth-service
go get github.com/segmentio/kafka-go
go get github.com/go-redis/redis/v8
go get github.com/lib/pq
'''
#тест регистрации по http
'''import requests

url = "http://localhost:8081/registration"
payload = {
    "email": "test@example.com",
    "login": "testuser",
    "password": "secret123"
}

response = requests.post(url, json=payload)

print("Status Code:", response.status_code)
print("Response Body:", response.text)'''

#тест login по http
'''import requests

url = "http://localhost:8081/login"
payload = {
    "login": "testuser",
    "password": "secret123"
}

response = requests.post(url, json=payload)

print("Status Code:", response.status_code)
print("Response Body:", response.text)'''

import requests

import requests

url = "http://localhost:8080/simplify"
headers = {
    "Authorization": "b8446067-8f5f-4a9d-a87b-3e0619e4630e",
    "Content-Type": "application/json"
}
data = {
    "text": "Какой-то технический текст, который нужно упростить"
}

response = requests.post(url, headers=headers, json=data)
print("Status:", response.status_code)
print("Response text:", response.text)

#S/opt/kafka/bin/kafka-broker-api-versions.sh --bootstrap-server broker:29092

'''/opt/kafka/bin/kafka-topics.sh \
  --create \
  --topic user_responses \
  --bootstrap-server broker:29092 \
  --partitions 1 \
  --replication-factor 1'''

'''import requests

url = "http://localhost:8081/registration"
data = {
    "email": "test@example.com",
    "login": "testuser",
    "password": "securepassword123"
}

response = requests.post(url, json=data)

print("Status Code:", response.status_code)
print("Response Text:", response.text)'''

'''import requests

url = "http://localhost:8081/login"
data = {
    "login": "testuser",
    "password": "securepassword123"
}

response = requests.post(url, json=data)

print("Status Code:", response.status_code)
print("Response JSON:", response.json() if response.status_code == 200 else response.text)'''
