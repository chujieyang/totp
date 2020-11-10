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

// accountsCmd represents the accounts command
var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "查看当前管理的账号 Secret 列表",
	Long:  `查看当前管理的账号 Secret 列表，数据来源当前目录 qr_imgs 目录下的二维码图片文件.`,
	Run: func(cmd *cobra.Command, args []string) {
		qrList := totp.ScanAccountQrImgs()
		fmt.Printf("\033[0;36m --------- 账号列表 --------- \033[0m\n")
		for index, qr := range qrList {
			qrContent := totp.ReadQRCode(qr)
			accountName, _ := totp.GetQrSecret(qrContent)
			fmt.Printf("\033[1;34mNo: 「%d」 Account: 「%s」\033[0m\n", index+1, accountName)
		}
	},
}

func init() {
	rootCmd.AddCommand(accountsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
