package integer_to_roman

import "fmt"

func intToRoman(num int) string {
	MList := []string{"", "M", "MM", "MMM"}
	CList := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	XList := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	IList := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return fmt.Sprintf("%s%s%s%s", MList[num/1000], CList[(num/100)%10], XList[(num/10)%10], IList[num%10])
}
