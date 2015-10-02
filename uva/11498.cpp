#include <stdio.h>

int main ()
{
  int x, y, a, b, n;
  while (scanf("%d", &n), n) {
    scanf("%d %d", &x, &y);
    while(n-- > 0) {
      scanf("%d %d", &a, &b);
      if (a > x) {
        if (b > y) {
          printf("NE\n");
        } else if (b < y) {
          printf("SE\n");
        } else {
          printf("divisa\n");
        }
      } else if (a < x) {
        if (b > y) {
          printf("NO\n");
        } else if (b < y) {
          printf("SO\n");
        } else {
          printf("divisa\n");
        }
      } else {
        printf("divisa\n");
      }
    }
  }
  return 0;
}
