# RadixTreeGo
Radix TreeによるIPアドレスの解決をシミュレートする

### 動作
プログラム実行時に`route-02.txt`からネットワークアドレスを読み取り，Radix Treeを構成していく．<br>
その後，任意のアドレスで検索を行うことで，そのアドレスが属するネットワークアドレスの`route-02.txt`におけるインデックスを返す．

### 例
41.74.1.1     -> 56183 (41.74/21)<br>
66.31.10.3    -> 134544(66.30/15)<br>
133.5.1.1     -> 370522(133.5/16)<br>
209.143.75.1  -> 729716(209.143.64/18)<br>
221.121.128.1 -> 775303(221.121.128/19)<br>
