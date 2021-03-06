# 最長共通部分列

`X = {a, b, c, b, d, a, b}`と`Y = {b, d, c, a, b, a}`の最長共通部分列は`{b, c, b, a}`となる

このように連続していなくても良い部分列の最大の長さを求める

## 考え方

まずは`X[m]`と`Y[n]`のLCSを求める問題を次のような部分問題に分割して考える

- `x[m] == y[n]`の場合は、`X[m-1]`と`Y[n-1]`のLCSに`x[m]`を連結したものが`X[m]`と`Y[n]`のLCSとなる
- `x[m] != y[n]`の場合は、`X[m-1]`と`Y[n]`のLCSまた`X[m]`と`Y[n-1]`のLCSの長い方が`X[m]`と`Y[n]`のLCSとなる

## 実行例
### 入力例

最初の行にデータセットの組の数、それ以降の行に2 x n行でn組のデータセットを記述する

```
3
abcbdab
bdcaba
abc
abc
abc
bc
```

### 出力例
```
4
3
2
```
