package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

func hey() {
	r := rand.New(rand.NewSource(99))
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}
	show("Float32", r.Int(), r.Int(), r.Int())
}
