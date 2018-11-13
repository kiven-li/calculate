package calculate

//加法
func add(a, b int) int {
	var sum int = a
	var carry int = b

	for carry != 0 {
		temp := sum
		sum = temp ^ carry
		carry = (temp & carry) << 1
	}

	return sum
}

//减法
func subtract(a, b int) int {
	var substrahend int = add(^b, 1)
	var sub int = add(a, substrahend)

	return sub
}

//用加法实现的乘法
func multiply(a, b int) int {
	//将乘数和被乘数都取绝对值
	var multiplier int = a
	if a < 0 {
		multiplier = add(^a, 1)
	}

	var multiplicand int = b
	if b < 0 {
		multiplicand = add(^b, 1)
	}

	//计算绝对值的乘积
	var product int = 0
	var count int = 0

	for count < multiplier {
		product = add(product, multiplicand)
		count = add(count, 1)
	}

	//计算乘积的符号
	if (a ^ b) < 0 {
		product = add(^product, 1)
	}

	return product
}

//改进的乘法
func multiplyEx(a, b int) int {
	//将乘数和被乘数都取绝对值
	var multiplier int = a
	if a < 0 {
		multiplier = add(^a, 1)
	}

	var multiplicand int = b
	if b < 0 {
		multiplicand = add(^b, 1)
	}

	//计算绝对值的乘积
	var product int = 0
	for multiplier != 0 {
		if multiplier&0x1 != 0 {
			product = add(product, multiplicand)
		}

		multiplicand = multiplicand << 1
		multiplier = multiplier >> 1
	}

	//计算乘积的符号
	if (a ^ b) < 0 {
		product = add(^product, 1)
	}

	return product
}

//减法实现的求商和余
func divide(a, b int) (int, int) {
	//对被除数和除数取绝对值
	var dividend int = a
	if a < 0 {
		dividend = add(^a, 1)
	}

	var divisor int = b
	if b < 0 {
		divisor = add(^b, 1)
	}

	//对被除数和除数的绝对值求商
	var remainder int = dividend
	var quotient int = 0

	for remainder >= divisor {
		remainder = subtract(remainder, divisor)
		quotient = add(quotient, 1)
	}

	//求商的符号
	if (a ^ b) < 0 {
		quotient = add(^quotient, 1)
	}

	//求余的符号
	if a < 0 {
		remainder = add(^remainder, 1)
	}

	return quotient, remainder
}

//改进的求商和余
func divideEx(a, b int) (int, int) {
	//对被除数和除数取绝对值
	var dividend int = a
	if a < 0 {
		dividend = add(^a, 1)
	}

	var divisor int = b
	if b < 0 {
		divisor = add(^b, 1)
	}

	//获得被除数的反序 比如dividend=101011 invert=1110101,invert最高位会多一个1,
	//这是为了防止dividend=1010,则直接反转为0101,这个时候原来的最低位的0就会丢失
	var invert int = 2
	for dividend != 0 {
		invert = invert | (dividend & 0x1)
		invert = invert << 1
		dividend = dividend >> 1
	}

	var quotient int = 0
	var remainder int = 0
	for invert > 1 { //排除最高位的1
		remainder = remainder << 1
		remainder = remainder | (invert & 0x1)
		invert = invert >> 1
		quotient = quotient << 1

		if remainder >= divisor {
			quotient = quotient | 0x1
			remainder = subtract(remainder, divisor)
		}
	}

	//求商的符号
	if (a ^ b) < 0 {
		quotient = add(^quotient, 1)
	}

	//求余的符号
	if a < 0 {
		remainder = add(^remainder, 1)
	}

	return quotient, remainder
}
