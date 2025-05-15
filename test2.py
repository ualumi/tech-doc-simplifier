import requests

# Замените на адрес вашего API Gateway
URL = "http://localhost:8080/history"

# Токен должен быть реальным — тот, который есть в Redis
TOKEN = "dfd198d1-82c0-46e6-af56-494795fc9faf"

headers = {
    "Authorization": f"Bearer {TOKEN}"
}

response = requests.get(URL, headers=headers)

print(f"Status code: {response.status_code}")
print("Response JSON:", response.json())
