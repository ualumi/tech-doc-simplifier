#FROM node:23-alpine AS build

#WORKDIR /app

#COPY package*.json ./
#RUN npm install

#COPY . .
#RUN npm run build

# Установка сервера serve
#RUN npm install -g serve

#FROM nginx:stable-alpine

#COPY --from=build /dist /usr/share/nginx/html
#COPY --from=build nginx.conf /etc/nginx/conf.d/default.conf

#EXPOSE 3000
#CMD ["nginx", "-g", "daemon off;"]


# Stage 1: Build the app
FROM node:18 AS build

WORKDIR /app
COPY . .
RUN npm install && npm run build

# Stage 2: Serve with Nginx
FROM nginx:stable-alpine

# Копируем билд из dist, а не build
COPY --from=build /app/dist /usr/share/nginx/html

# Кастомная конфигурация Nginx (если у тебя она есть)
COPY nginx.conf /etc/nginx/conf.d/default.conf

