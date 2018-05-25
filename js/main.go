package main

import (
	"github.com/emitter-io/stats"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("stats", map[string]interface{}{
		"restore": Restore,
	})

}

// Restore restores the snapshot
func Restore(snapshot []byte) (out []*js.Object) {
	snapshots, err := stats.Restore(snapshot)
	if err != nil {
		panic(err)
	}

	for _, s := range snapshots {
		out = append(out, js.MakeWrapper(&s))
	}
	return
}
