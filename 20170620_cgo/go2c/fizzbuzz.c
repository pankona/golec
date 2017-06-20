
#include <stdio.h>
#include <string.h>

int
fizzbuzz(int n, char* out, int out_len) {
    int ret;
    if (out_len < strlen("FizzBuzz")) {
        // buffer is too short
        return -1;
    }

    if (n%5 == 0 && n%3 == 0) {
        strcpy(out, "FizzBuzz");
        ret = strlen("FizzBuzz");
    } else if (n%5 == 0) {
        strcpy(out, "Buzz");
        ret = strlen("Buzz");
    } else if (n%3 == 0) {
        strcpy(out, "Fizz");
        ret = strlen("Fizz");
    } else {
        ret = snprintf(out, 24, "%d", n);
    }
    return ret;
}
