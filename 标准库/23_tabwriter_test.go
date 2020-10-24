package stand

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func Test_tabwrite_01(t *testing.T) { // 左对齐
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ',1)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmith\t")
	w.Flush()

}


func Test_tabwrite_02(t *testing.T) { // 右对齐
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ',tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmith\t")
	w.Flush()

}
