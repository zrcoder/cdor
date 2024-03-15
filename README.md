# cdor

Write Go+/Go code to make diagrams.

## cdor vs [d2](https://d2lang.com)

`cdor` is inspired and powered by `d2`, the difference is:

> d2 is a `DSL` (Domain Specific Language), But cdor is `SDF` (Specific Domain Friendliness), we just write common programming laguages (Go+/Go) to generate diagrams.

## examples

![hello](doc/examples/hello.svg)
![md](doc/examples/md.svg)
![latex](doc/examples/latex.svg)
![sql_table](doc/examples/sql_table.svg)
![class](doc/examples/class.svg)
![jsonn](doc/examples/jsonn.svg)
![obj](doc/examples/obj.svg)
![grid](doc/examples/grid.svg)
![sequence](doc/examples/sequence.svg)
![shapes](doc/examples/shapes.svg)
![icon](doc/examples/icon.svg)

see [doc](doc) for details.

## usage

Initialize your diagram project

```shell
mkdir demo
cd demo
gop mod init demo
gop get github.com/zrcoder/cdor@latest
```

### cases

#### single diagram

create `main_cdor.gox`, with content like:

```c
con "x", "y"

saveFile "hi.svg"
```

then run:

```shell
gop mod tidy
gop run .
```

you'll find `hi.svg` generated:

![hi](doc/usage/single.svg)

#### multiple diagrams

create `hi1_cdor.gox` with content:
```c
con "x", "y"
```

create `hi2_cdor.gox` with content:
```c
con "a", "b"
```

modify `main_cdor.gox` to:
```c
saveFiles ""
```

after `gop run`, you'll find `hi1.svg` and `hi2.svg` generated.

#### merge diagrams

you can merge all diagrams into a single diagram, and even creat more nodes and connections based on the sub diagrams, for example, modify `main_cdor.gox` to:

```c
con "a", "x"
merge().saveFile "merged.svg"
```

after `gop run`, we got a generated `merged.svg`:

![merged](doc/usage/merged.svg)

we can also merge specific diagrams by their names:

```c
merge("hi1", "hi2").saveFile("res.svg")
```

## config

```c
cfg.direction("right").sketch().elkLayout().theme(105).darkTheme(200)
```

directions: `down`(default), `up`, `left`, `right`

theme ids:

| Theme | ID |
|---|---|
| Neutral default |  0 |
| Neutral Grey |  1 |
| Flagship Terrastruct |  3 |
| Cool classics |  4 |
| Mixed berry blue |  5 |
| Grape soda |  6 |
| Aubergine |  7 |
| Colorblind clear |  8 |
| Vanilla nitro cola |  100 |
| Orange creamsicle |  101 |
| Shirley temple |  102 |
| Earth tones |  103 |
| Everglade green |  104 |
| Buttered toast |  105 |
| Terminal |  300 |
| Terminal Grayscale |  301 |
| Origami |  302 |
| Dark Mauve |  200 |
| Dark Flagship Terrastruct |  201 |

