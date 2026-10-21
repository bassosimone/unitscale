package unitscale_test

import (
	"fmt"

	"github.com/bassosimone/unitscale"
)

func ExampleIEC_bytes() {
	fmt.Println(unitscale.IEC(500, "B"))
	// Output: 500 B
}

func ExampleIEC_kibibytes() {
	fmt.Println(unitscale.IEC(2048, "B"))
	// Output: 2.0 KiB
}

func ExampleIEC_mebibytes() {
	fmt.Println(unitscale.IEC(3*1024*1024, "B"))
	// Output: 3.0 MiB
}

func ExampleIEC_gibibytes() {
	fmt.Println(unitscale.IEC(1.5*1024*1024*1024, "B"))
	// Output: 1.5 GiB
}

func ExampleIEC_tebibytes() {
	fmt.Println(unitscale.IEC(2.5*1024*1024*1024*1024, "B"))
	// Output: 2.5 TiB
}

func ExampleSI_base() {
	fmt.Println(unitscale.SI(42, "bit/s"))
	// Output: 42 bit/s
}

func ExampleSI_kilo() {
	fmt.Println(unitscale.SI(1500, "bit/s"))
	// Output: 1.5 kbit/s
}

func ExampleSI_mega() {
	fmt.Println(unitscale.SI(10e6, "bit/s"))
	// Output: 10.0 Mbit/s
}

func ExampleSI_giga() {
	fmt.Println(unitscale.SI(2.4e9, "bit/s"))
	// Output: 2.4 Gbit/s
}

func ExampleSI_tera() {
	fmt.Println(unitscale.SI(3.7e12, "bit/s"))
	// Output: 3.7 Tbit/s
}
