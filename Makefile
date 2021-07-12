BINDIR	= $(CURDIR)/bin

build:
	cd $(CURDIR)/src/check_consul_node && go get check_consul_node/src/check_consul_node && go build -o $(CURDIR)/bin/check_consul_node

all: depend build strip install

depend:
	# go mod will handle dependencies

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/check_consul_node

ifneq (, $(shell which upx 2>/dev/null))
	upx -9 $(BINDIR)/check_consul_node
endif

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/check_consul_node $(DESTDIR)/usr/lib64/nagios/plugins

clean:
	/bin/rm -f bin/check_consul_node

distclean: clean
	rm -rf src/github.com/
	rm -rf pkg/

uninstall:
	/bin/rm -f $(DESTDIR)/usr/lib64/nagios/plugins/check_consul_node

