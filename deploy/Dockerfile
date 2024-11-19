## Change this Dockerfile so the image size wont be huge
FROM golang:alpine3.20 

RUN apk update && apk add --no-cache git nodejs npm curl ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY package.json package-lock.json vite.config.js tailwind.config.js postcss.config.js ./

RUN npm install

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./app/cmd 

RUN npm run build

EXPOSE 8080

CMD ["./bin/main"]
