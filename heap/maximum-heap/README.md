# 最大・最小ヒープ

`A[i]`をmaxヒープ条件を満たすまで木の葉に向かって下降させていく`maxHeapify(A, i)`関数を実装する

例えば、次のようになっている時、`maxHeapify(A, 1)`を実行することを考える。

```
      5
  86     17
12  25 32  11
```

まずは5の子を比べてキーが大きい方を選び、それが現在のキーよりも大きければ交換していく。

```
     86
  5      17
12 25  32  11
```

```
     86
  25      17
12  5  32  11
```

このアルゴリズムにおけるmaxHeapifyの計算量は完全二分木の高さに比例するので`O(logH)`となる
