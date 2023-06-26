#File Transfer gRPC

## File Structure
```
file-transfer-service/
  ├── client/
  │   ├── client.go
  │   └── go.mod
  ├── server/
  │   ├── server.go
  │   └── go.mod
  ├── file_transfer.proto
  ├── go.mod
  └── README.md
```

## How to run
Simply execute the binaries in two different terminal instances

``` zsh
./server/server
```

``` zsh
./client/client
```

