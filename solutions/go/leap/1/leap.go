
package leap

// IsLeapYear tells if it a leap year
func IsLeapYear(year int) bool {
    if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
        return true
    }
    return false
}
