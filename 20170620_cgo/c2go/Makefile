
all: fizzbuzz.a
	gcc -o main main.c -I. ./fizzbuzz.a -lpthread

fizzbuzz.a:
	go build -buildmode=c-archive fizzbuzz.go

clean:
	rm -f ./main
	rm -f ./fizzbuzz.a
	rm -f ./fizzbuzz.h
