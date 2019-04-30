name: top
class: center, middle, top-background

<h1>Gopher と行く<br/> Lang な旅</h1>

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
.left-column[
  ## What is it?
  ## Why use it?
]
.right-column[
As the slideshow is expressed using Markdown, you may:

- Focus on the content, expressing yourself in next to plain text not worrying what flashy graphics and disturbing effects to put where

As the slideshow is actually an HTML document, you may:

- Display it in any decent browser

- Style it using regular CSS, just like any other HTML content

- Use it offline!

As the slideshow is contained in a plain file, you may:

- Store it wherever you like; on your computer, hosted from your Dropbox, hosted on Github Pages alongside the stuff you're presenting...

- Easily collaborate with others, keeping track of changes using your favourite SCM tool, like Git or Mercurial
]
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

  </div>
  <div class="col-md-4">

  </div>

</div>

---

# <span class="en">Introduction</span>

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
# エンジニアを募集しています
## 職種
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

# <span class="en">Beginer</span> 向けコンテンツになります


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
なので、遠慮なく

**質問** or **何か** 聞きたいとか

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

# Golang とは

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
  <td class="align-middle"><i class="icon-csharp"></i></td>
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

# Golang とは



* expressive
* concise
* clean
* efficient


<img src="images/what-is-golang.png" width="60%"/>

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
