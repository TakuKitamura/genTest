# GenTest
良質なテストコードの自動生成による､ソフトウェアテスト環境の構築の実装｡

### Setup
```sh
$ git clone https://github.com/TakuKitamura/genTest.git
$ cd genTest
$ go run main.go sample_code/hello_world.gent 
Hello, GenTest!
$ go run main.go sample_code/grammar_example.gent
("このようなことしかできません｡")
("できることその1: 演算")
("sqrt((2.0 + (3.0 - 4.0) / 5.0) + 0.1 * 2.0) is ", 1.4142135623730951)
("できることその2: 評価")
("1 < 2 && (3 != 3 || true) is ", true)
("できることその3: 文字列結合")
```

### Usage
```sh
$ go run main.go 
usage: gent file
exit status 1
```

### Author
- Taku Kitamura
