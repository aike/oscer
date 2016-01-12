oscer: oscer.go osc/osc.go
	go build oscer.go

test:
	cd ../test; go test -v

all: oscer.go osc/osc.go
	GOOS=darwin  GOARCH=amd64 go build -o ../build/MacOSX/oscer oscer.go 
	GOOS=windows GOARCH=amd64 go build -o ../build/Windows64/oscer.exe oscer.go 
	GOOS=windows GOARCH=386   go build -o ../build/Windows32/oscer.exe oscer.go 
	GOOS=linux   GOARCH=amd64 go build -o ../build/Linux64/oscer oscer.go 
	GOOS=linux   GOARCH=386   go build -o ../build/Linux32/oscer oscer.go 
	GOOS=linux   GOARCH=arm   go build -o ../build/LinuxARM/oscer oscer.go 
