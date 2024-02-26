# usage

## pre
Initialize your diagram project

```shell
mkdir demo
cd demo
gop mod init demo
gop get github.com/zrcoder/cdor@latest
```

## cases

### single diagram

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

you'll find `hi.svg` generated, rendered like:

![hi](single.svg)

### multiple diagrams

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
saveFiles
```

after `gop run`, you'll find `hi1.svg` and `hi2.svg` generated.

### merge diagrams

you can merge all diagrams into a single diagram, and even creat more nodes and connections based on the sub diagrams, for example, modify `main_cdor.gox` to:

```c
con "a", "x"
merge
saveFile "merged.svg"
```

after `gop run`, we got a generated `merge.svg`:

![merged](merged.svg)
