FROM golang:1.18

RUN mkdir app

WORKDIR app

COPY src .

RUN chmod +x ./setup.sh && ./setup.sh

EXPOSE 8096

CMD ["go", "run", "cmd/server/main.go", "--port", "8096"]