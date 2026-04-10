#include <stdio.h>

int probe(int v) {
    if (v > 5) return 1;
    return 0;
}

int main() {
    printf("sys_probe active\n");
    return 0;
}
