name: top
class: center, middle, top-background

<h1 style="line-height:1.5em; font-size:6rem;">.en[Gopher] と行く .en[Lang] な旅</h1>

<br/>

.title-img[![gopher](images/gophercolor.png)]

<p class="h2 en">Go Conference 2019 Spring</p> <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">
<img alt="クリエイティブ・コモンズ・ライセンス" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" />
</a>
<br />

<small>これは <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">クリエイティブ・コモンズ 表示 4.0 国際 ライセンス</a>の下に提供されています。</small>

---

## <span class="en">Introduction</span>

.left-split[
#### <span class="en">Profile</span>

```bash
 name       Ryo Terunuma
 handlename YooWaan
                 `-- Game / Twitter / Qiita
 company    BrainPad
                 `-- p 大文字
 role       Engineering Manager
```

#### <span class="en">Details</span>

```bash
$ Infra | Backend | Front
$ Member | TeamLead | Manager
```

#### <span class="en">Develop   </span>


```bash
$ Platform System
```

]

.right-split[

<a href="https://www.brainpad.co.jp/" target="_blank">
<img src="images/brainpad.png" width="90%"/>
</a>
]

---
class: center, middle

.big[
先に  
いつもの  
お約束
]

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

必ず採用担当までは,お繋ぎします
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
class: middle

.display-2[
.text-center[
- どう書いたらいいのか？
- どんな言語なのか？
]
]

<div class="row justify-content-center mb-4">
  <div class="col-md-5 text-center" >
  <p class="bubble display-2" style="width:420px;">不安を解決</p>
 </div>
</div>
<div class="row justify-content-center">
<img alt="gopher" src="images/gophercolor.png" width="10%" height="10%"/>
</div>

---
class: center, middle

## <span class="en">Beginner</span> 向けコンテンツになります


<p class="strong">
 <span class="badge badge-pill badge-primary"> ★ </span> <span class="text-primary">Junior / Beginner</span>
</p>

.display-2[
<span class="badge badge-pill badge-light"> 😰 </span> Senior (中級者)

<span class="badge badge-pill badge-light"> 😱 </span> Expert (上級者)
]

---
class: center, middle, eye-bg

.strong[
なので、遠慮なく途中で

**質問** or **何か聞きたい** とか

大丈夫です
]

---
class: center, middle

# Junior / Beginner

.display-4[
<br>

完全な初心者を想定していません

なんらかのプログラム(アプリケーション)を

書いたことがある事を

想定しています
]

---
class: center, middle, eye-bg

.strong[
旅

設定です
]

(謎)

---

# 📜 概要

.display-4[
1. 旅立ち 編
   - Gopher <i class="icon-go"></i> と準備と旅立ち
2. Golang な旅 編
   - Twelve Go Best Practices の紹介
3. Lang な旅 編
   - いろんな言語で楽しむ <span class="bg-warning small">&dagger;</span> <span class="en">(Ruby / Python / Java / Javascript / C#)</span>
4. 過去の Trend な旅 編
   - 今まであったトレンド名所巡り <span class="bg-warning small">&dagger;</span> <span class="en">(Linq / ReactiveX)</span>
5. We are one peace （まとめ <span class="small">旅の終わり</span>）
]

<span class="bg-warning small">&dagger;</span>（筆者の主観）

---
class: center, middle

.big[
旅立ち編
]

---
class: center, eye-bg

.big[
まず

Gopher に会う
]


---

## <i class="icon-go"></i> Golang とは


* expressive ... 表現が豊かで
* concise ... 簡潔で
* clean ... キレイで
* efficient ... 無駄がない(効率がいい)

<br/>

.text-center[
<img src="images/what-is-golang.png" width="60%"/>
]
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
  <td class="align-middle"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9d/Swift_logo.svg/40px-Swift_logo.svg.png"/></td>
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

参考リンク

<small>
[プログラム年表 wikipediaより](https://ja.wikipedia.org/wiki/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9E%E5%B9%B4%E8%A1%A8)
,[1G - 5G](https://bridgera.com/5g-promises-new-horizons-for-iot/)
</small>


???
https://dmitri.shuralyov.com/idiomatic-go


---
class: eye-bg

## 旅立ちの不安

.display-2[

.text-center[
Golang ? Gopher ?

どんな言語なのか？

どう書いたらいいのか？
]

]

<div class="row justify-content-center text-center">
  <div class="col-md-9 offset-md-3 text-center" >
  <p class="bubble display-2" style="width:64%;">不安な部分を解決</p>
 </div>
</div>

---
class: center, middle


<div class="h3 border-success rounded bg-info text-light pt-3 pb-3">
<b>不安のモトは何か？</b>
</div>

---
## モトを探る

.display-3[
ソフトウェアエンジニアがやりたいこと

- ものを作りたい
- 技法を適切に利用したい
- モチベ・哲学・思想を反映したい
]

---
## モトにある不安要素

.display-4[
- ものを作りたい
> → 実装できるか？
- 技法を適切に利用したい
> → 技法の適用できるか？  
> → Senior、Expert の威光（マウンティング）
- モチベ・哲学・思想を反映したい
> → いままでの経験はどうなる？  
> → 新しい慣習ってなんだろ？
]

---
## 旅を通じて感じていく事


.h1[
Golang <i class="icon-go"></i> の .bg-info[Simplicity] (単純・簡単)
]


.h2[
<br/>

* ものを作るイメージを広げる　 → Twelve Go Best Practices
* 技法の使い所を少し広める
* モチベ・哲学・思想の整理　　 → 各言語 / Trend
]

.text-right[
<a class="badge badge-pill badge-info" href="https://talks.golang.org/2015/simplicity-is-complicated.slide#1" target="_blank">
Simplicity is Complicated <i class="inline-b material-icons" style="display: inline-flex; vertical-align: middle;">exit_to_app</i>
</a>

<small>by Rob Pike</small>
]


???

詳細はPike さんのものを参照ください

---
class: center, middle

.big[
.en[Golang] な旅

.en[Twelve Go Bestpractice]
]

.title-img[![gopher](images/gophercolor.png)]

---

## Twelve Best Practice ?


.left-split[
.text-center[
<img src="images/bestpractice-about.png" width="90%"/>  
Developer Advocate for the Go team at Google and for Google Cloud Platform
]
]
.right-split[
.text-center[
<img src="images/talk-bestpractice.png" width="90%" />

Go talks 2013

]
.text-right[
<a class="badge badge-pill badge-info" href="https://talks.golang.org/"> Go Talks <i class="fa fa-arrow-right"></i></a>
]
]


---

# BestPractices の紹介


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

<a class="badge badge-pill badge-info" href="https://talks.golang.org/2013/bestpractices.slide#1" target="_blank">
Twelve Go Best Practices <i class="inline-b material-icons" style="display: inline-flex; vertical-align: middle;">exit_to_app</i>
</a>


---

## Twelve Go Best Practices

Golangの書き方のエッセンスを見ていきましょう


```bash
 Beginner はここから
1. Avoid nesting by handling errors first
2. Avoid repetition when possible
3. Important code goes first
4. Document your code
```

```bash
 中級に向けて
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent
```

今日はここまで

```bash
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

```go
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
    if err == nil {
        size += 4
        var n int
        // bad
        n, err = w.Write([]byte(g.Name))
        size += int64(n)
        if err == nil {
            // bad
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

ネストを回避して、１インデントでエラーハンドルをする

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

バイナリのWriterを作ってそれを再利用してコードを書く

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

```go
type FistData struct { /* some thing */}

func Func1() {}
func Func2() {}

type subData struct {}

```

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

package名は import を行った側で prefix として利用します  
<span class="text-muted">Don't forget that the package name will appear before the identifier you chose. </span>

```golang
pacakge json

var (
    JSONEncoder = someThing{} // bad
    Encoder = someThing{}     // good
)
```

```golang
import "json"

json.JSONEncoder.foo() // bad
json.Encoder.foo()     // good

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

.left-split[

```bash
  cmd/
    xxxget/
    xxxmount/
    xxxput/
```
]

.right-split[
```bash
  auth/
    auth.go
  blobref/
    blobref.go
    blobref_test.go
    chanpeek.go
    fetcher.go
```
]

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

## <span class="bg-info">小まとめ</span>


.left-split[

Gopher と仲良くなる為に

1. ネストエラーハンドルをしない
2. 繰り返しを避ける
3. 重要なコードから書いていく
4. ドキュメントを書く
5. 短い名前は良い
6. ファイルは小さく
7. go get 可能にする
8. 何が必要なのかを重要視する
9. パッケージの独立性を保つ
]

.right-split[


_.en[Golang] <i class="icon-go"></i> 初めての方は こちらから_


最初に見ると

* [Go言語の初心者が見ると幸せになれる場所](https://qiita.com/tenntenn/items/0e33a4959250d1a55045)

その他気をつける事

* [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
* [日本語翻訳](https://qiita.com/knsh14/items/8b73b31822c109d4c497)



上手にかける為に

* [Effective Go](https://golang.org/doc/effective_go.html)


コミュニティへ参加したい人は

* [Community Code of Conduct](https://golang.org/conduct)

]

---
class: center, middle

.display-1[
.en[Lang] な旅
]
.display-3[
いろんな言語で楽しむ
]

.title-img[![gopher](images/gophercolor.png)]


---
class: eye-bg

## いろんな言語の旅

.h1[.lh16[
* <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> (<i class="icon-ruby text-danger"></i> <i class="icon-csharp text-success"></i>)Class(クラス) の継承
* <i class="icon-ruby text-danger"></i> Ruby block variable
* <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await
* <i class="icon-java-duke"></i>  annotation, <i class="icon-python text-success"></i> decorator
* <i class="icon-java-duke"></i> interface, <i class="icon-python text-success"></i> duck typing
]]

---

## いろんな言語の旅

.h1[.lh16[
* <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承
]]

---
## <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承

<div class="text-center mermaid">
graph LR;
    Child-- 継承 -->Parent;

</div>


.left-split[

```java
// Java
class Parent {
    public void say() { System.out.println("parent"); }
}

class Child extends Parent {
    public void say() { System.out.println("child"); }
}
```

<button class="btn btn-raised" onclick="player.open('/src/clz.java', '/src/clz.py');">Eval</button>

]

.right-split[
```python
// python
class Parent:

    def say(self):
        print('parent')

class Child(Parent):

    def say(self):
        print('child')
```
]


---
## <i class="icon-go"></i> Class(クラス) Example

```golang
package main

type Parent struct {}
func (p Parent) say() string {return "parent" }
func (p Parent) Print() { println(p.say())}
func (p Parent) Hello() { println(p.say())}

type Child struct { Parent }
func (c Child) say() string {return "child" }
// Print() は Parent の処理を実行する
func (c Child) Hello() { println(c.say()) }

```

<button class="btn btn-raised" onclick="player.open('/src/clz.go');">Eval</button>


child の Hello() は "child" と出力するが、Print() は "parent" と出力します

---
## <i class="icon-go"></i> Class(クラス) Summary


[Effectiive Go#Embedding](https://golang.org/doc/effective_go.html#embedding)

Go does not provide the typical, type-driven notion of subclassing, ...

Golang サブクラスを提供していないので、interface を定義してそれを利用することを紹介してます


```golang
type Say interface {
	Print()
}

type Person struct {
	Say
}

type sayParent struct {}
func (p *sayParent) Print() { println("parent") }

type sayChild struct {}
func (p *sayChild) Print() { println("child") }

```

<button class="btn btn-raised" onclick="player.open('/src/clz_summary.go');">Eval</button>


---

## いろんな言語の旅

.h3[.lh16[
.text-muted[
* <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承
]]]

.h1[
* <i class="icon-ruby text-danger"></i> Ruby block variable
]


---
## <i class="icon-ruby text-danger"></i> Ruby block variable

```ruby
[1, 2, 3].each do |i|
    put i*2
end
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

<button class="btn btn-raised" onclick="player.open('/src/blockvar.rb');">Eval</button>

---
## <i class="icon-go"></i> Ruby block variable: Example

```golang
type Person struct {
    Name string
}

type People []Person

func (pp People) Each(hfn func(i int, p Person)) {
  for i, p := range pp {
     hfn(i, p)
  }
}

func main() {
   people := People{ Person{"tom"}, Person{"bom"} }

   names := []string{}
   people.Each(func(i int, p Person) {
       names = append(names, p.Name)
   })
   println(strings.Join(names, ","))
}
```

<button class="btn btn-raised" onclick="player.open('/src/blockvar.go');">Eval</button>

---
## <i class="icon-go"></i> Ruby block variable: Summary


<span class="bg-info h1">Golang</span>

* 下のように記述可能でブロック変数というような概念は不要

```go
for _, num := range []int{1, 2, 3}{
   // Do something
}
```

---

## いろんな言語の旅

.h3[.lh16[

.text-muted[
* <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承
* <i class="icon-ruby text-danger"></i> Ruby block variable
]]]

.h1[
* <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await
]

---
## <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await


<i class="icon-javascript-alt display-4" style="color:gold;"></i> の async/await

- async : 非同期処理関数を定義する <a class="badge badge-pill badge-info" href="https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Statements/async_function" target="_blank">async <i class="fa fa-external-link-alt"></i></a>
- await : Promiseのresolve, reject の結果を待つ  
awaitはasyncの関数内でしか使えない <a class="badge badge-pill badge-info" href="https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/await" target="_blank">async <i class="fa fa-external-link-alt"></i></a>


<i class="icon-csharp text-success display-4"></i> の async/await

- async <a class="badge badge-pill badge-info" href="https://docs.microsoft.com/ja-jp/dotnet/csharp/language-reference/keywords/async" target="_blank"> async <i class="fa fa-external-link-alt"></i></a> Task, Task<TResult> を返す非同期メソッド
- await <a class="badge badge-pill badge-info" href="https://docs.microsoft.com/ja-jp/dotnet/csharp/language-reference/keywords/await" target="_blank"> await  <i class="fa fa-external-link-alt"></i></a> 非同期処理(Task, Task<TResult>, ValueTask, ValueTask<TResult>)を待つ


---

## <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await 非同期処理


.left-split[

```js
function wait(waitSec) {
  return new Promise(resolve => {
    setTimeout(_ => {
      resolve(`resolved:${waitSec}`);
    }, waitSec);
  });
}

async function awaitNumsWait(nums) {
  while ( (num = nums.shift()) !== undefined ) {
    const result = await wait(num);
    console.log(`await [${result}]`);
  }
}

async function asyncNumsWait(nums) {
  Promise.all(nums.map(n => {
    return wait(n);
  })).then(values => {
	return console.log('asyncDone::' +values.join(','));
  });
}
```
]

.right-split[
```js
async function awaitAndAsyncCall(nums) {
  asyncNumsWait(nums);
  awaitNumsWait(nums);
  return 'awaitAndAsyncCall::done';
}

const nums = [2000, 1000, 1500];
awaitAndAsyncCall(nums).then((msg) => {
  console.log(msg);
});
```

<button class="btn btn-raised" onclick="player.open('/src/async_await.js');">Eval</button>
]



---

## async, await 非同期処理


<div class="text-center mermaid">
sequenceDiagram
    Process ->> AsyncLogic: 非同期処理
    AsyncLogic ->> BackendLogic1: 何かの処理
    AsyncLogic ->> BackendLogic2: 何かの処理
    Process ->> Process: 他の処理
    Note right of Process: 非同期処理を待たなければ、プロセスは終わる
</div>


---
## <i class="icon-go"></i> async, await: Example

.left-split[

```golang
func asyncFunc(sleep time.Duration, start time.Time) {
	time.Sleep(sleep)
	elapsed("async", start)
}

func awaitFunc(wg *sync.WaitGroup, sleep time.Duration, start time.Time) {
	defer wg.Done()
	// do somethig
	time.Sleep(sleep)
	elapsed("await", start)
}

```
]


.right-split[

```golang
func main() {
	start := time.Now()
	nums := []time.Duration{ 2 * time.Second, time.Second, 1500 * time.Millisecond }

	// async finished
	for _, sec := range nums {
		go asyncFunc(sec, start)
	}

	// await
	await := new(sync.WaitGroup)
	for _, sec := range nums {
		await.Add(1)
		go awaitFunc(await, sec, start)
	}
	await.Wait()

	println("finish")
}
```

<button class="btn btn-raised" onclick="player.open('/src/async_await.go');">Eval</button>

]



---
## <i class="icon-go"></i> async, await: Summary


<span class="bg-info h1">async</span>

* go routine で実行すれば大丈夫
* 待たないとそのままプロセスが終わります

<span class="bg-info h1">await</span>

* 標準パッケージの WaitGroup を利用すれば大丈夫
* 1つの処理なら単純に関数をそのまま呼ぶ


---

## いろんな言語の旅

.h3[
.lh16[
.text-muted[
* <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承
* <i class="icon-ruby text-danger"></i> Ruby block variable
* <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await
]]]

.h1[
* <i class="icon-java-duke"></i>  annotation, <i class="icon-python text-success"></i> decorator
]

---
## <i class="icon-java-duke"></i>  annotation, <i class="icon-python text-success"></i> decorator

.left-split[
```java
import java.lang.annotation.ElementType;
import java.lang.annotation.Target;

@Target(ElementType.METHOD) // メソッドにしか付けられません
public @interface Option {
    boolean isActive() default false;
}

class Api {

  @Option(isActive = true)
  public void get() String {
     return "get";
  }

  @Option
  public void put() String {
     return "put";
  }
}
```
]

.right-split[

```python
def api_option(fn):
    def func_wrapper():
        return "wrap[{0}]".format(fn())
    return func_wrapper

@api_option
def say():
    return 'say'

def hello():
    ''' no decrated function '''
    return 'hello'
```

<button class="btn btn-raised" onclick="player.open('/src/annotation.java','/src/decorator.py');">Eval</button>

]

---
## <i class="icon-java-duke"></i>  annotation

<div class="text-center mermaid">
graph LR;
	start((Start)) --> cb[Call API];
	cb --> ann{Annotation?}
	ann --> |yes| do[Related Annotation logic]
	ann --> |no| no[not work]
	do --> fin((End));
	no --> fin((End));

    style start fill:#2cb5e8;
    style fin fill:#2cb5e8;
</div>

```java
class Caller {
    public void call(Object instance, String methodName) {
        try {
            Method method = instance.getClass().getMethod(methodName);
            Option option = method.getAnnotation(Option.class);
            String result = option.isActive()
                ? String.format("warp[%s]", method.invoke(instance))
                : method.invoke(instance).toString();

            System.out.println("Result[" + result + "]");
        } catch (Throwable cause) {
            throw new RuntimeException(cause);
        }
    }
}
```


---
## <i class="icon-python text-success"></i> decorator


<div class="text-center mermaid">
graph LR;
	start((Start)) --> cb[Call API];
	cb --> wrap[Run wrapped code];
	wrap --> fn[Run original function code];
	fn --> fin((End));

    style start fill:#2cb5e8;
    style fin fill:#2cb5e8;
</div>

```python
def api_option(fn):
    def func_wrapper():
        return "wrap[{0}]".format(fn())
    return func_wrapper

@api_option
def say():
    return 'say'
```


---
## <i class="icon-go"></i>  anotation, decorator: Example


```golang
// api function
type ApiFunc func() string

// function adpter
func ApiOption(fn ApiFunc, isActive bool) ApiFunc {
  if isActive {
    return func() string {
       return fmt.Sprintf("wrap [%s]", fn())
    }
  }
  return fn  
}

func Get() string { return "get" }

type Api struct {
    Get ApiFuc `option:"active=true"`
    Put ApiFuc `option:"active=false"`
}


```

<button class="btn btn-raised" onclick="player.open('/src/decopartor_annotation.go');">Eval</button>


---
## <i class="icon-go"></i>  anotation, decorator: Summary

<span class="bg-info h1">annotation</span>

* struct tag で振る舞いを変更できる
* ※ 例は良くないので忘れてください  
  encoding/json とかを参照ください

```go
type Person struct {
    Id       int    `json:"id"` `db:"id"`
    Name     string `json:"name"` `db:"name"`
    Age      int `json:"age"` `db:"age"`
}
```


<span class="bg-info h1">decorator</span>

* Function Adapterで大丈夫

---

## いろんな言語の旅

.h3[
.lh16[
.text-muted[
* <i class="icon-java-duke"></i> <i class="icon-python text-success"></i> Class(クラス) の継承
* <i class="icon-ruby text-danger"></i> Ruby block variable
* <i class="icon-javascript-alt" style="color:gold;"></i> <i class="icon-csharp text-success"></i> async, await
* <i class="icon-java-duke"></i>  annotation, <i class="icon-python text-success"></i> decorator
]]]

.h1[
* <i class="icon-java-duke"></i> interface, <i class="icon-python text-success"></i> duck typing
]

---
## <i class="icon-java-duke"></i>  interface, <i class="icon-python text-success"></i> duck typing


.left-split[

```java
interface Animal {
    void run();
}

class Dog implements Animal {
    @Override
    public void run() {
        System.out.println("wow");
    }
}

class Cat implements Animal {
    @Override
    public void run() {
        System.out.println("miao");
    }
}

Animal dog = new Dog();
Animal cat = new Cat();
dog.run();
cat.run()
```
]

.right-split[

```python
'''
class Animal(metaclass=ABCMeta):

    @abstractmethod
    def run(self):
        pass

class Dog(Animal):

    def run(self):
        print('wow')
'''

class Dog:

   def run():
       print('wow')

class Cat:

   def run():
       print('miao')


Dog().run()
Cat().run()

```
]

<button class="btn btn-raised" onclick="player.open('/src/interface.java', '/src/ducktype.py');">Eval</button>


---

## <i class="icon-go"></i> interface, duck typing: Example

.left-split[

```golang
// Animal interface
type Animal interface {
	Run()
}

// Dog
type Dog struct{}
func (d Dog) Run() { println("wow") }

// Cat
type Cat struct{}
func (c Cat) Run() { println("miao") }

// Runnable interface
type Runnable interface {
	Run()
}

type Show struct{}
func (s Show) Run() { println("show")}
type Update struct{}
func (u Update) Run() { println("update")}
```
]


.right-split[

```golang
func main() {
	println("-- Animal --")
	var animal Animal
	animal = new(Dog)
	animal.Run()
	animal = new(Cat)
	animal.Run()

	println("-- Runnable --")
	var run Runnable
	run = new(Show)
	run.Run()
	run = new(Update)
	run.Run()

	println("-- Duck Typing --")
	run = new(Dog)
	run.Run()
	animal = new(Show)
	animal.Run()
}
```

<button class="btn btn-raised" onclick="player.open('/src/interface_ducktype.go');">Eval</button>

]


---
## <i class="icon-go"></i> interface, duck typing: Summary

<span class="bg-info h1">Interface</span>

* 呼び出し側へ明示的に関数の定義を公開できる
* 実装対象の責務範囲を明確にできる

<span class="bg-info h1">DuckTyping</span>

* 呼び出し側は対象が何に対してのものか？理解する
* 実装対象は曖昧に実装が可能

---
class: eye-bg

## <span class="bg-info">小まとめ</span>

.h1[

* interface, duck typing をサポートしているよ
]


---
class: center, middle

.display-1[
過去の

.en[Trend] な旅 編
]

.title-img[![gopher](images/gophercolor.png)]


---
class: center, middle, eye-bg

.big[<span class="en">Linq / ReactiveX</span>]


---

## <span class="en">Linq (Language Integrated Query)</span>

データに集合対して

* 標準化した方式で賭合わせを可能にした技術
* 言語レベルで高階関数を利用して問い合わせできるもの

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


### <span class="en">References</span>

- [Wikipedia EN](https://en.wikipedia.org/wiki/Language_Integrated_Query) /  [Wikipedia JA](https://ja.wikipedia.org/wiki/%E7%B5%B1%E5%90%88%E8%A8%80%E8%AA%9E%E3%82%AF%E3%82%A8%E3%83%AA)
- [C# Programming concepts Language Integrated Query (LINQ)](https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/) <i class="fab fa-windows"></i>
- <a class="badge badge-pill badge-info" href="https://github.com/ahmetb/go-linq" target="_blank">go-linq godoc</a> / <a class="badge badge-pill badge-info" href="https://godoc.org/github.com/ahmetb/go-linq" target="_blank">github</a>

---
## <span class="en">Linq Java</span>

```java
// 文字列Listを , で結合した文字列を返す関数
Function<IntStream,String> join = (s) -> {
	return String.join(",",
					   s.mapToObj(Integer::toString)
					   .collect(Collectors.toList()));
};

// 1 .. 9 までの配列
int[] data = IntStream.range(1, 10).toArray();

// stream の作成
IntStream linq = IntStream.of(data);

// 内容表示：配列内容を全部
System.out.println("data    : " + join.apply(linq));

// 4 より大きい数字のみフィルタ
linq = IntStream.of(data).filter(i -> i > 4);
// 内容表示
System.out.println("filtered: " + join.apply(linq));
```

<button class="btn btn-raised" onclick="player.open('/src/linq.java');">Eval</button>


---
## <span class="en">Linq Example</span>


go-linq を元に作成（汎用性を減らしてコピペ）してます

.h3[
.lh16[

* query
   * クエリを表すもの
* from(data []int)
   * data をクエリに変換
* selectBy()
   * １データへの操作
* where()
   * 条件絞り込み
* orderBy()
   * ソート
]]

---
## <span class="en">Linq Example</span>

.left-split[

```golang
// 反復を繰り返す関数
type itr func() (int, bool)

// 反復可能なクエリ
type query struct {
	itr func() itr
}

// map変換関数を受けて各要素で関数実行の反復を繰り返す
func (q query) selectBy(mapFn func(int) int) query {
	return query{
		itr: func() itr {
			next := q.itr()
			return func() (int, bool) {
				ret, ok := next()
				if ok {
					ret = mapFn(ret)
				}
				return ret, ok
			}
		},
	}
}

```
]

.right-split[
```go
// []int からクエリを作成する
func from(data []int) query {
	len := len(data)
	return query{
        // 実行する度に index をインクリメントして値を返す関数
		itr: func() itr {
			index := 0
			return func() (int, bool) {
				ok := index < len
				val := 0
				if ok {
					val = data[index]
					index++
				}
				return val, ok
			}
		},
	}
}
```
]

---
## <span class="en">Linq Example</span>

.left-split[
```go
// フィルタ関数でフィルタして反復を繰り返す
func (q query) where(filter func(i int) bool) query {
	return query{
		itr: func() itr {
			next := q.itr()
			return func() (int, bool) {
				var (
					val int
					ok  bool
				)
				for val, ok = next(); ok; val, ok = next() {
					if filter(val) {
						return val, ok
					}
				}
				return val, ok
			}
		},
	}
}
```
]

.right-split[

```go
// ソートを行う
func (q query) orderBy(less func(i, j int) bool) query {
	var nums []int
	q.apply(&nums)
	sort.Sort(sorter{nums: nums, less: less})
	return from(nums)
}

// ソート用
type nums []int
func (ns nums) Len() int      { return len(ns) }
func (ns nums) Swap(i, j int) { ns[i], ns[j] = ns[j], ns[i] }

type sorter struct {
	nums
	less func(i, j int) bool
}
func (s sorter) Less(i, j int) bool {
	return s.less(s.nums[i], s.nums[j])
}
```

]

---
## <span class="en">Linq Example</span>


.left-split[
<div class="text-center mermaid">
graph TD;
    start((data)) --> f1[selectBy next * 2];
    f1 --> f2[where next > 50];
    f2 --> odr[orderBy  sort DESC];
    odr --> fin((Finish));

    style start fill:#2cb5e8;
    style fin fill:#2cb5e8;
</div>
]

.right-split[

```go

func main() {
	data := []int{1, 10, 43, 28, 32, 5}
	var values []int
	from(data).apply(&values)
	fmt.Println("data: ", values)

	var selectValues []int
	from(data).selectBy(func(num int) int {
		return num * 2
	}).where(func(num int) bool {
		return num >= 50
	}).orderBy(func(i, j int) bool {
		// DESC
		return i > j
	}).apply(&selectValues)
	fmt.Println("select data: ", selectValues)
}
```

<button class="btn btn-raised" onclick="player.open('/src/linq.go');">Eval</button>

]

---
## <span class="en">Linq Summary</span>

<span class="bg-info h1"><i class="icon-go"></i>golang</span>

* 言語標準パッケージでの実装は無い
* <code class="remark-inline-code inline-b">go get github.com/ahmetb/go-linq</code> で取得しては利用可能
* .text-muted[個人的にSliceにその機能があると嬉しい]


---
## <span class="en">ReactiveX</span>

<div class="text-center mermaid">
graph LR;
    start((Start)) --> e1((Event1));
    e1((Event1)) --> e2((Event2));
    e2((Event2)) --> e3((Event3));
    e3((Event3)) --> fin((Finish));

    style start fill:#2cb5e8;
    style fin fill:#2cb5e8;
</div>


* .en[Observer, Iterator] パターンと .en[Functional Programming] のアイデアを組み合わせたもの


#### .en[References]

- .en[[ReactiveX](http://reactivex.io/)]
   - [リアクティブ宣言](https://www.reactivemanifesto.org/ja)
- .en[[RxGo github](https://github.com/ReactiveX/RxGo)]

---
## .en[ReactiveX::RxJS]


```javascript
document.addEventListener('click', () => console.log('Clicked!'));
  ↓↓↓
import { fromEvent } from 'rxjs';

fromEvent(document, 'click').subscribe(() => console.log('Clicked!'));
```


```js
import { Observable } from 'rxjs';

const observable = new Observable(subscriber => {
  subscriber.next(1); subscriber.next(2); subscriber.next(3);
});

console.log('just before subscribe');
observable.subscribe({
  next(x) { console.log('got value ' + x); },
  error(err) { console.error('something wrong occurred: ' + err); },
  complete() { console.log('done'); }
});
console.log('just after subscribe');```

```bash
just before subscribe
got value 1
got value 2
got value 3
just after subscribe
done
```

---
## <span class="en">ReactiveX Exmaple</span>

.left-split[

```go
func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	value := make(chan bool)
	done := make(chan bool)

	go func() {
		defer func() { done <- true }()
		num := 0
		for {
			select {
			case <-value:
				num++
				if num >= 10 {
					return
				} else {
					fmt.Println("wait...[", num, "]")
				}
			}
		}
	}()
```
]

.right-split[
```go
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t.Format("15:04:05"))
			value <- true
		}
	}
}
```
]

---
## <span class="en">ReactiveX Exmaple</span>

<div class="text-center mermaid">
graph LR;
    start((Event)) --> cb[Create Observable]
	cb --> s[Subscribe connect to observer];
	s --> fin((Continue));

    style start fill:#2cb5e8;
    style fin fill:#2cb5e8;
</div>

<div class="text-center mermaid">
graph LR;
	start((Continue)) --> ob{Observable Error?};
	ob --> |yes| do[Observer # onError]
	ob --> |no| no[Observer # onNext]
	do --> done[onCompleted];
	no --> done[onCompleted];
	done --> fin((End));

    style start fill:#2cb5e8;
    style fin fill:#2cb5e8;
</div>

1. Event ソースから Observable を作成する
2. handler から Subscribe() で Observer 接続する


---
## <span class="en">ReactiveX Exmaple</span>

.left-split[

```go
type (
    // Event Handle funcs
	Next  func(interface{})
	Done  func()
	Error func(error)

	observable <-chan interface{}

	iterable <-chan interface{}
)

func (it iterable) Next() (interface{}, error) {
	if next, ok := <-it; ok {
		return next, nil
	}
	return nil, ErrEndOfInterator
}

// Event Handlers
type observer struct {
	next Next
	done Done
	err  Error
}
```
]

.right-split[

```go
// 元のコードは onNext, onError, onCompleted や オプションの引数から
// 内部で observer を接続している
func (o observable) subscribe(ob observer) chan subscription {
	done := make(chan subscription)
	sub := subscription{}
    // イベントを受けて、エラーの場合は onError, 値は onNext に通知
	go func() {  // 元のコードはOptionに応じて、concurent で実行
	OuterLoop:
		for item := range o {
			switch item := item.(type) {
			case error:
				ob.err(item)
				sub.Error = item
				break OuterLoop
			default:
				ob.next(item)
			}
		}

		if sub.Error == nil {
			ob.done()
		}
		done <- sub.End()
	}()
	return done
}
```
]

---
## <span class="en">ReactiveX Example</span>

.left-split[

```go
func just(items ...interface{}) observable {
	source := make(chan interface{})
	go func() {
		for _, item := range items {
			source <- item
		}
		close(source)
	}()

	return observable(source)
}
```

]

.right-split[

```go
func from(it iterable) observable {
	source := make(chan interface{})
	go func() {
		for {
			val, err := it.Next()
			if err != nil {
				break
			}
			source <- val
		}
		close(source)
	}()
	return observable(source)
}
```

]

---
## <span class="en">ReactiveX Example</span>


.left-split[

```go
func justSubscribe() {
	score := 0

	observer := handler(
		// next
		func(n interface{}) {
			if num, ok := n.(int); ok {
				score += num
			}
		},
		// done
		func() {
			score *= 2
		},
	)
	sub := just([]interface{}{1, 2, 3, 4, 5}).subscribe(observer)
	<-sub

	fmt.Println("score:", score)
	close(sub)
}
```
<button class="btn btn-raised" onclick="player.open('/src/rx.go');">Eval</button>

]

.right-split[

```go
func ticker() {
	score := 0
	observer := handler(/* same code */)
	value := make(chan interface{})
	itr, _ := toItr(value)
	sub := from(itr).subscribe(observer)
	defer close(sub)

	ticker := time.NewTicker(time.Second)
	go func() {
		defer ticker.Stop()
		num := 0
		for {
			select {
			case t := <-ticker.C:
				num++
				fmt.Println("Current time: ", t.Format("15:04:05"), ",num:", num)
				if num > 5 {
					close(value)
					return
				}
				value <- num
			}
		}
	}()
	<- sub
	fmt.Println("score:", score)
}
```

]


---
## <span class="en">ReactiveX Sumamry</span>

<span class="bg-info h1">Golang</span>

* <code class="remark-inline-code inline-b">go get -u github.com/reactivex/rxgo</code> で取得しては利用可能
* .text-danger[This is an early project and your contributions will help shape its direction.] なので注意

<span class="bg-info h1">ReactiveX(FRP)</span>

* いろいろ大変（概念理解したり、実装力あげたり）
* 全てをストリームと見たり...
.text-center[
<image src="images/everything-is-stream.jpg" width="20%" />
]

---
class: center, middle, eye-bg

.display-1[
旅の終わり

まとめ
]

---

## 旅のまとめ

<span class="bg-info h1">Golang の旅</span>

* 簡潔に書く

<span class="bg-info h1">Lang な旅</span>

* SubClass を持たない
* block 変数とか特殊な変数やその記述はない
* async, await はサポートしない
* annotation, decorator はサポートしない  
  → StructTagをサポート <i class="icon-go"></i>
* interface, duck typing をサポート <i class="icon-go"></i>

<span class="bg-info h1">Trend な旅</span>

* 簡単に書けるよ
* 使うかは・・・あなた次第・・・


---
class: center, middle

.h1[
サポート機能少ない

けど簡単にかける
]

.big[
なんでか？
]

---

## まとめ と .en[Features in Go] <i class="icon-go"></i>

<div class="card">

<div class="card-body bg-light border-rounded">

<div class="h2">
<p>Go は違うものです</p>
<p class="text-muted">Go is different.</p>
<br>
<p>他の言語のようににはならないよ</p>
<p class="text-muted">Go does not try to be like the other languages.</p>
<br>
<p>機能を競わないよ</p>
<p class="text-muted">Go does not compete on features.</p>
</div>
</div>

</div>

<br/>

.text-center[
.en[by Rob Pike]
]

---

## 旅のまとめ からみる Goの機能

.left-split[

<span class="bg-info h1">Golang の旅</span>

* 簡潔に書く

<span class="bg-info h1">Lang な旅</span>

* SubClass を持たない
* block 変数とか特殊な変数やその記述はない
* async, await はサポートしない
* annotation, decorator はサポートしない  
  → StructTagをサポート <i class="icon-go"></i>
* interface, duck typing をサポート <i class="icon-go"></i>

<span class="bg-info h1">Trend な旅</span>

* 簡単に書けるよ
* 使うかは・・・あなた次第・・・

]

.right-split[

.h2[
<span class="bg-warning text-white h1">デメリット</span>

* 言語機能が増加すると覚える・することも増加
   * 保守性、可読性の減少

<br><br>

<span class="bg-primary text-white h1">やりたいこと</span>

* ものを作ったり、問題を解決したりしたい
* Trendの取り込みや技法の使いたい

<br><br>

<span class="bg-danger text-white h1">モチベ・思想・哲学</span>

* What???

]

]

---

## まとめ

<div class="text-center display-2">
<span class="bg-info pl-4 pr-4 pt-1 pb-1 en">
Simplicity
</span>
</div>

<br>

.en[
.text-center[
Go is simple, at least compared to established languages.

Simplicity has many facets.
]]

<br>

<div class="text-center display-2">
<span class="bg-info pl-4 pr-4 pt-1 pb-1 en">
Simplicity is Complicated
</span>
</div>

.text-center[
<a class="badge badge-pill badge-info" href="https://talks.golang.org/2015/simplicity-is-complicated.slide#1" target="_blank">
Simplicity is Complicated <i class="inline-b material-icons" style="display: inline-flex; vertical-align: middle;">exit_to_app</i>
</a>

<small class="en">by Rob Pike</small>
]

---
class: center, eye-bg


.display-2[

<br>

Go code that is

- simple,
- readable,
- maintainable.
]

---
class: center, eye-bg

<br>

.display-2[
<p class="pt-3 pb-3 bubble" style="line-height: 1.6em;">
<i class="icon-python text-success"></i>
<i class="icon-java-duke"></i>
<i class="icon-ruby text-danger"></i>
<i class="icon-javascript-alt" style="color:gold;"></i>
<i class="icon-csharp text-success"></i>
<i class="icon-go text-info"></i>
<br>
各言語ディスっている<br>
内容あればすみません

</p>
]

---
class: center, middle, eye-bg

.display-2[
<p class="pt-3 pb-3 bubble" style="line-height: 1.6em;">
ご静聴ありがとうございました
</p>
]

---

## .en[Simplicity] と .en[Generics]

.en[Declaration]

```go
type List(type T) []T

func Keys(type K, V)(m map[K]V) []K
```

.en[Uses]

```go
var ints List(int)

keys := Keys(int, string)(map[int]string{1:"one", 2: "two"})
```

.text-right[
<a class="badge badge-pill badge-info" href="https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md" target="_blank">
Generics — Problem Overview <i class="inline-b material-icons" style="display: inline-flex; vertical-align: middle;">exit_to_app</i>
</a>
]

.left-split[
* 個人的な気持ちでは嬉しい
* Simplicity を崩してしまうのでは無いか？とちょっと心配しています
]

.right-split[
<img src="images/mic-drop.png" width="40%">
]


---
## 余談

.text-center[
.display-3[
<p class="pt-3 pb-3 bubble" style="line-height: 1.6em;">
ものごとはハートで見なくちゃいけない、<br/>

っていうことなんだ。<br/>

大切なことは、目に見えないからね
</p>
]

<br/>

**引用:** _星の王子様_
]
