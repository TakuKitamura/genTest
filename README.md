# GenTest
良質なテストコードの自動生成による､ソフトウェアテスト環境の構築の実装｡

### Setup
```sh
$ cat sample_code/example1.gent && echo -e "\n\n---\n" && go run main.go sample_code/example1.gent 
print("演算")
resultOne = 1 + 2 - 3*4/5
print("結果1:", resultOne)
resultTwo = ((((1 + 2) - 3) * 4) / 5)
print("結果2:", resultTwo)
resultThree = resultTwo - resultOne
print("結果3:", resultThree)

---

演算
結果1: 0.6000000000000001
結果2: 0.0000000000000000
結果3: -0.6000000000000001
$ cat sample_code/example2.gent && echo -e "\n\n---\n" && go run main.go sample_code/example2.gent 
print("式評価")
boolOne = 1 < 2
print("結果1:", boolOne)
boolTwo = 1 > 2
print("結果2:", boolTwo)
boolThree = boolOne && boolTwo
print("結果3:", boolThree)

---

式評価
結果1: true
結果2: false
結果3: false
$ cat sample_code/example3.gent && echo -e "\n\n---\n" && go run main.go sample_code/example3.gent 
print("IF文")
if 1 > 2:
    print(1)
    if 3 == 4:
        print(2)
if 1 < 2:
    print(3)
    if 3 != 4:
        print(4)
print(5)

---

IF文
3
4
5
$ cat sample_code/example4.gent && echo -e "\n\n---\n" && go run main.go sample_code/example4.gent 
print("FOR文")
for i = 0; i <= 2; i = i + 1:
    for j = 0; j <= 2; j = j + 1:
        for k = 0; k <= 2; k = k + 1:
            print("i:", i ,"j:", j, "k:", k)

---

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
$ cat sample_code/example5.gent && echo -e "\n\n---\n" && go run main.go sample_code/example5.gent 
print("関数")
func addThree(a, b, c):
    return a + b + c
func hello():
    print("Hello!")
hello()
result = addThree(1, 2, 3)
print(result)

---

関数
Hello!
6.0000000000000000
$ cat sample_code/example6.gent && echo -e "\n\n---\n" && go run main.go sample_code/example6.gent 
func sqrt(n): // ニュートン法による､平方根の計算
    if n < 0:
        r = -1
    if n == 0:
        r = 0
    if n > 0:
        r = 1
        for j = 0; j < 10; j = j + 1:
            r = ((n / r) + r) / 2
    return r

rootTwo = sqrt(2)
a = 1
b = 1 / rootTwo
t = 1 / 4
x = 1
pi = 0

for i = 1; i <= 3; i = i + 1:
    an = (a + b) / 2
    b = sqrt(a * b)
    t = t - (x * (an - a) * (an - a))
    x = x * 2
    a = an
    pi = (a + b) * (a + b) / (4 * t)

print("π ≒", pi) // ガウス＝ルジャンドルのアルゴリズムによる計算

---

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
