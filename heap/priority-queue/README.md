# 優先度付きキュー

優先度付きキューは各要素がキューを持ったデータの集合Sを保持するデータ構造で、次の操作を行う

- `insert(S, k)` : 集合Sに要素kを挿入する
- `extractMax(S)` : 最大のキーを持つSの要素をSから削除してその値を返す

キーが大きいものを優先するmax-優先度付きキューはmax-ヒープにより実装することが可能。

max-優先度付きキューSにkeyを追加する操作insertは次のようなアルゴリズム

```
insert(key)
    H++
    A[H] = -INFTY
    heapIncreaseKey(A, H, key)  // A[H]にkeyを設定する
```

heapIncreaseKeyは二分ヒープの要素iのキー値を増加させる処理で、次のようなアルゴリズムになる

```
heapIncreaseKey(A, i, key)
    if key < A[i]
        エラー : 新しいキーは現在のキーより小さい
    A[i] = key
    while i > 1 && A[parent(i)] < A[i]
        A[i]とA[parent(i)]を交換
        i = parent(i)
```

