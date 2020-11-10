/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"totp/cmd/totp"

	"github.com/spf13/cobra"
)

var account string

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "查看指定账号的动态密码",
	Long:  "查看指定账号的动态密码",
	Run: func(cmd *cobra.Command, args []string) {
		hasAccount := false
		qrList := totp.ScanAccountQrImgs()
		for _, qr := range qrList {
			qrContent := totp.ReadQRCode(qr)
			accountName, secret := totp.GetQrSecret(qrContent)
			if account == accountName {
				hasAccount = true
				code, refresh := totp.GetTOTPToken(secret)
				fmt.Printf("\033[1;34m Code: 「%s」, 距离下次更新还有 %d 秒. \033[0m\n", code, refresh)
				break
			}
		}
		if !hasAccount {
			fmt.Printf("\033[1;34m %s \033[0m\n", "当前 qr_imgs 目录下未找到对应的账户信息.")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.
	codeCmd.Flags().StringVarP(&account, "account", "a", "", "需要查看动态密码的账户名")

}
