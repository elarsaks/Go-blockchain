# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.18.0

WORKDIR /app 

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY . .

EXPOSE 3000

CMD "air"