#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    if (argc > 1) {
        int val = atoi(argv[1]);

        if (val > 5) {
            printf("1\n");
        } else {
            printf("0\n");
        }
    }

    return 0;
}
