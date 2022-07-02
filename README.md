# Poker Equity Calculator

## Build
```
go build .
```


## Usage Examples
```
./go-poker-equity --iter 1 6s9c4hQcKd KsKc 2s3h
{"equity":{"KsKc":1},"time_delta":0,"iterations":1}
```
```
./go-poker-equity --iter 1000 6s9c4hQcKd 6c8d KsKh:0.3,2h3s
{"equity":{"6c8d":0.767},"time_delta":0.001,"iterations":1000
```
```
./go-poker-equity . --timeout 0.1 6s9c4hQcKd 6c8d KsKh:0.3,2h3s
{"equity":{"6c8d":0.7677752},"time_delta":0.1,"iterations":123501}
```

## Run tests
```
go test ./...
```