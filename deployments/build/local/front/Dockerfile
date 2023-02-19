FROM node:16

WORKDIR /app/frontend

COPY package*.json ./
RUN npm install -g @vue/cli

COPY . .

EXPOSE 8080

CMD ["npm", "run", "serve"]