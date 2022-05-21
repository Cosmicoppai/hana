FROM golang:1.8

RUN mkdir src

WORKDIR src

COPY . .

RUN chmod +x ./setup.sh && ./setup.sh

EXPOSE 8096

CMD ["go", "run", "cmd/server/main.go", "--port", "8096"]