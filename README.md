Syaro Markdown Wiki Server
====

[![Build Status](https://drone.io/github.com/OUCC/syaro/status.png)](https://drone.io/github.com/OUCC/syaro/latest)

Syaro is a simple and pretty wiki system supporting markdown.

![Markdown Viewer](sample/syaro.png)

Features
----
### Markdown Viewer
Syaro can handle markdown format files. [blackfriday] is used to convert
Markdown to HTML.

Viewer supports MathJax. LaTeX text surrounded by `$` (inline math) or `$$` 
(display math) is converted to beautiful mathematical expression.

Viewer also supports code syntax highlighting. This feature is powered by
[highlight.js].

### `[[WikiLink]]`
Texts surrounded by double bracket are interpreted as WikiLink. To link to
another wiki page, you can use WikiLink. 
another wiki page, you can use WikiLink. For example,

```
[[My Profile Page]]
```

is a link to `My Profile Page.[md,mkd,markdown]`. Syaro searches all files under
wikiroot, so two files haven't be in same folder.

### Files/Folders under ...
In main page of a directory, file/folder list is automatically appended to
page's end.

Example file tree

```
WIKIROOT/ - Home.md
            kinmoza.md
            gochiusa/ - gochiusa.md
                        cocoa.md
                        chino.md
                        rize.md
                        syaro.md
                        chiya.md
```

In this example, main page of `gochiusa/` is `gochiusa.md`, and file/folder list
is appended when you see `gochiusa.md`.

### Print ready
![Print preview](sample/printing.png)

### Powerful Markdown editor
* Realtime preview (including MathJax rendering and code highlighting)
* Markdown syntax highlight
* User fiendly table editor (inspired by 
[Sublime Table Editor](https://sublime.wbond.net/packages/Table%20Editor))
  - auto format (alignment, cell size) table
  - move cursor by tab/shift-tab
  - support full width character (ex. CJK)

![Editor](sample/editor.png)

![Table editor](sample/table.gif)

### Supports MathJax and code highlighting
See [[Sample Page]]

VS. [Gollum](https://github.com/gollum/gollum)
----
* Syaro supports CJK filename and text.
* File list on wiki page
* User friendly Markdown editor
* Table editor

Build & Install
----
Required packages: go, bower, cmake, pkg-config

```bash
go get github.com/OUCC/syaro
cd $GOPATH/src/github.com/OUCC/syaro
./build.sh
```

Usage
----
```bash
syaro --wikiroot=/path/to/wiki

# or
sudo mkdir /usr/local/share/syaro
sudo cp public views /usr/local/share/syaro/ # place template html etc in your system
cd path/to/wiki
syaro

# If you want to use MathJax or highlight.js,
# syaro --mathjax --highlight --wikiroot=...
```

Then open `localhost:8080/Home` in your browser.

`syaro -h` or `syaro --help` you can see more options.

Contribution
----
Fork and pull requests welcome. I hadn't receive any pullreq ever so please give
me your first pullreq!

Donate
----
* [Donate $3](https://gumroad.com/l/Jwtx)
* [誰かに買って欲しいものリスト](http://www.amazon.co.jp/registry/wishlist/1MVMC2QBIJYY)

About
----
### Author
* @yuntan
* @susisu (Table editor)

This software is released under MIT License.

Following softwares are used:

* [Go]  (BSD)
* [blackfriday]  (BSD)
* [go-logging]
* [Twitter Bootstrap]  (MIT)
* [jQuery]  (MIT)
* [Ace]  (BSD)
* [East Asian Width]  (MIT)
* [MathJax]  (Apache)
* [highlight.js]  (BSD)


[Go]: http://golang.org/
[blackfriday]: https://github.com/russross/blackfriday
[go-logging]: https://github.com/op/go-logging
[Twitter Bootstrap]: http://getbootstrap.com
[jQuery]: http://jquery.com
[Ace]: http://ace.c9.io
[East Asian Width]: https://github.com/komagata/eastasianwidth
[Mathjax]: http://www.mathjax.org/
[highlight.js]: https://highlightjs.org/
