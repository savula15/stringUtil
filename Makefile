RM=rm -f
RMFORCE=rm -rf

CODE_BASE=$(abspath $(GOPATH))
DESTDIR=$(CODE_BASE)/bin

SRCS="github.com/stringUtil/cmd/stringUtil"
COMP_NAME=stringUtil

all: clean format-code exe

.PHONY: all

clean:
	$(RM) $(DESTDIR)/$(COMP_NAME)

format-code:
	go fmt ./...

exe:
	@echo "Createing build for stringUtil"
	$(GO_BUILD_PREFIX) GOOS=linux go build -i -o $(DESTDIR)/$(COMP_NAME) $(SRCS)
