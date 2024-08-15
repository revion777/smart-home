set GOOS=linux
set GOARCH=arm64

echo Building create_device...
cd cmd/create_device
go build -o bootstrap main.go
build-lambda-zip.exe -o create_device.zip bootstrap
cd ../..

echo Building get_device...
cd cmd/get_device
go build -o bootstrap main.go
zip get_device.zip bootstrap
cd ../..

echo Building update_device...
cd cmd/update_device
go build -o bootstrap main.go
zip update_device.zip bootstrap
cd ../..

echo Building delete_device...
cd cmd/delete_device
go build -o bootstrap main.go
zip delete_device.zip bootstrap
cd ../..

echo Building process_sqs...
cd cmd/process_sqs
go build -o bootstrap main.go
zip process_sqs.zip bootstrap
cd ../..

echo Packaging AWS Layer...
cd layer/go
zip -r ../../smart-home-layer.zip *
cd ../..

echo Build process complete!
