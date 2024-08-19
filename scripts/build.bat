set GOOS=linux
set GOARCH=arm64
set CGO_ENABLED=0

echo Building create_device...
cd %APP_PATH%\handlers\createDevice\main
go build -tags lambda.norpc -o bootstrap create_device.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o create_device.zip bootstrap

echo Building get_device...
cd %APP_PATH%\handlers\getDevice\main
go build -tags lambda.norpc -o bootstrap get_device.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o get_device.zip bootstrap

echo Building update_device...
cd %APP_PATH%\handlers\updateDevice\main
go build -tags lambda.norpc -o bootstrap update_device.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o update_device.zip bootstrap

echo Building delete_device...
cd %APP_PATH%\handlers\deleteDevice\main
go build -tags lambda.norpc -o bootstrap delete_device.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o delete_device.zip bootstrap

echo Building process_sqs...
cd %APP_PATH%\handlers\processSqs\main
go build -tags lambda.norpc -o bootstrap process_sqs.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o process_sqs.zip bootstrap
