# Any copyright is dedicated to the Public Domain.
# http://creativecommons.org/publicdomain/zero/1.0/

VER=0.1
ARCH?=x64

ifneq ($(filter-out x86 x64,$(ARCH)),)
	$(error ARCH should be x86 or x64)
endif

ifeq ($(ARCH),x86)
export GOARCH=386
else
export GOARCH=amd64
endif

BUILD=build_$(ARCH)
DIST=usudo_v$(VER)_$(ARCH).zip

7ZDIR?=C:/Program Files/7-Zip
7Z?=$(7ZDIR)/7z.exe

$(shell mkdir -p $(BUILD))

GO=go build -o $@ $^

USUDO=$(BUILD)/usudo.exe $(BUILD)/usudo-w.exe
HELPER=$(BUILD)/usudo-helper.exe
ALL=$(USUDO) $(HELPER)

.PHONY: all clean dist

all: $(ALL)

$(USUDO): $(BUILD)/%.exe: %.go run.go shell32.go
	$(GO)

$(HELPER): usudo-helper.go
	$(GO)

clean:
	rm -rf $(BUILD)

dist: $(ALL)
	rm -f $(BUILD)/$(DIST)
	cd $(BUILD) && "$(7Z)" a $(DIST) $(notdir $(ALL)) ../LICENSE ../README.rst
