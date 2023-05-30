default:build ver

generate:
	@go generate ./...

optimize:
	@#go get github.com/orijtech/structslop/cmd/structslop
	@structslop -verbose ./...

build:clean generate
	@./build.sh

linux:clean generate
	@LINUX=1 ./build.sh

docker:
	@docker build -t flexdb .
	@docker image list

dockerboot:docker dockermysql
	@docker run --user 10001 -p 10002:8000 -d flexdb
	@#docker run --link flexsql:mysql --user 10001 -p 10002:8000 flexdb -migrate -dsn "root:@tcp(mysql:3306)/easyapp"
	@#docker run --user 10001 -p 10002:8000 -d flexdb -dsn "root:@tcp(mysql:3306)/easyapp"

dockerclean:
	@echo 'docker rm --force $$(docker ps -a -q)'
	@echo 'docker rmi $$(docker images -f "dangling=true" -q)'

dockermysql:
	docker run -d --name flexsql -e MYSQL_DATABASE=easyapp -e MYSQL_ALLOW_EMPTY_PASSWORD=1 mysql:5.7

it:
	docker run -it -u 10001 flexdb /bin/sh

dockerpush:
	@docker push dddplus/flexdb

clean:
	@go clean ./cmd/flexdb/
	@rm -f flexdb
	@rm -f *.log.*
	@rm -f *.log
	@find . -name bindata.go -exec rm -f {} \;

ver:build
	@./flexdb -ver

apis:build
	@./flexdb -apis -cf conf/flexdb.cf

pprof:
	@open http://localhost:10120/debug/pprof/

run:build
	@./flexdb -loglevel=debug -cf conf/flexdb.cf # -logfile f.log

fmt:
	@go fmt ./...

test:
	@#go test -covermode=atomic -bench=. -race ./pkg/...	
	@go test -covermode=atomic ./pkg/...

resetdb:build
	@./flexdb -migrate=true -cf conf/flexdb.cf

deploy:linux
	@scp flexdb root@11.51.197.168:/home/admin/workspace/flexdb
