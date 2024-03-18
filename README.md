# cdor

Generate diagrams with Go+/Go code.

## cdor vs [d2](https://d2lang.com)

`cdor` is inspired and powered by `d2`, the difference is:

d2 is a `DSL` (Domain Specific Language), But cdor is `SDF` (Specific Domain Friendliness), we just write common programming laguages (Go+/Go) to generate diagrams.

## examples

![hello](doc/examples/hello.svg)
![md](doc/examples/md.svg)
![latex](doc/examples/latex.svg)
![sql_table](doc/examples/sql_table.svg)
![jsonn](doc/examples/jsonn.svg)
![icon](doc/examples/icon.svg)
![shapes](doc/examples/shapes.svg)
![near](doc/examples/near.svg)

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

after `gop run`, we got a generated `merged.svg`.

we can also merge specific diagrams by their names:

```c
merge("hi1", "hi2").saveFile("res.svg")
```

#### range diagrams

you can range all the diagrams:

```c
rangeCdors (name, cdor, err) => {
    // do something with cdor
}
```

or range part of the diagrams: 

```c
rangeCdors (name, cdor, err) => {
    // do somthing with cdor
}, "hi1", "hi2"
```
