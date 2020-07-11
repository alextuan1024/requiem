.SILENT:
.PHONY:

WITH_ENV =

NAME:=requiem
ROOF:=github.com/alextuan1024/$(NAME)
DATE:=$(shell date '+%Y%m%d')
TAG:=$(shell git describe --tags --always)
LDFLAGS:= -X $(ROOF)/cmd.version=$(TAG)
GO=$(shell which go)

vet:
	echo "checking ./..."
	$(GO) vet ./...

install: vet
	echo "build and install"
	$(GO) install -ldflags "$(LDFLAGS)"

build: vet
	echo "building..."
	$(GO) build -ldflags "-X $(ROOF)/cmd.version=debug" -o $(NAME)
