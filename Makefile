all: install

install:
	@for dir in $(GO_SRC_DIRS); do \
		go install $$dir; \
	done;
	
help:
	@echo "Acciones válidas:"
	@echo "	-all"
	@echo "	-install"
	@echo "	-test"
	@echo "	-fmt"


# Lanzo los test con los resultados de la búsqueda de archivos de test
test:
	@for dir in $(GO_TEST_DIRS); do\
		go test -v $$dir; \
	done;
	
fmt: 
	@for dir in $(GO_SRC_DIRS); do \
		go fmt $$dir; \
	done;

# Busco los archivos .go
GO_SRC_DIRS := $(shell \
	find . -name "*.go" | \
	xargs -I {} dirname {}  | \
	uniq)	
	
# Busco los archivos _test.go
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go"| \
	xargs -I {} dirname {}  | \
	uniq)
	


.PHONY: test
