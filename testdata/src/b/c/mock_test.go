package c

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

// mocktail:Pineapple
// mocktail:Coconut

func TestName(t *testing.T) {
	var p Pineapple = newPineappleMock(t).
		OnHello(Water{}).TypedReturns("a").Once().
		OnWorld().TypedReturns("a").Once().
		OnGoo().TypedReturns("", 1, Water{}).Once().
		OnCoo("", Water{}).TypedReturns(Water{}).
		TypedRun(func(s string, water Water) {}).Once().
		Parent

	p.Hello(Water{})
	p.World()
	p.Goo()
	p.Coo(context.Background(), "", Water{})

	fn := func(st Strawberry, stban Strawberry) Pineapple {
		return p
	}

	fnMatcher := mock.MatchedBy(func(fn func(i int) int) bool {
		return fn(42) == 43
	})

	var c Coconut = newCoconutMock(t).
		OnLoo("a", 1, 2).TypedReturns("foo").Once().
		OnMoo(fn).TypedReturns("").Once().
		OnSooRaw(fnMatcher).TypedReturns(43).Once().
		Parent

	c.Loo("a", 1, 2)
	c.Moo(fn)
	c.Soo(func(i int) int { return i + 1 })
}
