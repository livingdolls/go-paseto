run :
	go run cmd/http/runner.go

test :
	go test -v -cover ./...

mock :
	mockgen -package mockdb -destination internal/infra/mock/user_mock.go github.com/livingdolls/go-paseto/internal/core/port/repository/user UserPortRepository

.PHONY : run