# Chat

## Description

Chat connecting users identified by username.

```
// duda
To: mbampi
Message: Bom dia!

mbampi: Buenoss
mbampi: Dormiu bem?
```
```
// mbampi
duda: Buenoss

To: duda
Message: Buenoss
To: duda
Message: Dormiu bem?
```

A client-server TCP connection, handling multiple clients using goroutines and channels.

## Usage
- Clone this repository
  - ```git clone https://github.com/mbampi/ConcurrentTCPServer.git``` 
  - `cd ConcurrentTCPServer`
  
- Run the Server
  - `go run Server/main.go` 
  
- Run Client (possible to run multiple clients, each client an user)
  - `go run Client/main.go` 