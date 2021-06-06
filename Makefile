.PHONY: all install clean re

GONAME		= n-puzzle

GOPATH		= $(shell pwd)
GOBIN		= $(GOPATH)/bin
GOENV		= GOPATH=$(GOPATH) GOBIN=$(GOBIN)

FILES		= src/main.go \
				src/a_impliment.go \
				src/parse.go \
				src/exist_solution.go\
				src/hueristic.go \
				src/utils.go \
				src/parse_flag.go \

EXEPATH		= ./bin/$(GONAME)

all:		$(EXEPATH)

$(EXEPATH):	$(FILES)
			$(GOENV) go build -o $(EXEPATH) $(FILES)

get:
		$(GOENV) go get .

install:
		$(GOENV) go install $(GOFILES)

clean:
		$(GOENV) go clean
		rm -rf ./bin/

re:		fclean all

