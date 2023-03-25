### RATE LIMITER GO

#### Tech Stack Used
- Http package 
- Dotenv

### Setup on linux machine
```
git clone https://github.com/axrav/_goratelimiter
cd _goratelimiter
echo PORT="8080" >> .env
go run ./cmd/. -N 3 # 3 is the number of requests per second passed through command args
```
### Build
```
git clone https://github.com/axrav/_goratelimiter
cd _goratelimiter
echo PORT="8080" >> .env
go build ./cmd/main.go
./main -N 3 # 3 is the number of requests per second passed through command args

```
### Run Tests
```
git clone https://github.com/axrav/_goratelimiter
cd _goratelimiter
go test ./tests/...
```


