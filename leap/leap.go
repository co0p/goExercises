package leap


// Leap Years are any year that can be evenly divided by 4 (such as 2012, 2016, etc)
func canBeDividedBy4(year int) bool {
  return year % 4 == 0
}

// except if it can be evenly divided by 100, then it isn't (such as 2100, 2200, etc)
func canBeDividedBy100(year int) bool {
    return year % 100 == 0
}

// except if it can be evenly divided by 400, then it is (such as 2000, 2400)
func canBeDividedBy400(year int) bool {
    return year % 400 == 0
}

func IsLeapYear(year int) bool  {
  if !canBeDividedBy4(year) {
      return false
  }

  if canBeDividedBy100(year) {
    if (canBeDividedBy400(year)) {
      return true
    }
    return false
  }

  return true
}
