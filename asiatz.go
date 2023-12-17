package asiatz

import (
	"fmt"
	"strconv"
)

// ToUTC converts a time string (HH:mm) from a specified timezone to UTC time string (HH:mm).
func ToUTC(timezoneOffset float64, time string) (string, error) {
	hour, err := strconv.Atoi(time[:2])
	if err != nil {
		return "", err
	}
	if hour < 0 || hour > 23 {
		return "", fmt.Errorf("hour must be between 0 and 23")
	}
	minute, err := strconv.Atoi(time[3:])
	if err != nil {
		return "", err
	}
	if minute < 0 || minute > 59 {
		return "", fmt.Errorf("minute must be between 0 and 59")
	}
	totalMinutes := hour*60 + minute
	utcTotalMinutes := ((totalMinutes-int(timezoneOffset*60))%1440 + 1440) % 1440
	utcHour := utcTotalMinutes / 60
	utcMinute := utcTotalMinutes % 60
	utcTime := fmt.Sprintf("%02d:%02d", utcHour, utcMinute)
	return utcTime, nil
}

// ShanghaiToUTC converts a Shanghai time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Shanghai is equivalent to "00:00" in UTC.
func ShanghaiToUTC(shanghaiTime string) (string, error) {
	return ToUTC(8, shanghaiTime)
}

// TokyoToUTC converts a Tokyo time string (HH:mm) to UTC time string (HH:mm).
// For example, "09:00" in Tokyo is equivalent to "00:00" in UTC.
func TokyoToUTC(tokyoTime string) (string, error) {
	return ToUTC(9, tokyoTime)
}

// HongKongToUTC converts a Hong Kong time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Hong Kong is equivalent to "00:00" in UTC.
func HongKongToUTC(hongKongTime string) (string, error) {
	return ToUTC(8, hongKongTime)
}

// SingaporeToUTC converts a Singapore time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Singapore is equivalent to "00:00" in UTC.
func SingaporeToUTC(singaporeTime string) (string, error) {
	return ToUTC(8, singaporeTime)
}

// TaipeiToUTC converts a Taipei time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Taipei is equivalent to "00:00" in UTC.
func TaipeiToUTC(taipeiTime string) (string, error) {
	return ToUTC(8, taipeiTime)
}

// SeoulToUTC converts a Seoul time string (HH:mm) to UTC time string (HH:mm).
// For example, "09:00" in Seoul is equivalent to "00:00" in UTC.
func SeoulToUTC(seoulTime string) (string, error) {
	return ToUTC(9, seoulTime)
}

// BeijingToUTC converts a Beijing time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Beijing is equivalent to "00:00" in UTC.
func BeijingToUTC(beijingTime string) (string, error) {
	return ToUTC(8, beijingTime)
}

// DubaiToUTC converts a Dubai time string (HH:mm) to UTC time string (HH:mm).
// For example, "04:00" in Dubai is equivalent to "00:00" in UTC.
func DubaiToUTC(dubaiTime string) (string, error) {
	return ToUTC(4, dubaiTime)
}

// DelhiToUTC converts a Delhi time string (HH:mm) to UTC time string (HH:mm).
// For example, "05:30" in Delhi is equivalent to "00:00" in UTC.
func DelhiToUTC(delhiTime string) (string, error) {
	return ToUTC(5.5, delhiTime)
}

// JakartaToUTC converts a Jakarta time string (HH:mm) to UTC time string (HH:mm).
// For example, "07:00" in Jakarta is equivalent to "00:00" in UTC.
func JakartaToUTC(jakartaTime string) (string, error) {
	return ToUTC(7, jakartaTime)
}

// BangkokToUTC converts a Bangkok time string (HH:mm) to UTC time string (HH:mm).
// For example, "07:00" in Bangkok is equivalent to "00:00" in UTC.
func BangkokToUTC(bangkokTime string) (string, error) {
	return ToUTC(7, bangkokTime)
}

// PyongyangToUTC converts a Pyongyang time string (HH:mm) to UTC time string (HH:mm).
// For example, "09:00" in Pyongyang is equivalent to "00:00" in UTC.
func PyongyangToUTC(pyongyangTime string) (string, error) {
	return ToUTC(9, pyongyangTime)
}
