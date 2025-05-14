import requests

# Замените на адрес вашего API Gateway
URL = "http://localhost:8080/history"

# Токен должен быть реальным — тот, который есть в Redis
TOKEN = "fdb459f2-0cfe-4fd1-b412-d311fcbd5284"

headers = {
    "Authorization": f"Bearer {TOKEN}"
}

response = requests.get(URL, headers=headers)

print(f"Status code: {response.status_code}")
print("Response JSON:", response.json())
