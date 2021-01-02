package id

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
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

/*
 * Gen Withdraw No
 */
func GenWithdrawNo(prefix string, suffix string) string {
	rand.Seed(time.Now().UnixNano())
	return prefix + strconv.Itoa(100+rand.Intn(999)) + strconv.Itoa(int(time.Now().Unix())) + suffix
}


/*
 * Gen 6 char code: 2 number 4 alphabet
 */
func Gen6CharCode() string {
	var numberPool = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var alphabetPool = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var pool []string = append(numberPool, alphabetPool...)
	var code string = ""

	rand.Seed(time.Now().Unix())

	for i := 1; i <= 6; i++ {
		index := rand.Intn(len(pool))
		code += pool[index]
	}

	return code
}
