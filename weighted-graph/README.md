# 重み付きグラフ

グラフの辺に重み(値)が設定された重み付きグラフに関して見ていく

## 最小全域木

木とは閉路を持たないグラフであった。グラフG=(V, E)の全域木G=(V', E')とは、グラフの全ての頂点Vを含む部分グラフであり(V=V')、木であるかぎりできるだけ多くの辺を持つものを指す。

グラフの全域木は深さ優先探索や幅優先探索により求められるが、1つとは限らない。

特に最小全域木とは、グラフの全域木の中で辺の重みの総和が最小のものを指す。

## 最短経路

最短経路問題とは、重み付きグラフにおいて、与えられた頂点の組を接続する経路の中で、辺の重みの総和が最小であるパスを求める問題。

頂点sからGの全ての頂点に対してパスがある時、sを根としてsからGの全ての頂点への最短経路を包含するGの全域木が存在し、そのことを最短経路木と呼ぶ。
