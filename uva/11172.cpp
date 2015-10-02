#include <stdio.h>

int main ()
{
  int i, j, n;
  scanf("%d", &n);
  while (n-- > 0) {
    scanf("%d %d", &i, &j);
    if (i > j) {
      printf(">\n");
    } else if (i < j) {
      printf("<\n");
    } else {
      printf("=\n");
    }
  }
  return 0;
}
