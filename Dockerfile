FROM golang:1.17
WORKDIR /app
COPY . .

RUN go mod init
RUN go get github.com/labstack/echo/v4
RUN go get -u gorm.io/gorm
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/joho/godotenv
RUN go go get github.com/lib/pq

CMD ["go", "run", "main.go"]