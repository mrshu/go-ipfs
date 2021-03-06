package perf

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ipfs/go-ipfs/Godeps/_workspace/src/github.com/bradfitz/iter"
	"gx/ipfs/QmZwjfAKWe7vWZ8f48u7AGA1xYfzR1iCD9A2XSCYFRBWot/testify/assert"
)

func TestTimer(t *testing.T) {
	tr := NewTimer()
	tr.Stop("hiyo")
	tr.Stop("hiyo")
	t.Log(em.Get("hiyo").(*buckets))
}

func BenchmarkStopWarm(b *testing.B) {
	tr := NewTimer()
	for range iter.N(b.N) {
		tr.Stop("a")
	}
}

func BenchmarkStopCold(b *testing.B) {
	tr := NewTimer()
	for i := range iter.N(b.N) {
		tr.Stop(strconv.FormatInt(int64(i), 10))
	}
}

func TestExponent(t *testing.T) {
	for _, c := range []struct {
		e int
		d time.Duration
	}{
		{-1, 10 * time.Millisecond},
		{-2, 5 * time.Millisecond},
		{-2, time.Millisecond},
		{-3, 500 * time.Microsecond},
		{-3, 100 * time.Microsecond},
	} {
		tr := NewTimer()
		time.Sleep(c.d)
		assert.Equal(t, c.e, bucketExponent(tr.Stop(fmt.Sprintf("%d", c.e))), "%s", c.d)
	}
	assert.Equal(t, `{"-1": 1}`, em.Get("-1").String())
	assert.Equal(t, `{"-2": 2}`, em.Get("-2").String())
	assert.Equal(t, `{"-3": 2}`, em.Get("-3").String())
}
