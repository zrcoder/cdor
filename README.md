# cdor

Write Go+/Go code to make diagrams.

## cdor vs [d2](https://d2lang.com)

`cdor` is inspired and powered by `d2`, the difference is:

> d2 is a `DSL` (Domain Specific Language), But cdor is `SDF` (Specific Domain Friendliness), we just write common programming laguages (Go+/Go) to generate diagrams.

## examples

![hello](examples/hello.svg)
![id](examples/example-id.svg)
![shape](examples/shape.svg)
![connections](examples/connections.svg)
![containers](examples/containers.svg)

text & code:
![markdown](examples/md.svg)
![latext](examples/latex.svg)

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

- [x] arrorw options
    - [ ] arrow head filled
- [ ] special shapes
    - [x] text/code
- [ ] surpport worker classfile
