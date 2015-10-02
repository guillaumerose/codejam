#include <stdio.h>
#include <algorithm>
#include <climits>

int main ()
{
  int n, b, h, w, price, beds;
  while (scanf("%d %d %d %d\n", &n, &b, &h, &w) != EOF) {
    int min = INT_MAX;
    for (int i = 0; i < h; ++i)
    {
      scanf("%d\n", &price);
      for (int j = 0; j < w; ++j)
      {
        scanf("%d", &beds);
        if (beds >= n) {
          min = std::min(min, n * price);
        }
      }
    }
    if (min > b) {
      printf("stay home\n");
    } else {
      printf("%d\n", min);
    }
  }
  return 0;
}
