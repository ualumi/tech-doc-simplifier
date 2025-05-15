import requests

url = "http://localhost:8081/registration"
data = {
    "email": "test@example1.com",
    "login": "test",
    "password": "securepassword1231"
}

response = requests.post(url, json=data)

print("Status Code:", response.status_code)
print("Response Text:", response.text)

import requests

url = "http://localhost:8081/login"
data = {
    "login": "test",
    "password": "securepassword1231"
}

response = requests.post(url, json=data)

print("Status Code:", response.status_code)
print("Response JSON:", response.json() if response.status_code == 200 else response.text)