# GO Kafka Example

This is repository how to use kafka in golang, with shopify sarama

## Tutorial

You can see tutorial to make this repository in [here](https://medium.com/easyread/implementasi-kafka-menggunakan-golang-testing-db183e0b3c29)

## How to run

```bash
# clone this repo
git clone git@github.com:mufti1/kafka-example.git
# or
git clone https://github.com/mufti1/kafka-example.git

# Add kafka to your hosts
sudo nano /etc/hosts
127.0.0.1 kafka

cd kafka-example

# running docker compose
docker-compose up

# run consumer with command
go run cmd/consumer/main.go

# run publisher with command
go run cmd/publiser/main.go
```
