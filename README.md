# 素因数分解

## prime_dev.rb

第一引数を素因数分解します。
入力値を2から順に割っていき、余りが0になるものを探します。

http://idm.s9.xrea.com/factorization/naive.html
の改良前のものを参考にしています。


## prime_dev_6.rb

第一引数を素因数分解します。

prime_dev.rbでは2から順に割っていくため計算量が多いです。
以下の改良を施しています。

* 2, 3の倍数を除くと6m-1と6m+1のみ考えればよい。
* sqrt(n)以上に素因数が存在しない。

http://idm.s9.xrea.com/factorization/naive.html
を参考にしています。

## primeDev6.go

prime_dev_6.rbのGo言語版です。
引数ではなく、直接ソース中に対象の合成数を記載します。

