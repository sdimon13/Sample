# Используем официальный образ node.js в качестве базового
FROM node:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы package.json и package-lock.json
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install
RUN npm install -g @vue/cli-service
RUN npm install -g @vue/cli-plugin-babel
RUN npm install @vue/cli-plugin-eslint

# Копируем остальные файлы проекта
COPY . .

# Запускаем проект
CMD [ "npm", "run", "serve" ]