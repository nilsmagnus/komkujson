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


test: *.go
	go test
	touch test


server.PID: $(APP)  
	kill `cat server.PID` 2> /dev/null || echo "Nothing to kill before starting $(APP)"
	./$(APP) & echo $$! > $@; 
	echo "Started $(APP) on with PID `cat server.PID` "


watch: clean server.PID
	while true ; do \
		date; \
		make server.PID ;\
		inotifywait -qre close_write . ; \
	done

clean:
	killall $(APP) 2> /dev/null || echo ""
	rm -f .restartServer server.PID $(APP) test 

