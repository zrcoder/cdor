import (
	"strings"
)

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

arr := strings.fields(shapes)

for sh <- arr {
	// N("cloud", Opt().Sh("cloud")),
	echo `N("${sh}", Opt().Sh("${sh}")),`
}

/* arrow head shapes
triangle (default)
Can be further styled as style.filled: false.
arrow (like triangle but pointier)
diamond
Can be further styled as style.filled: true.
circle
Can be further styled as style.filled: true.
cf-one, cf-one-required (cf stands for crows foot)
cf-many, cf-many-required
*/
