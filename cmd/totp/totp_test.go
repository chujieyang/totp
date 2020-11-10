/**
 * @author [kevinyang]
 * @email [yangchujie6@mail.com]
 * @create date 2020-11-10 10:34
 * @modify date 2020-11-10 10:34
 * @desc [description]
 */
package totp

import (
	"fmt"
	"testing"
)

func TestGetTOTPToken(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test get totp code",
			args: args{
				secret: "Y4RHCF4JNK3QGZYCZZ53WLA47PJS66FXJ73OF7OW276EN6M2Y4UOIEVN5FJOLPJB",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if code, refresh := GetTOTPToken(tt.args.secret); code != tt.want {
				t.Errorf("getTOTPToken() = %v, %v, want %v", code, refresh, tt.want)
			} else {
				fmt.Printf("code: %s, refresh: %d\n", code, refresh)
			}
		})
	}
}
