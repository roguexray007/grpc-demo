# grpc-demo
trying out grpc

# Command to generate pb.go and _grpc.pb.go 
```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto```

# running server
```
    go run greeter_server/main.go --port=50052
```    
or    
```
    go run greeter_server/main.go --port=50052
```

# running client
```
    go run greeter_client/main.go --addr=localhost:50052 --name roguexray007
```    
or
```
    go run greeter_client/main.go
```
