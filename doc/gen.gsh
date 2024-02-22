#!/usr/bin/env gop run

shapes := `
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
`

for sh <- shapes.fields {
	echo `node("${sh}").shape("${sh}"),`
}

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

echo "---"

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