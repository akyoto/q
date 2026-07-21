main() {
	a := 1 + 1
	b := 2 + 2
	c := 3 + 3
	d := 4 + 4
	e := 5 + 5
	f := 6 + 6
	g := 7 + 7
	h := 8 + 8
	i := 9 + 9
	j := 10 + 10
	k := 11 + 11
	l := 12 + 12
	m := 13 + 13
	n := 14 + 14
	o := 15 + 15
	p := 16 + 16
	q := 17 + 17
	r := 18 + 18
	s := 19 + 19
	t := 20 + 20
	u := 21 + 21
	v := 22 + 22
	w := 23 + 23
	x := 24 + 24
	y := 25 + 25
	z := 26 + 26
	total := sum(a, b, c)
	total += sum(d, e, f)
	total += sum(g, h, i)
	total += sum(j, k, l)
	total += sum(m, n, o)
	total += sum(p, q, r)
	total += sum(s, t, u)
	total += sum(v, w, x)
	total += sum(y, z, 0)
	assert total == 351 + 351
}

sum(a int, b int, c int) -> int {
	return a + b + c
}