# run the application
run:
    go run ./cmd/app/main.go

# tidy the application
tidy:
    go mod tidy

# isntall dependencies to local
vendor:
    go mod vendor


