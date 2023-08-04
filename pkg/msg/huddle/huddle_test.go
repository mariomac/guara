package huddle

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHuddle(t *testing.T) {
	// GIVEN a huddle
	h := Huddle[string]{}

	// with multiple people joining
	m1, m2, m3 := h.Join(), h.Join(), h.Join()

	// WHEN a member says something
	m1.Say("hello")

	// THEN the rest of members listen it
	assert.Equal(t, "hello", todoListenBeforeTimeout(t, m2.Listen(), time.Second))
	assert.Equal(t, "hello", todoListenBeforeTimeout(t, m3.Listen(), time.Second))
	// but not the actual sender
	todoNotReceivingBefore(t, m1.Listen(), 20*time.Millisecond)

	// AND WHEN some members have nothing to say
	m1.NothingToSay()
	m2.NothingToSay()

	// THEN the huddle does not end and all the listen channels are still open
	todoChannelOpen(t, m1.Listen())
	todoChannelOpen(t, m2.Listen())
	todoChannelOpen(t, m3.Listen())

	// AND WHEN all the members finish
	m3.NothingToSay()

	// THEN the huddle ends
	todoChannelClosed(t, m1.Listen())
	todoChannelClosed(t, m2.Listen())
	todoChannelClosed(t, m3.Listen())
}

// TODO: to be implemented as an actual test function
func todoListenBeforeTimeout[T any](t *testing.T, ch <-chan T, timeout time.Duration) T {

}

func todoNotReceivingBefore[T any](t *testing.T, ch <-chan T, timeout time.Duration) {

}

// warn: it might alter the status of the channel
func todoChannelOpen[T any](t *testing.T, ch <-chan T) T {
	select {
	case o, ok := <-ch:
		if !ok {
			assert.Fail(t, "expecting channel to be open but it is closed")
		}
		return o
	default:
		var o T
		return o
	}
}

func todoChannelClosed[T any](t *testing.T, ch <-chan T) T {
	select {
	case o, ok := <-ch:
		if ok {
			require.Fail(t, "expecting channel to be closed but it remains open")
		}
		return o
	default:
		require.Fail(t, "expecting channel to be closed but it remains open")
		var o T
		return o
	}
}
