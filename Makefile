# Makefile for building Kommunkod.Restclone
# Runs go build -buildmode=c-shared -o src/Kommunkod.Restclone/restclone.dll cmd/csharp/
all: go dotnet
	@echo "Building Everything"


go:
	@echo "Building Go Shared Library"
	go build -buildmode=c-shared -o src/Kommunkod.Restclone/restclone.dll ./cmd/cshared/

dotnet:
	@echo "Building C# Test Program"
	cd testing/csharp && dotnet build -c Release

package:
	@echo "Packaging C# Test Program"
	dotnet build -c Release -o ./src/Kommunkod.Restclone/obj/Kommunkod.Restclone
	dotnet pack ./src/Kommunkod.Restclone/ -c Release -o ./src/Kommunkod.Restclone/obj/Kommunkod.Restclone.nupkg

