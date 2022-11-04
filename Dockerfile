FROM node:lts-alpine3.15 AS vuebuilder

COPY frontend/shortener-fe/ /usr/bin/shortener-fe

WORKDIR /usr/bin/shortener-fe

RUN npm install && npm run build

FROM golang:alpine3.15 AS release

COPY backend/shortenerBE/ /usr/bin/shortenerBE

WORKDIR /usr/bin/shortenerBE

RUN go mod tidy

COPY --from=vuebuilder /usr/bin/shortener-fe/dist ./views

# RUN go build -o /main
# CMD ["go", "run", "main.go"]

ENTRYPOINT ["go", "run", "main.go"]
# ENTRYPOINT ["/bin/sh"]