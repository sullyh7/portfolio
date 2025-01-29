FROM golang:1.22-alpine as builder

WORKDIR /app

RUN apk add --no-cache make nodejs npm git

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN make install
RUN make build

RUN make templ

RUN touch /app/.env

FROM builder

COPY --from=builder /app/bin/main /main

# COPY --from=builder /app/.env .env

EXPOSE 3000

ENTRYPOINT [ "/portfolio" ]
