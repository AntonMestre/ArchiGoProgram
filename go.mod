module github.com/AntonMestre/AirportProject

go 1.17

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	go.mongodb.org/mongo-driver v1.8.1
	util v1.0.0
)

require (
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace util v1.0.0 => ./internal/util
