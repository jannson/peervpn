#CFLAGS+=-O2
CFLAGS+=-fno-inline -O0 -g -Wstrict-prototypes -Wall
#CFLAGS+=-I/usr/include/openssl -L/usr/lib/x86_64-linux-gnu
LIBS+=-lcrypto -lz

all: peervpn
peervpn: peervpn.o
	$(CC) $(LDFLAGS) peervpn.o $(LIBS) -o $@
peervpn.o: peervpn.c

clean:
	rm -f peervpn peervpn.o
