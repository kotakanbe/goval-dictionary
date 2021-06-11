.PHONY: \
	build \
	install \
	all \
	vendor \
 	lint \
	vet \
	fmt \
	fmtcheck \
	pretest \
	test \
	integration \
	cov \
	clean

SRCS = $(shell git ls-files '*.go')
PKGS =  ./commands ./config ./db ./db/rdb ./fetcher ./models ./util
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'
GO := GO111MODULE=on go
GO_OFF := GO111MODULE=off go

all: build

build: main.go pretest
	$(GO) build -a -ldflags "$(LDFLAGS)" -o goval-dictionary $<

b: 	main.go pretest
	$(GO) build -ldflags "$(LDFLAGS)" -o goval-dictionary $<

install: main.go pretest
	$(GO) install -ldflags "$(LDFLAGS)"

lint:
	$(GO_OFF) get -u golang.org/x/lint/golint
	golint $(PKGS)

vet:
	echo $(PKGS) | xargs env $(GO) vet || exit;

fmt:
	gofmt -s -w $(SRCS)

mlint:
	$(foreach file,$(SRCS),gometalinter $(file) || exit;)

fmtcheck:
	$(foreach file,$(SRCS),gofmt -s -d $(file);)

pretest: lint vet fmtcheck

test: 
	$(GO) test -cover -v ./... || exit;

unused:
	$(foreach pkg,$(PKGS),unused $(pkg);)

cov:
	@ go get -v github.com/axw/gocov/gocov
	@ go get golang.org/x/tools/cmd/cover
	gocov test | gocov report

clean:
	echo $(PKGS) | xargs go clean || exit;
	echo $(PKGS) | xargs go clean || exit;

PWD := $(shell pwd)
BRANCH := $(shell git symbolic-ref --short HEAD)
build-integration:
	# @ git stash save
	$(GO) build -ldflags "$(LDFLAGS)" -o integration/goval-dictionary.new
	@ git stash save
	git checkout $(shell git describe --tags --abbrev=0)
	@git reset --hard
	$(GO) build -ldflags "$(LDFLAGS)" -o integration/goval-dictionary.old
	git checkout $(BRANCH)
	@ git stash apply stash@{0} && git stash drop stash@{0}

clean-integration:
	-pkill goval-dictionary.old
	-pkill goval-dictionary.new
	-rm integration/goval-dictionary.old integration/goval-dictionary.new integration/oval.old.sqlite3 integration/oval.new.sqlite3
	-docker kill redis-old redis-new
	-docker rm redis-old redis-new

fetch-rdb:
	# integration/goval-dictionary.old fetch-debian --dbpath=$(PWD)/integration/oval.old.sqlite3 7 8 9 10
	# integration/goval-dictionary.old fetch-ubuntu --dbpath=$(PWD)/integration/oval.old.sqlite3 14 16 18 19 20
	integration/goval-dictionary.old fetch-redhat --dbpath=$(PWD)/integration/oval.old.sqlite3 5 6 7 8
	
	# integration/goval-dictionary.new fetch-debian --dbpath=$(PWD)/integration/oval.new.sqlite3 7 8 9 10
	# integration/goval-dictionary.new fetch-ubuntu --dbpath=$(PWD)/integration/oval.new.sqlite3 14 16 18 19 20
	integration/goval-dictionary.new fetch-redhat --dbpath=$(PWD)/integration/oval.new.sqlite3 5 6 7 8

fetch-redis:
	docker run --name redis-old -d -p 127.0.0.1:6379:6379 redis
	docker run --name redis-new -d -p 127.0.0.1:6380:6379 redis

	# integration/goval-dictionary.old fetch-debian --dbtype redis --dbpath "redis://127.0.0.1:6379/0"
	integration/goval-dictionary.old fetch-redhat --dbtype redis --dbpath "redis://127.0.0.1:6379/0"

	# integration/goval-dictionary.new fetch-debian --dbtype redis --dbpath "redis://127.0.0.1:6380/0"
	integration/goval-dictionary.new fetch-redhat --dbtype redis --dbpath "redis://127.0.0.1:6380/0"

diff-cveid:
	# @ python integration/diff_server_mode.py cveid debian 7 8 9 10
	# @ python integration/diff_server_mode.py cveid ubuntu 14 16 18 19 20
	@ python integration/diff_server_mode.py cveid redhat 5 6 7 8

diff-package:
	# @ python integration/diff_server_mode.py package debian 7 8 9 10
	# @ python integration/diff_server_mode.py cveid ubuntu 14 16 18 19 20
	@ python integration/diff_server_mode.py package redhat 5 6 7 8

diff-server-rdb:
	integration/goval-dictionary.old server --dbpath=$(PWD)/integration/oval.old.sqlite3 --port 1325 > /dev/null & 
	integration/goval-dictionary.new server --dbpath=$(PWD)/integration/oval.new.sqlite3 --port 1326 > /dev/null &
	make diff-cveid
	# make diff-package
	pkill goval-dictionary.old 
	pkill goval-dictionary.new

diff-server-redis:
	integration/goval-dictionary.old server --dbtype redis --dbpath "redis://127.0.0.1:6379/0" --port 1325 > /dev/null & 
	integration/goval-dictionary.new server --dbtype redis --dbpath "redis://127.0.0.1:6380/0" --port 1326 > /dev/null &
	make diff-cveid
	# make diff-package
	pkill goval-dictionary.old 
	pkill goval-dictionary.new

diff-server-rdb-redis:
	integration/goval-dictionary.new server --dbpath=$(PWD)/integration/oval.new.sqlite3 --port 1325 > /dev/null &
	integration/goval-dictionary.new server --dbtype redis --dbpath "redis://127.0.0.1:6380/0" --port 1326 > /dev/null &
	make diff-cveid
	# make diff-package
	pkill goval-dictionary.new