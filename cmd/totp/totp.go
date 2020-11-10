/**
 * @author [kevinyang]
 * @email [yangchujie6@mail.com]
 * @create date 2020-11-10 10:28
 * @modify date 2020-11-10 10:28
 * @desc [description]
 */
package totp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func GetTOTPToken(secret string) (otp string, refresh int) {
	now := time.Now()
	interval := now.Unix() / 30
	refresh = int(math.Abs(float64(30 - now.Second())))
	otp, err := getHOTPToken(secret, interval)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	return
}

func getHOTPToken(secret string, interval int64) (otp string, err error) {
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	if err != nil {
		return
	}
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(interval))
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)
	o := (h[19] & 15)
	var header uint32
	r := bytes.NewReader(h[o : o+4])
	if err = binary.Read(r, binary.BigEndian, &header); err != nil {
		return
	}
	h12 := (int(header) & 0x7fffffff) % 1000000
	otp = strconv.Itoa(int(h12))
	return
}
