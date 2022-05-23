package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventually_Error(t *testing.T) {
	innerTest := &testing.T{}
	Eventually(innerTest, 10*time.Millisecond, func(t require.TestingT) {
		require.True(t, false)
	})
	assert.True(t, innerTest.Failed())
}

func TestEventually_Fail(t *testing.T) {
	innerTest := &testing.T{}
	Eventually(innerTest, 10*time.Millisecond, func(t require.TestingT) {
		t.FailNow()
	})
	assert.True(t, innerTest.Failed())
}

func TestEventually_Timeout(t *testing.T) {
	innerTest := &testing.T{}
	Eventually(innerTest, 10*time.Millisecond, func(t require.TestingT) {
		time.Sleep(5 * time.Second)
	})
	assert.True(t, innerTest.Failed())
}

func TestEventually_Success(t *testing.T) {
	num := 3
	Eventually(t, 5*time.Second, func(t require.TestingT) {
		require.Equal(t, 0, num)
		num--
	})
}

func TestEventually_Interval(t *testing.T) {
	innerTest := &testing.T{}
	executions := 0
	Eventually(innerTest, 10*time.Millisecond, func(t require.TestingT) {
		executions++
		t.FailNow()
	}, Interval(20*time.Second))
	assert.True(t, innerTest.Failed())
	assert.Equal(t, 1, executions)
}
