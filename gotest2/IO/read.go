package IO

import (
	"bytes"
	"fmt"
	"os"
)

/*
  @Description:
*/

func main() {
	// buffer both implements io.Reader and io.Writer
	var b bytes.Buffer
	b.Write([]byte("Hello"))

	// Fprintf write string to io.Writer
	fmt.Fprintf(&b," World")
	b.WriteTo(os.Stdout)

}