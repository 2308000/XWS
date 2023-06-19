module accommodation_booking/api_gateway

go 1.20

replace accommodation_booking/common => ../common

require (
	accommodation_booking/common v1.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/joho/godotenv v1.4.0
	github.com/rs/cors v1.9.0
	go.mongodb.org/mongo-driver v1.9.1
	google.golang.org/grpc v1.47.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20220617124728-180714bec0ad // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
