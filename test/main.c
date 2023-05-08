// run `make lib` to drop lib and header files in .build directory
#include <stdio.h>
#include <libqr.h>
int main() {
    printf("Testing qr C library!\n");
    printf("%s", qrCodeSmall("chey wuz here", 0));
    printf("%s", qrCodeBig("chey wuz here", 0));

    Png png = qrCodePng("chey wuz here", 0);
    printf("PNG Size: %d\n", png.n);

    int expected_size = 649;
    /*
    char * fname = "chey.png";
    FILE* fp = fopen(fname, "w");
    fwrite(png.b, 1, png.n, fp);
    fclose(fp);
    printf("PNG written to: %s\n", fname);
    */
    if (png.n != expected_size) {
      fprintf(stderr, "lib produced incorrect size png of %d", expected_size);
      return 1;
    }
    return 0;
}
