FROM node:15.11.0-alpine3.10

WORKDIR /app/frontend

# COPY package*.json ./

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli

# COPY . .

CMD ["npm", "run", "serve"]

# CMD ["/bin/ash"]