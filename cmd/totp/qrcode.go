package totp

import (
	"fmt"
	"os"
	"regexp"

	qrcodeReader "github.com/tuotoo/qrcode"
)

func ReadQRCode(filename string) (content string) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcodeReader.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	content = qrmatrix.Content
	return
}

func GetQrSecret(content string) (account string, secret string) {
	flysnowRegexp := regexp.MustCompile(`^otpauth://totp/(.*)\?secret=(.*)\&issuer=(.*)$`)
	//otpauth://totp/Aliyun:123456@qq.com?secret=Y4RHCF4JNK36EN6M2Y4UOIEVN5FJOLPJB&issuer=Aliyun
	params := flysnowRegexp.FindStringSubmatch(content)
	if len(params) == 4 {
		account = params[1]
		secret = params[2]
	}
	return
}
