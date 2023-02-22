# rabbitMQ-go
```
1.Download RabbitMq exe for windows ->https://www.rabbitmq.com/install-windows.html
2.Download erlang -> https://www.erlang.org/downloads
3.Install both
4.Go to rabbitMQ folder and then in sbin and open cmd and run the command
Enable plugins -> rabbitmq-plugins enable rabbitmq_management
5. Restart the system.
6.http://localhost:15672
```


```
curl --location 'http://localhost:9000/rabbitMQ/send_to_queue' \
--header 'Content-Type: application/json' \
--data '{
    "msg": {
        "contactNumber": "1234567891",
        "message": "Hello Rk",
        "message2":"Hello"
    }
}'
```