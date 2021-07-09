package generator

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
 * Gen Open ID
 */
func GenOpenId() string {
	// static $guid = '';
	// 		$uid  = uniqid("", true);
	// var data string
	// 		$data .= $_SERVER['REQUEST_TIME'];
	// data = $_SERVER['HTTP_USER_AGENT'];
	// data += $_SERVER['SERVER_ADDR'];
	// 		$data .= $_SERVER['SERVER_PORT'];
	// 		$data .= $_SERVER['REMOTE_ADDR'];
	// 		$data .= $_SERVER['REMOTE_PORT'];
	// 		$hash = strtoupper(hash('ripemd128', $uid . $guid . md5($data)));
	// 		$guid =
	// 			substr($hash, 0, 8) .
	// 			substr($hash, 8, 4) .
	// 			substr($hash, 12, 4) .
	// 			substr($hash, 16, 4) .
	// 			substr($hash, 20, 12);

	// 		return $guid;
	var err error
	var u string = uuid.Must(uuid.NewV4(), err).String()
	u = strings.ToUpper(strings.Replace(u, "-", "", -1))

	return u
}

func GenUUIDv1() (id string, err error) {
	id = uuid.Must(uuid.NewV1(), err).String()
	return
}

func GenUUIDv4() (id string, err error) {
	id = uuid.Must(uuid.NewV4(), err).String()
	return
}

/*
 * Gen Serial No
 */
func GenSerialNo(prefix string, suffix string) string {
	rand.Seed(time.Now().UnixNano())
	return prefix + strconv.Itoa(100+rand.Intn(999)) + strconv.Itoa(int(time.Now().Unix())) + suffix
}

//var numberPool = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
//var upperAlphabetPool = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
//var lowerAlphabetPool = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

type StringElem int

const (
	ElemLowerCharAndNum = StringElem(1)
	ElemUpperCharAndNum = StringElem(2)
	ElemAllChar         = StringElem(3)
)

/*
 * Gen Rand String
 */
func GenRandString(length int, elem StringElem) (code string) {
	var pool []string
	switch elem {
	case ElemLowerCharAndNum:
		pool = []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		}
		break
	case ElemUpperCharAndNum:
		pool = []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		}
		break
	case ElemAllChar:
		pool = []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		}
		break
	default:
		pool = []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		}
	}

	poolSize := len(pool)

	for i := 1; i <= length; i++ {
		rand.Seed(time.Now().UnixNano())
		code += pool[rand.Intn(poolSize)]
	}

	return code
}
