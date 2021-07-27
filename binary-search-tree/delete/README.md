# 二分探索木のノード削除

二分探索木Tから節点zを削除するdeleteNodeは次のようなアルゴリズムとなる

```
deleteNode(T, z)
    if z.left == NIL && z.right == NIL  // zが子を持たない時
        // zの親から自身を削除する
        if z.parent == NIL  // zが根の時
            'root of T' = NIL  // 削除するだけ
        else if z == z.parent.left  // zが親の左の子の時
            z.parent.left = NIL  // 親の左の子を削除する
        else  // zが親の右の子の時
            z.parent.right = NIL  // 親の右の子を削除する
    else if z.left != NIL || z.right == NIL  // zが片方の子のみ持つ場合
        // その片方の子をxとする
        if z.left != NIL
            x = z.left
        else
            x = z.right
        // zの子xから自身を削除して自身の親を自身の子xの新たな親とし、zの親から自身を削除して自身の子xを自身の親の新たな子とする
        if z.parent == NIL  // zが根の時
            x.parent = NIL  // zの子xから自身を削除する
            'root of T' = x  // zの子xを新たな根とする
        else if z == z.parent.left  // zが親の左の子の時
            x.parent = z.parent  // zの子xから自身を削除して自身の親を自身の子xの新たな親とする
            z.parent.left = x  // zの親から自身を削除して自身の子xを自身の親の新たなことする
        else  // zが親の右の子の時
            x.parent = z.parent  // zの子xから自身を削除して自身の親を自身の子xの新たな親とする
            z.parent.right = x  // zの親から自身を削除して自身の子xを自身の親の新たなことする
    else  // zが子を両方持つ場合
        y = getSuccessor(z)  // zの次節点(inorderで巡回した時の次の節点)をyと置く. yは子を1つ以下しか持たない
        deleteNode(T, y)  // zではなくzの次節点yを消してしまう
        z.key = y.key  // yのデータをzにコピーする
```

以上のように実装をすれば、挿入・探索・削除が`O(logn)`で自走することができる。

## 実行例
### 入力例

```
18
insert 8
insert 2
insert 3
insert 7
insert 22
insert 1
find 1
find 2
find 3
find 4
find 5
find 6
find 7
find 8
print
delete 3
delete 7
print
```

### 出力例

```
yes
yes
no
no
no
yes
yes
1 2 3 7 8 22
8 2 1 3 7 22
1 2 8 22
8 2 1 22
```