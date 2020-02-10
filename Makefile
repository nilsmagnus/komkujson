SHELL := /bin/bash
.PHONY=start stop watch-test

ifndef VERBOSE
.SILENT:
endif

APP=komkujson

$(APP): *.go 
	@echo "Building backend"
	@go build -o $@
	@echo "Built backend $@ "


.system: $(APP) 
	touch .system

test: *.go
	go test
	touch test


.restartServer: *.go start
	touch .restartServer

start: server.PID

server.PID: .system
	killall -SIGINT $(APP) || echo "nothing to kill"
	kill `cat server.PID` 2> /dev/null || echo "Nothing to kill before starting $(APP)"
	./$(APP) & echo $$! > $@; 
	echo "Started $(APP) on with PID `cat server.PID` "

watch: clean server.PID
	while true ; do \
		date; \
		make .restartServer ;\
		inotifywait -qre close_write . ; \
	done

clean:
	killall $(APP) 2> /dev/null || echo ""
	rm -f .restartServer server.PID $(APP) test 

