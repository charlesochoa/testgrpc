# Prueba de concepto de uso de gRPC y Golang para manejo de notificaciones



## Para actualizar proto:
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative notification/notification.proto