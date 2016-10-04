RM=rm -f
RMFORCE=rm -rf
DESTDIR=$(SR_CODE_BASE)/snaproute/src/out/bin
SRCS=main.go
COMP_NAME=client_thrift
all: exe

exe: $(SRCS)
	go build -o $(COMP_NAME) $(SRCS)

guard:
ifndef SR_CODE_BASE
	$(error SR_CODE_BASE is not set)
endif

install:
	@echo "Test thrift has no files to install"
clean:guard
	$(RM) $(COMP_NAME) 
	#$(RMFORCE) $(GENERATED_IPC)/$(COMP_NAME)
