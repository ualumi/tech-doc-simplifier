FROM python:3.13-slim

WORKDIR /app



COPY requirements.txt .
ENV PIP_DEFAULT_TIMEOUT=100
ENV PIP_NO_CACHE_DIR=1

RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    pkg-config \
    git \
    wget \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN pip install --default-timeout=100 -r requirements.txt


COPY . .

CMD ["python", "app.py"]
