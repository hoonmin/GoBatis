pushd cmd\gobatis
go install
popd
go generate ./...
go test -v ./...