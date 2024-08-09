FROM golang

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o ./bin/server ./cmd/app.go

EXPOSE 8082

CMD [ "./bin/server" ]