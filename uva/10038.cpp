#include <cstdio>
#include <algorithm>
#include <vector>

using namespace std;

int verify(vector<int> differences, int n) {
  for (int i = 0; i < n - 1; i++)
  {
    if (differences[i] != i + 1) {
      return 1;
    }
  }
  return 0;
}

int main()
{
  int n, p, memory;
  while (scanf("%d", &n) != EOF) {
    vector<int> differences(n - 1);
    for (int i = 0; i < n; i++)
    {
      scanf("%d", &p);
      // printf("p = %d, memory = %d\n", p, memory);
      if (i > 0) {
        differences[i - 1] = abs(p - memory);
      }
      memory = p;
    }
    sort(differences.begin(), differences.end());
    if (verify(differences, n) != 0) {
      printf("Not jolly\n");
    } else {
      printf("Jolly\n");
    }
  }
  return 0;
}
