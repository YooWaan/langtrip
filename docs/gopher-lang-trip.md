name: top
class: center, middle, top-background

<h1 style="line-height:1.5em; font-size:6rem;">Gopher と行く<br/> Lang な旅</h1>

<br/>

.title-img[![gopher](images/gophercolor.png)]


---

# Pre test


```golang
func main() {
   println("hello")
}
# comment
```

<button onclick="hello();">Hello</button>

<p><i class="material-icons">account_circle</i> Hello</p>

<div class="center">
<button class="badge badge-pill pl-5 pr-5 pt-3 pb-3">Border </button>
</div>

---
template: inverse

## How does it work, then?

---

# Components


<div class="row mb-4">

  <div class="col-md-4">

    <div class="card">
      <div class="card-header">
        Featured
      </div>
      <div class="card-body">
        <h5 class="card-title">Hello</h5>
        <p class="card-text">
        Detail xxxxxxxx
        </p>
      </div>
    </div>

  </div>
</div>

<div class="row mb-4">

  <div class="col-md-4 mb-4">
	<p class="h1">
	  <span class="badge badge-pill badge-info">Info</span>
	  <span class="badge badge-pill badge-light">Light</span>
	  <span class="badge badge-pill badge-dark">Dark</span>
	</p>
  </div>
  
  <div class="col-md-4">

    <i class="material-icons">mood_bad</i>
    <i class="material-icons">mood</i>
    
    <i class="fa fa-link"></i> 
    <i class="fa fa-arrow-right"></i> 
  </div>
  <div class="col-md-4">

  </div>

</div>

---

## <span class="en">Introduction</span>

<div class="row mt-5 mb-4">
  <div class="col-md-3">
  <h2>Name:</h2>
  </div>
  <div class="col-md-9">
  <h2>xxxxxxxxx</h2>
  </div>
</div>

### Details

.display-2[
- <span><i class="material-icons align-middle">location_city</i> aaaa</span>
- <span><i class="material-icons align-middle">perm_identity</i> エンジニアリングマネージャー</span>
   - <span><i class="material-icons align-middle">assignment_ind</i> 開発部基盤チーム</span>
]

<div class="row mb-4">
  <div class="col-md-8">
	<p class="h1">
	  <span class="badge badge-pill badge-info">Info</span>
	  <span class="badge badge-pill badge-light">Light</span>
	  <span class="badge badge-pill badge-dark">Dark</span>
	</p>
  </div>
  <div class="col-md-4">
  <p class="bubble">aaaa</p>
  </div>
</div>

---
class: center, middle

.big[先に]

---
class: middle


.text-center[
## エンジニアを募集しています
]

<p class="strong">
  <span class="badge badge-pill badge-warning align-middle">🤓 Data Engineer</span>
  <span class="badge badge-pill badge-primary align-middle">⛑ Site Reliability Engineer</span>
  <span class="badge badge-pill badge-light">Application Engineer</span>
</p>

.text-center[
よろしくお願いします
]

---
class: center, middle

.big[
今日のお話
]

.title-img[
![gopher](images/gophercolor.png)
]

---
class: center, middle

.display-2[
- どう書いたらいいのか？
- どんな言語なのか？
]

<div class="row justify-content-center mb-4">
  <div class="col-md-5" >
  <p class="bubble display-2" style="width:420px;">不安を解決</p>
 </div>
</div>
<div class="row justify-content-center">
<img alt="gopher" src="images/gophercolor.png" width="10%" height="10%"/>
</div>

---
class: center, middle

## <span class="en">Beginer</span> 向けコンテンツになります


<p class="strong">
 <span class="badge badge-pill badge-primary"> ★ </span> <span class="text-primary">Junior</span>
</p>

.display-2[
<span class="badge badge-pill badge-light"> 😰 </span> Senior (中級者)

<span class="badge badge-pill badge-light"> 😱 </span> Expert (上級者)
]

---
class: center, middle

.strong[
なので、遠慮なく途中で

**質問** or **何か聞きたい** とか

大丈夫です
]

---

# 📜 概要

.display-4[
1. そもそもの話
    - Golang の基本的なことを見ていきます
2. BestPractice12の紹介とそのTips
    - Golang の基本的なことを見ていきます
3. 各言語 (Each Language) のTips
    - 各言語の特徴（**筆者の主観**）をGoで書くとどうなるか
	   - <span class="en">Ruby / Python / Java / Javascript / C#</span>
4. 実装手法 (Implementation) のTips
    - 今まであったトレンド（**筆者の主観**）でGoで書くとどうなるか
	   - <span class="en">Linq / ReactiveX</span>
]



---
class: center, middle


.big[
前置きは終わりで

始めていきます
]

---
class: center, middle

.big[
そもそもの話
]

---
## そもそもの話

<div class="container border rounded">
<div class="row fh ml-3 ml-3 mt-5" style="height:40vh;">
.display-4[
ものを認識 と 懸念事項 を考慮して 取っ掛かり の話
]

.display-3[
▶ Golangとは （認識）️ 

開発の要素 (懸念事項)

始める準備 （取っ掛かり）
]

</div>
</div>


---


## Golang とは

.history-text[
<table class="table">
  <tr>
  <td>
	<i class="icon-python text-success"></i><br/>
	<i class="icon-java-duke"></i>
  </td>
  <td class="align-middle"><i class="icon-ruby text-danger"></i></td>
  <td class="align-middle"><i class="icon-javascript-alt" style="color:gold;"></i></td>
  <td></td>
  <td class="align-middle"><i class="icon-csharp text-success"></i></td>
  <td></td>
  <td></td>
  <td class="align-middle"><i class="icon-go text-info"></i></td>

  <td class="align-middle"><img class="history-img"src="https://upload.wikimedia.org/wikipedia/commons/b/b5/Kotlin-logo.png"/></td>
  <td><img  src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9d/Swift_logo.svg/40px-Swift_logo.svg.png"/></td>
  <td></td>
  </tr>
<tr class="en">
<td>1991</td>
<td>1993</td>
<td>1994</td>
<td>1995</td>
<td>2002</td>
<td>2006</td>
<td>2007</td>
<td>2009</td>
<td>2011</td>
<td>2014</td>
<td></td>
</tr>
  <tr>
  <td></td>
  <td></td>
  <td><img class="history-img" src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a8/Sega-Saturn-Console-Set-Mk1.jpg/1920px-Sega-Saturn-Console-Set-Mk1.jpg"  /></td>
  <td><img class="history-img" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/PSX-Console-wController.png/280px-PSX-Console-wController.png"/></td>
  <td>3G</td>
  <td><img class="history-img" src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/d3/PS3Versions.png/250px-PS3Versions.png"/></td>
  <td>iPhone</td>
  <td></td>
  <td>4G</td>
  <td></td>
  <td></td>
  </tr>
</table>

]


<small>
[プログラム年表 wikipediaより](https://ja.wikipedia.org/wiki/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9E%E5%B9%B4%E8%A1%A8)
</small>
https://bridgera.com/5g-promises-new-horizons-for-iot/

http://wiki.c2.com/?AntiPatternsCatalog
https://dmitri.shuralyov.com/idiomatic-go


---

## Golang とは



* expressive
* concise
* clean
* efficient


<img src="images/what-is-golang.png" width="60%"/>

---
## そもそもの話

<div class="container border rounded">
<div class="row fh ml-3 ml-3 mt-5" style="height:40vh;">
.display-4[
ものを認識 と 懸念事項 を考慮して 取っ掛かり の話
]

.display-3[
✔️ Golangとは （認識）️ 

▶ 開発の要素 (懸念事項)

始める準備 （取っ掛かり）
]

</div>
</div>


---

## 開発の生産性

.display-4[
* ものを作る

* 設計・実装スキル

* TODO
]

---

.left-column[
  ### ものを作る
]
.right-column[

早くコーディング

- 理解しやすさ
- 冗長にならない書きやすさ
- IDE, Eco System


早く環境構築

- オーサリング
- モジュール管理

早くローンチ,デリバリー

- Build
- Test
- Deploy
]


---

.left-column[
  ### ものを作る
  ### 設計・実装スキル
]
.right-column[

xxx指向

- OOP,FP
- ドメイン駆動
- and more

Lang

- Syntax Sugar
- Idiom

<i class="material-icons">mood_bad</i> Senior, Expert からのプレッシャー

- 設計がぁー
- こっちの言語だとホゲホゲだけど、 Golangではフガフガがぁー

]

---

.left-column[
  ### ものを作る
  ### 設計・実装スキル
  ### TODO
]
.right-column[

TODO Things 1

- xxxx

TODO Things 2

- one
- two

]

---
## そもそもの話

<div class="container border rounded">
<div class="row fh ml-3 ml-3 mt-5" style="height:40vh;">
.display-4[
ものを認識 と 懸念事項 を考慮して 取っ掛かり の話
]

.display-3[

✔️ Golangとは （認識）️ 

✔ 開発の要素 (懸念事項)

▶ 始める準備 （取っ掛かり）
]

</div>
</div>


---

## 取っ掛かりとして

TODO

---
## コラム.1

<div class="center">
.strong[
<p class="bubble">
大事な事は目では見えないよ<br/>
心で見ないと
</p>
]
</div>

---


そもそも


- Go言語って

   - 歴史
   - 雑なお話

https://www.golang-book.com/books/intro

- 不安に要素

   - senior,expert からのプレッシャー
   - 


- 何から始めたら


- 

---
class: center, middle

.big[
.en[Go Bestpractice Tips]
]

---

# BestPractices の紹介と振り返り

[Twelve Go BestPractices](https://talks.golang.org/2013/bestpractices.slide#1) よりGolangの書き方のエッセンスを見ていきましょう


.left-split[

1. Avoid nesting by handling errors first
2. Avoid repetition when possible
3. Important code goes first
4. Document your code
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent
10. Avoid concurrency in your API
11. Use goroutines to manage state
12. Avoid goroutine leaks
]

.right-split[
    <img src="images/go-talks.png" width="50%">
]

<a class="badge badge-pill badge-info" href="https://talks.golang.org/"> Go Talks <i class="fa fa-arrow-right"></i></a>

---

## Twelve Best Practice ?


.left-split[
.text-center[
<img src="images/bestpractice-about.png" width="90%"/>  
Developer Advocate for the Go team at Google and for Google Cloud Platform
]
]
.left-split[
.text-center[
<img src="images/talk-bestpractice.png" width="90%" />

Go talks 2013
]
]


---

## Twelve BestPractice Tipsおさらい

```bash
 Beginner はここから
1. Avoid nesting by handling errors first
2. Avoid repetition when possible
3. Important code goes first
4. Document your code

 中級に向けて
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent

 中級以上のお話
10. Avoid concurrency in your API
11. Use goroutines to manage state
12. Avoid goroutine leaks
```


---

## Some code

```golang
type Gopher struct {
    Name     string
    AgeYears int
}
```

```golang
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
    if err == nil {
        size += 4
        var n int
        n, err = w.Write([]byte(g.Name))
        size += int64(n)
        if err == nil {
            err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
            if err == nil {
                size += 4
            }
            return
        }
        return
    }
    return
}
```


---

## 1. Avoid nesting by handling errors first

```golang
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
    if err != nil {
        return
    }
    size += 4
    n, err := w.Write([]byte(g.Name))
    size += int64(n)
    if err != nil {
        return
    }
    err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
    if err == nil {
        size += 4
    }
    return
}
```

---

## 2. Avoid repetition when possible

```golang
type binWriter struct {
    w    io.Writer
    size int64
    err  error
}

func (w *binWriter) Write(v interface{}) {
    if w.err != nil {
        return
    }
    if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
        w.size += int64(binary.Size(v))
    }
}
```

```golang
func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
    bw := &binWriter{w: w}
    bw.Write(int32(len(g.Name)))
    bw.Write([]byte(g.Name))
    bw.Write(int64(g.AgeYears))
    return bw.size, bw.err
}
```

---

## 3. Important code goes first

ライセンス情報、ビルドタグ(Build Tags)、package ドキュメント
import 文を関係するグループ毎に blank line で分けましょう

```golang
import (
    "fmt"
    "io"
    "log"

    "golang.org/x/net/websocket"
)
```

重要な type から始めて、最後の方に補助的な func, type を記述する

---
## 4. Document your code

パッケージ名の前に内容を書きましょう  
<span class="text-muted">Package name, with the associated documentation before. </span>

```golang
// Package playground registers an HTTP handler at "/compile" that
// proxies requests to the golang.org playground service.
package playground
```

外部公開するものは、ちゃんと書きましょう  
<span class="text-muted">Exported identifiers appear in godoc, they should be documented correctly.</span>

```golang
// Author represents the person who wrote and/or is presenting the document.
type Author struct {
    Elem []Elem
}

// TextElem returns the first text elements of the author details.
// This is used to display the author' name, job title, and company
// without the contact details.
func (p *Author) TextElem() (elems []Elem) {
```

<a class="badge badge-pill badge-info" href="https://blog.golang.org/godoc-documenting-go-code" target="_blank"> Godoc: documenting code <i class="fa fa-arrow-right"></i></a>

---

## 5. Shorter is better

長い名前はいいとは限らないよ  
<span class="text-muted">or at least longer is not always better.</span>

自明な短い名前を付けましょう  
<span class="text-muted">Try to find the shortest name that is self explanatory</span>

*  **MarshalWithIndentation** より **MarshalIndent** が好ましい 

package名は import を行ったで prefix として利用します  
<span class="text-muted">Don't forget that the package name will appear before the identifier you chose. </span>

```golang
pacakge json

var (
    JSONEncoder = someThing{} // bad
    Encoder = someThing{}  // good
)
```

```golang
import "json"

json.JSONEncoder.foo() // bad
json.Encoder.foo()  // good

```

---

## 6. Packages with multiple files

* 長い（大きい）ファイルは避ける
* code と testcode は分ける （test code はテストの時のみコンパイルされます）
* package のドキュメントは分ける

.text-center[
    <i class="icon-go text-primary" aria-hidden="true"></i> package が大きくなるようなら doc.go にドキュメントを書く事といいです
    <br/><br/>
    <img src="images/http-pkg.png" width="50%">
]


???

httpパッケージはそのような形をとっています

---

## Make your packages "go get"-able

<code class="remark-inline-code inline-b">"go get"</code> 可能な単位でpackageを作成しましょう

* コマンド等は再利用できないものもある
* 再利用可能なものはその単位


---

## 8. Ask for what you need

```golang
type Gopher struct {
    Name     string
    AgeYears int
}
```

```golang
// 1. method 定義ができます
func (g *Gopher) WriteToFile(f *os.File) (int64, error) {
// 2. 具体的な型はテストしにくいので、Interfaceを利用する
func (g *Gopher) WriteToReadWriter(rw io.ReadWriter) (int64, error) {
// 3. 関数で必要なことにを検討する io.ReaderWriter → io.Writer
func (g *Gopher) WriteToWriter(f io.Writer) (int64, error) {
```

× ファイルへ書き込みたい

　　↓↓↓

◯ 情報を書き込みたい

---

## 9. Keep independent packages independent

```golang
import (
    "golang.org/x/talks/2013/bestpractices/funcdraw/drawer"
    "golang.org/x/talks/2013/bestpractices/funcdraw/parser"
)
```

```golang
    // Parse the text into an executable function.
    f, err := parser.Parse(text)
    if err != nil {
        log.Fatalf("parse %q: %v", text, err)
    }

    // Create an image plotting the function.
    m := drawer.Draw(f, *width, *height, *xmin, *xmax)

    // Encode the image into the standard output.
    err = png.Encode(os.Stdout, m)
    if err != nil {
        log.Fatalf("encode image: %v", err)
    }
```

---

## その他のTips


.display-4[
10． Avoid concurrency in your API  
11． Use goroutines to manage state  
12． Avoid goroutine leaks

今回は割愛します
]


---

## 小まとめ


---
class: center, middle


.big[
各言語の .en[Tips]
]

---
## <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承

<div class="text-center mermaid">
graph LR;
    Child-- 継承 -->Parent;
</div>



```java
// Java
class Parent {
    public void say() { System.out.println("parent"); }
}

class Child extends Parent {
    public void say() { System.out.println("child"); }
}
```

```python
// python
class Parent:

    def say(self):
        print('parent')

class Child(Parent):

    def say(self):
        print('child')
```

---
## <i class="icon-go"></i> Class(クラス) Example

---
## <i class="icon-go"></i> Class(クラス) Summary


---
## <i class="icon-ruby text-danger"></i> Ruby block variable

```ruby
# code & result
[1, 2, 3].each do |i|
    put i*2
end

# result
1
4
9
```

.text-right[
<a class="badge badge-pill badge-info" href="http://docs.ruby-doc.com/docs/ProgrammingRuby/html/intro.html#S6" target="_blank">docs: Blocks and Iterators <i class="fa fa-arrow-right"></i></a>
]

```ruby
class List
  def initialize(arr)
    @arr = arr
  end
  def each
    i = 0
    while i < @arr.length
      yield @arr[i]
      i += 1
    end
  end
end
```

---
## <i class="icon-go"></i> Ruby block variable: Example

---
## <i class="icon-go"></i> Ruby block variable: Summary


---
## <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await


---
## <i class="icon-go"></i> async, await: Example

---
## <i class="icon-go"></i> async, await: Summary

---
## <i class="icon-java-duke"></i>  anotation, <i class="icon-python text-success"></i> decorator

---
## <i class="icon-go"></i>  anotation, decorator: Example

---
## <i class="icon-go"></i>  anotation, decorator: Sumamry

---
## <i class="icon-java-duke"></i>  interface, <i class="icon-python text-success"></i> duck typing

---
## <i class="icon-go"></i> interface, duck typing: Example

---
## <i class="icon-go"></i> interface, duck typing: Summary


---
class: center, middle

.big[実装Tips]


---

## Linq

データに集合対して

* 標準化した方式で賭合わせを可能にした技術
* 言語レベルで高階関数を利用して賭合わせできるもの

```csharp
  // Specify the data source.
  int[] scores = new int[] { 97, 92, 81, 60 };
  // Define the query expression.
  IEnumerable<int> scoreQuery =
      from score in scores
      where score > 80
      select score;
  // Execute the query.
  foreach (int i in scoreQuery) {
      Console.Write(i + " ");
  }
```


### References

- [Wikipedia EN](https://en.wikipedia.org/wiki/Language_Integrated_Query) /  [Wikipedia JA](https://ja.wikipedia.org/wiki/%E7%B5%B1%E5%90%88%E8%A8%80%E8%AA%9E%E3%82%AF%E3%82%A8%E3%83%AA)
- [C# Programming concepts](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/)
- go-linq [godoc](https://godoc.org/github.com/ahmetb/go-linq) / [github](https://github.com/ahmetb/go-linq)


---
## Linq　Example


---
## Linq　Summary


---
## ReactiveX

Observer, Iterator パターンと Functional Programming のアイデアを組み合わせたもの

---
## ReactiveX Exmaple

---
## ReactiveX Sumamry


### References

- [ReactiveX](http://reactivex.io/)
- [RxGo github](https://github.com/ReactiveX/RxGo)


---
class: center, middle

.big[まとめ]


---

１、概要

- 過去の発表であった BestPractice１２といろいろな言語や実装記述のTipsをベースに
  BestPracticeな部分をPickupしてく話を説明します
- 発表における Best Practice の説明します
   * Golang の良さ simple, readable, maintainable を軸としている旨の説明
   * Tipsとのエッセンスから恩恵やGolang（筆者）視点での考察
-

---

２、 BestPractice１２のTipsおさらい

* Beginner 向け (これを守れるだけでも十分すごいですな話)
1. Avoid nesting by handling errors first
2. Avoid repetition when possible
3. Important code goes first
4. Document your code

* 中級に向けて (知らずに中級者の仲間入りですな話)
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent

* 中級からの（省略）
10. Avoid concurrency in your API
11. Use goroutines to manage state
12. Avoid goroutine leaks

---
３、 言語のTips

- Java,Python ... : Class 継承ぽいTips
  * （エッセンス）悪手になりがちなのでやめよう
- Ruby block variable : ブロック変数ぽい Tips
  * （エッセンス）Slap になってればOKなのかも
- JS,C# await async
  * （エッセンス）goruotine を書こう
- Java,Python
  * annotation, decorator
  * Interface, duck typing
  * （エッセンス）誰得かはアナタ次第な話

---
４，実装Tips

- Linq
  * （エッセンス）Slice がサポートしていると嬉しいよね

- ReactiveX
  * （エッセンス）書けるとかっこいいかもしれないけど、宗教戦争みたいになったり
   そもそもがもう simple じゃない

---
５、 まとめ

 - Tips要素を引くと Golang の形になる旨を説明
 - go 2.x で generics が入いる事への思い
   - go 2 の survey summary にも触れつつ
 - Demo tips を絡めたコード実装例(Optional 発表に間に合えば)
