FROM node:17.0.0
RUN apt update && apt upgrade -y
WORKDIR /app
COPY package.json /app/package.json
RUN chown node:node /app
USER node
RUN npm install 
COPY --chown=node . /app
EXPOSE 3000
