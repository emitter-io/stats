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
func Restore(snapshot []byte) map[string]*js.Object {
	snapshots, err := stats.Restore(snapshot)
	if err != nil {
		panic(err)
	}

	// Create a snapshot map for easier consumption
	out := make(map[string]*js.Object)
	for _, s := range snapshots {
		out[s.Name()] = js.MakeWrapper(&s)
	}
	return out
}
