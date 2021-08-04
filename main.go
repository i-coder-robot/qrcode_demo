package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
	"image/png"
	"os"
)

func GenerateQrCode(file_name, content string, w, h int) {
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		fmt.Println(err)
	}
	qrCode, err = barcode.Scale(qrCode, w, h)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	png.Encode(file, qrCode)
}

func DecodeQrCode(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	decode, err := qrcode.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(decode.Content)
}

func main() {
	file_name := "title.png"
	GenerateQrCode(file_name, "面向加薪学习", 256, 256)
	DecodeQrCode(file_name)
}
