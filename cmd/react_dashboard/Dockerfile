FROM node:17.0.0
RUN apt update && apt upgrade -y
WORKDIR /app
COPY package.json /app/package.json
RUN chown node:node /app
RUN chown -R node:node /app 
USER node
RUN npm install 
COPY keys /app/keys
COPY --chown=node . /app
EXPOSE 3000
