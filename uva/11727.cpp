#include <stdio.h>
#include <vector>
#include <algorithm>

int main ()
{
  int n;
  int i = 1;
  std::vector<int> v(3);
  scanf("%d", &n);
  while (n-- > 0) {
    scanf("%d %d %d", &v[0], &v[1], &v[2]);
    std::sort(v.begin(), v.end());
    printf("Case %d: %d\n", i++, v[1]);
  }
  return 0;
}
