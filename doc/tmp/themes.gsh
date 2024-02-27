/* shapes
rectangle
square
page
parallelogram
document
cylinder
queue
package
step
callout
stored_data
person
diamond
oval
circle
hexagon
cloud
*/

/* arrow head shapes
triangle: (default) Can be further styled as style.filled: false.
arrow: (like triangle but pointier)
diamond: Can be further styled as style.filled: true.
circle: Can be further styled as style.filled: true.
cf-one: (cf stands for crows foot)
cf-one-required:
cf-many:
cf-many-required:
*/

/*specia shapes:
text/code/latex
node("code").code("tex", `\lim_{h \rightarrow 0 } \frac{f(x+h)-f(x)}{h}`)
*/

themes := `
Available themes:
Light:
- Neutral default: 0
- Neutral Grey: 1
- Flagship Terrastruct: 3
- Cool classics: 4
- Mixed berry blue: 5
- Grape soda: 6
- Aubergine: 7
- Colorblind clear: 8
- Vanilla nitro cola: 100
- Orange creamsicle: 101
- Shirley temple: 102
- Earth tones: 103
- Everglade green: 104
- Buttered toast: 105
- Terminal: 300
- Terminal Grayscale: 301
- Origami: 302
Dark:
- Dark Mauve: 200
- Dark Flagship Terrastruct: 201
`

echo "| Theme | ID |"
echo "|---|---|"

for line <- themes.Split("\n") {
	line = line.trimSpace
	line = line.trimLeft("- ")
	i := line.index(":")
	if i == -1 {
		continue
	}
	theme, id := line[:i], line[i+1:]
	if id == "" {
		continue
	}
	echo "| ${theme} | ${id} |"
}
