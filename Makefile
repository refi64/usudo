# Any copyright is dedicated to the Public Domain.
# http://creativecommons.org/publicdomain/zero/1.0/

$(shell mkdir -p build)

GO=go build -o $@ $^

USUDO=build/usudo.exe build/usudo-w.exe
HELPER=build/usudo-helper.exe

.PHONY: all strip

all: $(USUDO) $(HELPER)

$(USUDO): build/%.exe: %.go run.go shell32.go
	$(GO)

$(HELPER): usudo-helper.go
	$(GO)
