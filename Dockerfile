FROM golang:1.22

WORKDIR /app
COPY . .
RUN go mod tidy

RUN go build -o coffeeshop .

EXPOSE 8080

RUN chmod a+x coffeeshop

CMD [ "./coffeeshop" ]