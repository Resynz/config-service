SHELL=/bin/bash

EXE = config-service

all: $(EXE)

config-service:
	@echo "building $@ ..."
	$(MAKE) -s -f make.inc s=static

clean:
	rm -f $(EXE)