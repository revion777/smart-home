echo "Running tests..."
cd %APP_PATH%
set GOOS=windows
set GOARCH=amd64
go clean -testcache
go test -v ./tests/...
echo "Tests completed."