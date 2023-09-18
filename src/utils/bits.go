package utils

// Bits is a utility class with methods to extract a sequence of bits from a 32-bit vector.

/*
ExtractSigned extracts, from a 32-bit vector, length bits starting at the index bit start
and returns the two's complement signed interpretation of the value.
*/
func ExtractSigned(value int32, start, length int) int32 {
	end := (start + length) - 1

	if start < 0 || end > 31 {
		panic("IllegalArgumentException: start and end must be in the interval [0, 31]")
	}

	if length < 0 {
		panic("IllegalArgumentException: length must be non-negative")
	}

	shiftLeft := 32 - (start + length)
	leftArithmeticShiftedBits := value << uint(shiftLeft)

	shiftRight := 32 - length
	rightArithmeticShiftedBits := leftArithmeticShiftedBits >> uint(shiftRight)

	return rightArithmeticShiftedBits
}

/*
ExtractUnsigned extracts, from a 32-bit vector, length bits starting at the index bit start
and returns the unsigned interpretation of the value.
*/
func ExtractUnsigned(value int32, start, length int) uint32 {
	end := (start + length) - 1

	if start < 0 || end > 31 {
		panic("IllegalArgumentException: start and end must be in the interval [0, 31]")
	}

	if length < 0 || length == 32 {
		panic("IllegalArgumentException: length must be non-negative and not equal to 32")
	}

	shiftLeft := 32 - (start + length)
	leftArithmeticShiftedBits := value << uint(shiftLeft)

	shiftRight := 32 - length
	rightLogicShiftedBits := uint32(leftArithmeticShiftedBits) >> uint(shiftRight)

	return rightLogicShiftedBits
}
