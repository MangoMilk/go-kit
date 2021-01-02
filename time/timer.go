package time

import (
	"time"
)

const (
	DASH_DATE_FORMAT     = "2006-01-02"
	SEAMLESS_DATE_FORMAT = "20060102"
)

/*
 * IntervalDay
 * @param int64 bt: big timestamp
 * @param int64 lt: little timestamp
 */
func IntervalDay(bt int64, lt int64) int {

	var btDate string = time.Unix(bt, 0).Format(DASH_DATE_FORMAT)
	btLoc, _ := time.LoadLocation("Local")
	btTmp, _ := time.ParseInLocation(DASH_DATE_FORMAT, btDate, btLoc)
	var btDateTimestamp int64 = btTmp.Unix()

	var ltDate string = time.Unix(lt, 0).Format(DASH_DATE_FORMAT)
	ltLoc, _ := time.LoadLocation("Local")
	ltTmp, _ := time.ParseInLocation(DASH_DATE_FORMAT, ltDate, ltLoc)
	var ltDateTimestamp int64 = ltTmp.Unix()

	diffTime := btDateTimestamp - ltDateTimestamp

	var diffDay int = int(diffTime / (24 * 3600))

	return diffDay
}

func GetDashDateFormat(t time.Time) string {
	return t.Format(DASH_DATE_FORMAT)
}

func GetSeamlessDateFormat(t time.Time) string {
	return t.Format(SEAMLESS_DATE_FORMAT)
}

func GetCurDayPassS() int {
	t := time.Now()
	return int(time.Now().Unix() - time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())
}

func GetCurDayPassMS() int64 {
	t := time.Now()
	return (time.Now().UnixNano() - time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixNano()) / 1e6
}
