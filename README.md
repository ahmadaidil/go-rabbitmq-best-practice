# go-rabbitmq-best-practice

## run as publisher
```
$ go run publisher/publisher.go
```

publish message example
- user queue
```
curl --location --request POST 'localhost:8081/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "John Doe",
    "email": "john@doe.com"
}'
```

- product queue
```
curl --location --request POST 'localhost:8081/product' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Indomie Goreng",
    "price": 3500
}'
```

## run as consumer
```
$ go run consumer/consumer.go
```

received message log example
- user
```
2022/03/05 21:12:22 Got message from queue "user": {"name":"John Doe","email":"john@doe.com"}
```

- product
```
2022/03/05 21:13:24 Got message from queue "product": {"name":"Indomie Goreng","price":3500}
```
