
# ==============================================================================
# Tools commands

test-local:
	go mod tidy -compat=1.17
	go test -tags dynamic ./... -coverprofile coverage.out
	go tool cover -func=coverage.out


cover-local:
	go test ./... -coverprofile coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
