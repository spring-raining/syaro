Sample Page
====

Title
----
### Title with ID {#titleWithID}
link to [Title with ID](#titleWithID)

#### h4

##### h5

###### h6

Plain text
----
Lorem ipsum dolor sit amet, consectetur adipiscing elit. 

Aliquam mi orci, porta vitae nisl sit amet, imperdiet fringilla dolor.

Rich text
----
Lorem *ipsum* **dolor** ***sit*** ~~amet~~, consectetur adipiscing elit[^lorem-ipsum].

### "smarty pants" [^smarty-pants]
SmartyPants can perform the following transformations:

* Straight quotes ( `"` and `'` ) into "curly" quote HTML entities
* Backticks-style quotes (` ``like this'' `) into "curly" quote HTML entities
* Dashes (`--` and `---`) into en-- and em---dash entities
* Three consecutive dots (`...`) into an ellipsis entity...
* Slash separated digits (`1/2`) into fraction HTML entities (1/2)

### Link
[google](http://google.co.jp)

### List
* item 1
* item 2
    - item 2-1
        + item 2-1-1
    - item 2-2
* item 3

1. item 1
2. item 2
    1. item 2-1
    2. item 2-2

#### Task list

* [ ] checkbox
* [x] checkbox2

### Quote
> quote

> > nest

> end

### Table

| Left align | Right align | Center align |
|:-----------|------------:|:------------:|
| This       |        This |     This     |
| column     |      column |    column    |
| will       |        will |     will     |
| be         |          be |      be      |
| left       |       right |    center    |
| aligned    |     aligned |   aligned    |

### Horizontal line

---

Code
----
To use code highlighter, run `syaro --highlight`

```HTML
<!DOCTYPE html>
<head>
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title></title>
```

```css
body { display: none; }
```

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, golang")
}
```

`<code></code>`

WikiLink
----
* [[Home]]
* [[page not available]]
* [[さんぷるぺーじ]](Sample Page)

Math
----
To use [MathJax](http://www.mathjax.org/), run `syaro --mathjax`.

inline math $\mathrm{e}^{i\theta}=\cos\theta+i\sin\theta$

$$ S=\sum^\infty_{n=1}s_n $$

[^lorem-ipsum]: In publishing and graphic design, lorem ipsum is a filler text commonly used to demonstrate the graphic elements of a document or visual presentation.
[^smarty-pants]: http://daringfireball.net/projects/smartypants/
