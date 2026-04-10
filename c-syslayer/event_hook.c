#include <stdio.h>

void hook(const char *event) {
    printf("[hook] %s\n", event);
}
