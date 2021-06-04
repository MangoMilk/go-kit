package time

import (
	"strconv"
	"strings"
	"time"
)

type Month string

const (
	FormatDashDatetime = "2006-01-02 15:04:05"
	FormatDashDate     = "2006-01-02"
	FormatSeamlessDate = "20060102"

	JANUARY   = Month("1")
	FEBRUARY  = Month("2")
	MARCH     = Month("3")
	APRIL     = Month("4")
	MAY       = Month("5")
	JUNE      = Month("6")
	JULY      = Month("7")
	AUGUST    = Month("8")
	SEPTEMBER = Month("9")
	OCTOBER   = Month("10")
	NOVEMBER  = Month("11")
	DECEMBER  = Month("12")

	//JAN = "1"
	//FEB = "2"
	//MAR = "3"
	//APR = "4"
	//MAY  = "5"
	//JUN  = "6"
	//JUL  = "7"
	//AUG  = "8"
	//SEPT = "9"
	//OCT  = "10"
	//NOV  = "11"
	//DEC  = "12"
)

/*
 * IntervalDay
 * @param int64 bt: big timestamp
 * @param int64 lt: little timestamp
 */
func IntervalDay(bt int64, lt int64) int {

	var btDate string = time.Unix(bt, 0).Format(FormatDashDate)
	btLoc, _ := time.LoadLocation("Local")
	btTmp, _ := time.ParseInLocation(FormatDashDate, btDate, btLoc)
	var btDateTimestamp int64 = btTmp.Unix()

	var ltDate string = time.Unix(lt, 0).Format(FormatDashDate)
	ltLoc, _ := time.LoadLocation("Local")
	ltTmp, _ := time.ParseInLocation(FormatDashDate, ltDate, ltLoc)
	var ltDateTimestamp int64 = ltTmp.Unix()

	diffTime := btDateTimestamp - ltDateTimestamp

	var diffDay int = int(diffTime / (24 * 3600))

	return diffDay
}

func GetFormatDashDate(t time.Time) string {
	return t.Format(FormatDashDate)
}

func GetFormatSeamlessDate() string {
	return time.Now().Format(FormatSeamlessDate)
}

func GetCurDayPassS() int {
	t := time.Now()
	return int(time.Now().Unix() - time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())
}

func GetCurDayPassMS() int64 {
	t := time.Now()
	return (time.Now().UnixNano() - time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).UnixNano()) / 1e6
}

func GetTSByFormatDate(formatDate string) int64 {

	loc, _ := time.LoadLocation("Local") // get time zone
	timeStruct, _ := time.ParseInLocation(FormatDashDatetime, formatDate, loc)
	ts := timeStruct.Unix()

	return ts
}

// TODO: need to support in different year
func GetDateRange(startDate string, endDate string) (dateRange []interface{}) {
	s := strings.Split(startDate, "-")
	e := strings.Split(endDate, "-")

	endYear, _ := strconv.Atoi(e[0])
	startYear, _ := strconv.Atoi(s[0])

	if endYear < startYear {
		return
	}

	endMonth, _ := strconv.Atoi(e[1])
	startMonth, _ := strconv.Atoi(s[1])

	if endMonth < startMonth {
		return
	}

	endDay, _ := strconv.Atoi(e[2])
	startDay, _ := strconv.Atoi(s[2])

	if endMonth == startMonth && endDay < startDay {
		return
	}

	var startP int
	var endP int
	for i := startMonth; i <= endMonth; i++ {
		curMonth := strconv.Itoa(i)
		dayInMonth := GetNumsOfDayInMonth(Month(curMonth))

		if i == startMonth {
			startP = startDay
		} else {
			startP = 1
		}

		if i == endMonth {
			endP = endDay
		} else {
			endP = dayInMonth
		}

		for j := startP; j <= endP; j++ {
			if len(curMonth) < 2 {
				curMonth = "0" + curMonth
			}

			curDay := strconv.Itoa(j)

			if len(curDay) < 2 {
				curDay = "0" + curDay
			}

			dateRange = append(dateRange, strconv.Itoa(startYear)+"-"+curMonth+"-"+curDay)
		}
	}

	return
}

func GetNumsOfDayInMonth(month Month) int {
	switch month {
	case JANUARY:
		return 31
	case FEBRUARY:
		// TODO: need to cal real math of day in feb
		return 28
	case MARCH:
		return 31
	case APRIL:
		return 30
	case MAY:
		return 31
	case JUNE:
		return 30
	case JULY:
		return 31
	case AUGUST:
		return 30
	case SEPTEMBER:
		return 30
	case OCTOBER:
		return 31
	case NOVEMBER:
		return 30
	case DECEMBER:
		return 31
	default:
		return 0
	}
}
