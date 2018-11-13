package calculate

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		left  int
		right int
		sum   int
	}{
		{1, 2, 3},
		{101, 97, 198},
		{1999, 1, 2000},
		{2020202, 1010101, 3030303},
	}

	for _, v := range tests {
		if add(v.left, v.right) != v.sum {
			t.Fatalf("%d + %d = %d turn out:%d", v.left, v.right, v.sum, add(v.left, v.right))
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		left  int
		right int
		sub   int
	}{
		{1, 2, -1},
		{101, 97, 4},
		{1999, 1, 1998},
		{2020202, 1010101, 1010101},
		{232, 53434, -53202},
	}

	for _, v := range tests {
		if subtract(v.left, v.right) != v.sub {
			t.Fatalf("%d - %d = %d turn out:%d", v.left, v.right, v.sub, subtract(v.left, v.right))
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		left    int
		right   int
		product int
	}{
		{1, 2, 2},
		{101, 97, 9797},
		{1999, -100, -199900},
		{-15, -15, 225},
		{87, 46, 4002},
	}

	for _, v := range tests {
		if multiply(v.left, v.right) != v.product {
			t.Fatalf("multiply %d - %d = %d turn out:%d", v.left, v.right, v.product, multiply(v.left, v.right))
		}
		if multiplyEx(v.left, v.right) != v.product {
			t.Fatalf("multiplyEx %d - %d = %d turn out:%d", v.left, v.right, v.product, multiplyEx(v.left, v.right))
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		left      int
		right     int
		quotient  int
		remainder int
	}{
		{10, 3, 3, 1},
		{-111, 33, -3, -12},
	}

	for _, v := range tests {
		{
			q, r := divide(v.left, v.right)
			if q != v.quotient {
				t.Fatalf("divide %d / %d = %d turn out:%d", v.left, v.right, v.quotient, q)
			}
			if r != v.remainder {
				t.Fatalf("divide %d MOD %d = %d turn out:%d", v.left, v.right, v.remainder, r)
			}
		}
		{
			q, r := divideEx(v.left, v.right)
			if q != v.quotient {
				t.Fatalf("divideEx %d / %d = %d turn out:%d", v.left, v.right, v.quotient, q)
			}
			if r != v.remainder {
				t.Fatalf("divideEx %d MOD %d = %d turn out:%d", v.left, v.right, v.remainder, r)
			}
		}
	}
}
