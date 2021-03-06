# 二分探索木

実行時に必要なメモリ領域を確保し動的に集合を取り扱うデータ構造として、データの追加・削除・探索を行うことができる連結リストがあった。

一方、連結リストの探索には`O(n)`の計算量が必要がであった。

そこで、動的な木構造を利用することでデータの追加・削除・探索をより効率的に行うことができる。

## 二分探索木とは

探索木は、挿入・検索・削除などの操作が行えるデータ構造で、辞書や優先度付きキューとしての利用が可能。

探索木の中でも最も基本的なのが二分探索木である。

二分探索木は各節点にキーを持ち、次の二分探索木条件を常に満たすように構築される。

xを二分探索木に属するある節点とする。
- yをxの左部分木に属する節点とすると、yのキーはxのキー以下となる
- zをxの右部分木に属する節点とすると、zのキーはxのキー以上となる

この二分木に対して中間順巡回を行うと、昇順に並べられたキーの列を得ることができる。

データの挿入・削除が行われても全ての節点でこの条件を満たすべきで、リストと同様に節点をポインタで連結することで木を表し、各節点にはキーに加えて親・左の子・右の子へのポインタを持たせる。

## 標準ライブラリによる集合の管理

listやvectorなどのシーケンスコンテナでは、追加される要素はその値に関係なく特定の位置に配置され、その位置は挿入した時間と場所に依存する

set, map, multiset, multimapなどの連想コンテナでは、追加される要素の位置は何らかのソート基準に従って決められる

連想コンテナでは要素は自動的にソートされて集合が管理され、常に二分探索が行えることが特徴。

### set, map

| 関数名 | 機能 | 計算量 |
| ---- | ---- | ---- |
| `size()` | セットの中の要素数を返す | `O(1)` |
| `clear()` | セットの中を空にする | `O(n)` |
| `begin()` | セットの先頭を指すイテレータを返す | `O(1)` |
| `end()` | セットの末尾を指すイテレータを返す | `O(1)` |
| `insert(key)` | セットに要素keyを挿入する | `O(logn)` |
| `erase(key` | keyを持つ要素を削除する | `O(logn)` |
| `find(key)` | keyに一致する要素を探索し、その要素へのイテレータを返す。要素がなければ末尾`end()`を返す | `O(logn)` |
