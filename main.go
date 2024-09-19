package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/liyue201/goqr"
	"golang.design/x/clipboard"
	"golang.org/x/term"
)

var dataTypes = map[int]string{
	1: "Numeric",
	2: "Alpha",
	4: "Byte",
	8: "Kanji",
}

var eccLevels = map[int]string{
	0: "M",
	1: "L",
	2: "H",
	3: "Q",
}

var eciEncodings = map[uint32]string{
	0:  "ISO-8859-1 (default)",
	1:  "ISO-8859-1",
	2:  "IBM437",
	4:  "ISO-8859-2",
	5:  "ISO-8859-3",
	6:  "ISO-8859-4",
	7:  "ISO-8859-5",
	8:  "ISO-8859-6",
	9:  "ISO-8859-7",
	10: "ISO-8859-8",
	11: "ISO-8859-9",
	13: "Windows874",
	15: "ISO-8859-13",
	17: "ISO-8859-15",
	20: "ShiftJIS",
	21: "UTF-8",
}

func main() {
	var err error
	var f io.Reader
	if len(os.Args) > 1 {
		if f, err = os.OpenFile(os.Args[1], os.O_RDONLY, 0); err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
	} else if !term.IsTerminal(int(os.Stdin.Fd())) {
		f = os.Stdin
	} else {
		err := clipboard.Init()
		if err != nil {
			fmt.Println("Error initializing clipboard:", err)
			os.Exit(1)
		}
		b := clipboard.Read(clipboard.FmtImage)
		if len(b) == 0 {
			fmt.Println("No qr found in stdin or clipboard")
			os.Exit(1)
		}
		f = bytes.NewReader(b)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		os.Exit(1)
	}

	codes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Println("Error recognizing QR code:", err)
		os.Exit(1)
	}

	for _, code := range codes {
		fmt.Println("Data:", string(code.Payload))
		fmt.Println("Version:", code.Version)
		fmt.Println("ECC:", eccLevels[code.EccLevel])
		fmt.Println("Mask:", code.Mask)
		fmt.Println("Data Type:", dataTypes[code.DataType])
		fmt.Println("ECI:", eciEncodings[code.Eci])
	}
}
