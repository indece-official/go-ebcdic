# Project config
BUILD_VERSION ?= SNAPSHOT

all: build_ebcdictoutf8 build_utf8toebcdic

build_ebcdictoutf8:
	cd ./cmd/ebcdictoutf8 && BUILD_VERSION=${BUILD_VERSION} make --always-make

build_utf8toebcdic:
	cd ./cmd/utf8toebcdic && BUILD_VERSION=${BUILD_VERSION} make --always-make
