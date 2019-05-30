# GenTest
良質なテストコードの自動生成による､ソフトウェアテスト環境の構築の実装｡

### Setup
```sh
$ git clone https://github.com/TakuKitamura/genTest.git
$ cd genTest
$ go run main.go sample_code/hello_world.gent 
Hello, GenTest!
$ go run main.go sample_code/grammar_example.gent
このようなことしかできません｡
できることその1: 演算
sqrt((2.0 + (3.0 - 4.0) / 5.0) + 0.1 * 2.0) is 1.4142135623730951
できることその2: 評価
1 < 2 && (3 != 3 || true) is true
できることその3: 文字列結合
できることその4: if文による分岐
2020 年は
うるう年です｡
演算
結果1: 0.6000000000000001
結果2: 0.0000000000000000
結果3: -0.6000000000000001
式評価
結果1: true
結果2: false
結果3: false
IF文
3
4
5
FOR文
i: 0 j: 0 k: 0
i: 0 j: 0 k: 1.0000000000000000
i: 0 j: 0 k: 2.0000000000000000
i: 0 j: 1.0000000000000000 k: 0
i: 0 j: 1.0000000000000000 k: 1.0000000000000000
i: 0 j: 1.0000000000000000 k: 2.0000000000000000
i: 0 j: 2.0000000000000000 k: 0
i: 0 j: 2.0000000000000000 k: 1.0000000000000000
i: 0 j: 2.0000000000000000 k: 2.0000000000000000
i: 1.0000000000000000 j: 0 k: 0
i: 1.0000000000000000 j: 0 k: 1.0000000000000000
i: 1.0000000000000000 j: 0 k: 2.0000000000000000
i: 1.0000000000000000 j: 1.0000000000000000 k: 0
i: 1.0000000000000000 j: 1.0000000000000000 k: 1.0000000000000000
i: 1.0000000000000000 j: 1.0000000000000000 k: 2.0000000000000000
i: 1.0000000000000000 j: 2.0000000000000000 k: 0
i: 1.0000000000000000 j: 2.0000000000000000 k: 1.0000000000000000
i: 1.0000000000000000 j: 2.0000000000000000 k: 2.0000000000000000
i: 2.0000000000000000 j: 0 k: 0
i: 2.0000000000000000 j: 0 k: 1.0000000000000000
i: 2.0000000000000000 j: 0 k: 2.0000000000000000
i: 2.0000000000000000 j: 1.0000000000000000 k: 0
i: 2.0000000000000000 j: 1.0000000000000000 k: 1.0000000000000000
i: 2.0000000000000000 j: 1.0000000000000000 k: 2.0000000000000000
i: 2.0000000000000000 j: 2.0000000000000000 k: 0
i: 2.0000000000000000 j: 2.0000000000000000 k: 1.0000000000000000
i: 2.0000000000000000 j: 2.0000000000000000 k: 2.0000000000000000
関数
Hello!
6.0000000000000000
π ≒ 3.1415926535897940
```

### Usage
```sh
$ go run main.go 
usage: gent file
exit status 1
```

### Author
- Taku Kitamura
