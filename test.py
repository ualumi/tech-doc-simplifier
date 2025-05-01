import requests

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
print("Response Body:", response.text)

#faa6d19245ab74032ecc9b95ef3a07757b837c2231492983c964b57cf5bff469
#kafka-net id