ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: clean
clean:
	rm -rf build

.PHONY: test
test:
	cd brewing_support && \
	go test ./... && \
	cd $(ROOT_DIR)

.PHONY: build
build: clean
	mkdir build && \
	cd brewing_support && \
	export GOOS=linux && \
	go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o ../build/brewing_support_linux_amd64.bin . && \
	export GOOS=windows && \
	go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o ../build/brewing_support_windows_amd64.exe . && \
	cd $(ROOT_DIR) && \
	cp -r resources build/

.PHONY: migrateUP
migrateUP:
	migrate -path resources/migrations -database postgres://brewing_support:brewing_support@brewing_support_postgres_db:5432/brewing_supportdb?sslmode=disable up

.PHONY: migrateDOWN
migrateDOWN:
	migrate -path resources/migrations -database postgres://brewing_support:brewing_support@brewing_support_postgres_db:5432/brewing_supportdb?sslmode=disable down

.PHONY: migrateCreate
migrateCreate:
	ARG=""
	migrate create -ext sql -dir migrations -seq ${ARG}