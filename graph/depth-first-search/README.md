# 深さ優先探索

可能な限り隣接する頂点を訪問するという戦略に基づく探索アルゴリズム

未探索の接続辺が残されている頂点の中で最後に発見した頂点vの接続辺を再帰的に探索する

探索の際には元の始点から到達可能な全ての頂点を発見するまで続き、未発見の頂点が残っていればその中の番号が一番小さい1つを新たな始点として探索を始める

## 手順

DFSでは、「まだ探索中の頂点」を一時的に保持しておくため、スタックを用いて管理する

1. 一番最初に訪問する頂点をスタックに入れておく
2. スタックの頂点が積まれている限り、次の処理を繰り返す
    - スタックのトップにある頂点uを訪問する
    - 現在訪問中の頂点uから次の頂点vへ移動する時、vをスタックに積む。ただし訪問中の頂点uに未訪問の隣接する頂点がない場合、uをスタックから削除する
    
最終的に全ての頂点に訪問が完了したら終了

次のように変数を準備しよう

| 変数 | 意味 |
| ---- | ---- |
| `color[n]` | 頂点iの訪問状態をWHITE, GRAY, BLACKのいずれかで表す |
| `W[n][n]`  | 頂点iから頂点jに辺がある場合`M[i][j]`がtrueとなるような隣接行列 |
| `Stack S` | 訪問途中の頂点を退避しておくスタック |

## 考察

- 隣接行列を用いた深さ優先探索は、各頂点について全ての頂点に隣接しているかどうかを調べるので`O(|V|^2)`のアルゴリズムとなり大きなグラフに対しては適当ではない
- 隣接リストを用いることでより高速に実装が可能
