package matchstr

import "regexp"

var (
	//Phone regular expression
	PhoneMathers = regexp.MustCompile("(^\\+[0-9]{2}|^\\+[0-9]{2}\\(0\\)|^\\(\\+[0-9]{2}\\)\\(0\\)|^00[0-9]{2}|^0)([0-9]{9}$|[0-9\\-\\s]{10}$)")
	//Barcode regular expression
	Barcode=regexp.MustCompile("^123456[0-9]{8}$")
	// Email  regular expression
	EmailMathers= regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
// Check Email Address is valid
func EmailMatcher(email string) (bool)  {
	isValid:=EmailMathers.MatchString(email)
	if isValid {
		  return true;
	}
	return false

}
//Check phone number is valid
func PhoneMatcher(phone string) (bool)  {
	isValid:=PhoneMathers.MatchString(phone)
	if isValid {
		return true;
	}
	return false

}
//Check phone number is valid
func BarcodeMatcher(barcode string) (bool)  {
	isValid:=Barcode.MatchString(barcode)
	if isValid {
		return true;
	}
	return false

}