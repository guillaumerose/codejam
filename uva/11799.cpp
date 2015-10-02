#include <stdio.h>
#include <algorithm>

using namespace std;

int main ()
{
  int n, p, c, m;
  int i = 1;
  scanf("%d", &n);
  while (n-- > 0) {
    scanf("%d", &p);
    m = -1;
    while(p-- > 0) {
      scanf("%d", &c);
      m = max(m, c);
    }
    printf("Case %d: %d\n", i++, m);
  }
  return 0;
}
