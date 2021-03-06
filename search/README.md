# 探索

データの集合の中からもこうてきの要素を探し出す処理を探索という。

例えば、`{8, 13, 5, 7, 21, 11}`の中で7が何番目に現れるか、などの問題を解決する。

## 基本的な探索手法

### 線形探索

- 線形探索は、配列の先頭から各要素が目的の値と等しいかどうかを順番に調べる。
- 等しいものが見つかった時点でその位置を返して探索を終了する。
- 末尾まで調べて目的の値が存在しなかった場合は、そのことを示す特別な値を返す。
- アルゴリズムの効率は悪いが、データの並び方に関係なく適用可能。

### 二分探索

- コンピュータで扱うデータは、多くの場合ある項目によって整列され管理されている。
- このような場合より効率的な探索アルゴリズムの適用が可能。
- 二分探索はデータの大小関係を利用した高速な探索アルゴリズム

### ハッシュ法

- ハッシュ法ではハッシュ関数と呼ばれる関数値によって要素の格納場所を決定する
- ハッシュテーブルと呼ばれる表を用いるアルゴリズム
- 要素のキーを引数とした関数を呼び出すだけでその位置の特定ができるので、データの種類によっては高速