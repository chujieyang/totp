/**
 * @author [kevinyang]
 * @email [yangchujie6@mail.com]
 * @create date 2020-11-10 10:42
 * @modify date 2020-11-10 10:42
 * @desc [description]
 */
package totp

import (
	"fmt"
	"io/ioutil"
)

func ScanAccountQrImgs() (data []string) {
	qr_dir := "./qr_imgs"
	files, _ := ioutil.ReadDir(qr_dir)
	for _, f := range files {
		data = append(data, fmt.Sprintf("%s/%s", qr_dir, f.Name()))
	}
	return
}
