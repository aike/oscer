oscer: oscer.go osc/osc.go
	go build oscer.go

test:
	cd ../test; go test -v

all: oscer.go osc/osc.go
	GO111MODULE=off GOOS=darwin  GOARCH=amd64 go build -o ../build/MacOSX/oscer oscer.go 
	GO111MODULE=off GOOS=darwin  GOARCH=arm64 go build -o ../build/MacOSX_M1/oscer oscer.go 
	GO111MODULE=off GOOS=windows GOARCH=amd64 go build -o ../build/Windows64/oscer.exe oscer.go 
	GO111MODULE=off GOOS=windows GOARCH=386   go build -o ../build/Windows32/oscer.exe oscer.go 
	GO111MODULE=off GOOS=linux   GOARCH=amd64 go build -o ../build/Linux64/oscer oscer.go 
	GO111MODULE=off GOOS=linux   GOARCH=386   go build -o ../build/Linux32/oscer oscer.go 
	GO111MODULE=off GOOS=linux   GOARCH=arm   go build -o ../build/LinuxARM/oscer oscer.go 
