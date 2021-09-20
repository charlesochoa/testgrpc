# Prueba de concepto de uso de gRPC y Golang para manejo de notificaciones



## Para actualizar proto:
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative notification/notification.proto


## Computer login
gcloud auth login

## Build client docker image 
pack build --builder gcr.io/buildpacks/builder:v1

docker run --rm -p 8080:8080 client

## Testing cloud run access
grpcurl -proto grpc-protos/echo.proto -d '{"request": "Request Content"}' test-grpc-malvira-wibdbssi3q-ew.a.run.app:8080 echo.DirectEcho