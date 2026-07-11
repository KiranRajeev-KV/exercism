// will tell if a year is leap  year
package leap

// IsLeapYear finds leap years
func IsLeapYear(year int) bool {
    if (year%4 == 0 && (year%100 != 0 || year%400 == 0)) {
        return true
    }
    return false
}
