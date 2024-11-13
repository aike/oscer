oscer: oscer.go osc/osc.go
	go build oscer.go

test:
	go test -v oscer_test.go

all: oscer.go osc/osc.go
	set GOOS=darwin
	set GOARCH=arm64
	go build -o ../build/MacArm/oscer oscer.go 
	set GOOS=darwin
	set GOARCH=amd64
	go build -o ../build/MacAmd/oscer oscer.go 
	set GOOS=windows
	set GOARCH=amd64
	go build -o ../build/Windows64/oscer.exe oscer.go 
	set GOOS=windows
	set GOARCH=386
	go build -o ../build/Windows32/oscer.exe oscer.go 
	set GOOS=linux
	set GOARCH=amd64
	go build -o ../build/Linux64/oscer oscer.go 
	set GOOS=linux
	set GOARCH=386
	go build -o ../build/Linux32/oscer oscer.go 
	set GOOS=linux
	set GOARCH=arm
	go build -o ../build/LinuxARM/oscer oscer.go 
