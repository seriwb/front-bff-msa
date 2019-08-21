FROM node:12.8-slim

EXPOSE 3000

RUN apt-get update \
    && apt-get install -y git \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /web-app/front

COPY . .

RUN yarn install

# CMD [ "yarn", "run", "dev" ]