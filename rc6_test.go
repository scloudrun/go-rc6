/***************************************************************************
 *
 * Copyright (c) 2019 github.com, Inc. All Rights Reserved
 * rc6 encrypt test
 * Author scloudrun
 *
**************************************************************************/

package rc6lib

import (
	"reflect"
	"testing"
)

var (
	defaultKey     = []byte("0000000000000000")
	defaultIv      = []byte("1111111111111111")
	defaultEncData = []byte("scloudrun")
	defaultEncByte = []byte{148, 250, 198, 116, 20, 46, 103, 128, 40, 160, 83, 239, 210, 246, 151, 15}
)

func Test_Rc6Enc(t *testing.T) {
	type args struct {
		key           []byte
		iv            []byte
		plantText     []byte
		paddingStatus bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"Test_Rc6Enc",
			args{defaultKey, defaultIv, defaultEncData, true},
			defaultEncByte,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rc6Enc(tt.args.key, tt.args.iv, tt.args.plantText, tt.args.paddingStatus)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rc6Enc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rc6Enc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Rc6Dec(t *testing.T) {
	type args struct {
		key           []byte
		iv            []byte
		ciphertext    []byte
		paddingStatus bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"Test_rc6dec",
			args{defaultKey, defaultIv, defaultEncByte, true},
			[]byte(defaultEncData),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rc6Dec(tt.args.key, tt.args.iv, tt.args.ciphertext, tt.args.paddingStatus)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rc6Dec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rc6Dec() = %v, want %v", got, tt.want)
			}
		})
	}
}
