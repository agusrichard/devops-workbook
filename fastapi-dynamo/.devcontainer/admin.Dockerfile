FROM node:12.16.1

WORKDIR /app

RUN npm install -g dynamodb-admin

CMD ["dynamodb-admin"]