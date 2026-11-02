package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}
	shift %= len(runes)
	if shift < 0 {
		shift += len(runes)
	}
	// Reversing each part and then the whole slice swaps the parts.
	reverse := func(part []rune) {
		for left, right := 0, len(part)-1; left < right; left, right = left+1, right-1 {
			part[left], part[right] = part[right], part[left]
		}
	}
	if shift != 0 {
		reverse(runes[:shift])
		reverse(runes[shift:])
		reverse(runes)
	}
	return string(runes)
}
