gvm use go1.4 --default
gdb -tui interview
/home/jjosephy/.gvm/gos/go1.4/src/runtime/runtime-gdb.py
 go build -gcflags "-N -l" main.go
 go get github.com/julienschmidt/httprouter

Set GOPATH to the path of where you are doing go work (e.g. ~/Source/go)
RUN
 $GOPATH/bin/interview

BUILD
go build interview
go install -gcflags "-N -l" interview (DEBUG)

export GOPATH=$HOME/work
