package asiatz

// ConvertShanghaiToUTC converts a Shanghai time string (HH:mm) to UTC time string (HH:mm).
// For example, "08:00" in Shanghai is equivalent to "00:00" in UTC.
func ConvertShanghaiToUTC(shanghaiTime string) (string, error) {
	shanghaiHour, err := strconv.Atoi(shanghaiTime[:2])
	if err != nil {
		return "", err
	}
	shanghaiMinute, err := strconv.Atoi(shanghaiTime[3:])
	if err != nil {
		return "", err
	}
	shanghaiTotalMinutes := shanghaiHour*60 + shanghaiMinute
	utcTotalMinutes := (shanghaiTotalMinutes - 480 + 1440) % 1440
	utcHour := utcTotalMinutes / 60
	utcMinute := utcTotalMinutes % 60
	utcTime := fmt.Sprintf("%02d:%02d", utcHour, utcMinute)
	return utcTime, nil
}
