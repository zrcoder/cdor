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
