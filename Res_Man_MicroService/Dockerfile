FROM golang
MAINTAINER sourav241196@gmail.com
WORKDIR /app
COPY . .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify
RUN go build -v -o ./Res_Man_MicroService/bin/server ./Res_Man_MicroService/cmd/
EXPOSE 4000
CMD ["/app/bin/server"]
