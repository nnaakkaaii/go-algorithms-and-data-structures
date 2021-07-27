# 二分探索木への挿入

二分探索木Tに新たに値vを挿入するには、次のようなinsertを実行する。

insertでは、キーがv, 左の子がNIL, 右の子がNILであるような点zを受け取ってTの正しい位置に挿入する。

このアルゴリズムでは挿入操作の計算量は、木の高さをhとするとO(h)となる。

つまり、節点の個数nに対して、入力に偏りがなければ`O(logn)`の計算量。

二分探索木のバランスが良いほど計算量は理論値に近づく。

```
insert(T, z):
    y = NIL  // xの親
    x = 'Tの根'
    
    // 葉に到達するまで正しい大小関係を保ちながら深くへ推移
    while x != NIL
        y = x  // 親を設定
        if z.key < x.key
            x = x.left  // 左の子へ移動
        else
            x = x.right  // 右の子へ移動
    z.p = y  // 最後の葉がzの親に当たる
    
    if y == NIL  // Tが空の場合 (木を新規で作成)
        'Tの根' = z
    else if z.key < y.key
        y.left = z  // zをyの左の子にする
    else
        y.right = z  // zをyの右の子にする
```