import requests

# Замените на адрес вашего API Gateway
URL = "http://localhost:8080/history"

# Токен должен быть реальным — тот, который есть в Redis
TOKEN = "083a4abf-ffe0-4c81-8a19-9e0b78a748b7"

headers = {
    "Authorization": f"Bearer {TOKEN}"
}

response = requests.get(URL, headers=headers)

print(f"Status code: {response.status_code}")
print("Response JSON:", response.json())
