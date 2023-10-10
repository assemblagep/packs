Unit tests
```
go test -v ./...
```

Run web server
```
go run .
```

Send request
```
curl -d packs=1,500,1000,250,2000 "localhost:8081/amount"
curl -d amount=1001 "localhost:8081/amount"
```

Calculation result
```
map[pack1:amount1 pack2:amount2 ...]
```