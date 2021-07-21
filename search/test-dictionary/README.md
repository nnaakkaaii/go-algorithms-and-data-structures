# ハッシュ

- findもinsertも`O(1)`で行うことができる
- リストに文字列を入力する際、文字列をハッシュ関数に代入して得たハッシュ値の指す位置に挿入する
    - 既にデータがある場合、`i`をインクリメントしていく

## 実行例
### input

```
6
insert AAA
insert AAC
find AAA
find CCC
insert CCC
find CCC
```
