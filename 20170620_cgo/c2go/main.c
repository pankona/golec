#include <stdio.h>
#include <fizzbuzz.h>

int
main(int argc, char** argv) {
    int i;
    GoInt n = 17;
    char *str;

    for(i = 0; i < n; i++) {
      str = fizzbuzz(i);
      printf("%s\n", str);
    }
}
