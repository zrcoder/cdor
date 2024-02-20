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
