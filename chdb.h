#ifndef CHDB_H
#define CHDB_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct local_result
{
    char * buf;
    size_t len;
    void * _vec; // std::vector<char> *, for freeing

    double elapsed;
} local_result;

struct local_result* query_stable(int arg, char ** argv);
void free_result(struct local_result *result);


static char**makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}

static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}

static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++) {
                free(a[i]);
	}
        free(a);
}







#ifdef __cplusplus
}
#endif

#endif
