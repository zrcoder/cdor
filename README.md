# cdor

Write Go+/Go code to make diagrams.

## cdor vs [d2](https://d2lang.com)

`cdor` is inspired and pwored by `d2`, the difference is:

> d2 is a `DSL` (Domain Specific Language), But cdor is `SDF` (Specific Domain Friendliness), we just write Go+/Go to generate diagrams.

## examples

```c
con("Go+", "Go").label("cdor")
```

<center><img src='hello.svg' width='62%'/></center>

see more in [example_test.go](example_test.go).

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

## TODO

- [x] Go+ classfile
- [x] arrorw options
    - [ ] arrow head filled
- [x] config api
    - [ ] elk layout
- [ ] special shapes
- [ ] surpport multiple cdor files in same directory
