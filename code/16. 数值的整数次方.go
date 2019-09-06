func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent&1 == 0 {
		t := Power(base, exponent/2)
		return t * t
	}
	if exponent < 0 {
		return 1 / base * Power(base, exponent+1)
	}
	return base * Power(base, exponent-1)
}