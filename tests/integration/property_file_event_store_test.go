package integration

import (
	"fmt"
	"os"
	"testing"
	"overhelloworld/internal/app"
	"overhelloworld/internal/domain"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestFileEventStore_Sequence_Property(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 10
	props := gopter.NewProperties(parameters)

	props.Property("All messages in a sequence are persisted and replayed in order", prop.ForAll(
		func(msgs []string) bool {
			os.Remove("test_seq_events.jsonl")
			store := &app.FileEventStore{Path: "test_seq_events.jsonl"}
			for i, m := range msgs {
				evt := domain.HelloSaid{ID: fmt.Sprintf("%d", i), Message: m}
				if err := store.Append(evt); err != nil {
					return false
				}
			}
			var replayed []string
			if err := store.Replay(func(evt map[string]interface{}) {
				if m, ok := evt["Message"].(string); ok {
					replayed = append(replayed, m)
				}
			}); err != nil {
				return false
			}
			if len(replayed) != len(msgs) {
				return false
			}
			for i := range msgs {
				if msgs[i] != replayed[i] {
					return false
				}
			}
			return true
		},
		gen.SliceOf(gen.RegexMatch("[a-zA-Z0-9 .,!?:;'\"-]{0,100}")),
	))

	props.TestingRun(t)
}

func TestFileEventStore_EmptyAndLarge_Property(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 10
	props := gopter.NewProperties(parameters)

	props.Property("Empty and large messages are persisted and replayed", prop.ForAll(
		func(msg string) bool {
			os.Remove("test_edge_events.jsonl")
			store := &app.FileEventStore{Path: "test_edge_events.jsonl"}
			evt := domain.HelloSaid{ID: "edge", Message: msg}
			if err := store.Append(evt); err != nil {
				return false
			}
			var replayed []string
			if err := store.Replay(func(evt map[string]interface{}) {
				if m, ok := evt["Message"].(string); ok {
					replayed = append(replayed, m)
				}
			}); err != nil {
				return false
			}
			if len(replayed) != 1 {
				return false
			}
			return replayed[0] == msg
		},
		gen.OneConstOf("", string(make([]byte, 10000))),
	))

	props.TestingRun(t)
}
