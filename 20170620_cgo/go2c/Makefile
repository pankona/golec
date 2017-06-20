

all: libfizzbuzz.a
	go build

libfizzbuzz.a:
	gcc -c fizzbuzz.c
	ar crv libfizzbuzz.a ./fizzbuzz.o

clean:
	rm -f libfizzbuzz.a
	rm -f fizzbuzz.o
	rm -f go2c
