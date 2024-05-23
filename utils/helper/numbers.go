package helper

const maxDigit = 18

var Pow10 = func() (ret []int64) {
	cur := int64(1)
	ret = append(ret, cur)

	for i := 1; i <= maxDigit; i++ {
		cur *= 10
		ret = append(ret, cur)
	}
	return
}()

func IsBetween[T int | int64](val, lo, hi T) bool {
	return lo <= val && val <= hi
}

func HasLen(val int64, len int) bool {
	return IsBetween(val, Pow10[len-1], Pow10[len]-1)
}

func GetLen(val int64) int {
	for i := 0; i < maxDigit; i++ {
		if val < Pow10[i] {
			return i
		}
	}
	return maxDigit
}

// 1-based indexing
func GetSubDigit(val int64, len, leftDigit, rightDigit int) int64 {
	targetLen := rightDigit - leftDigit + 1
	return val / Pow10[len-rightDigit] % Pow10[targetLen]
}
