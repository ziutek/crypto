// xxtea reads text numbers from its stdin and outputs encrypted/decrypted
// text numbers to its stdout using KEYFILE. KEYFILE should contains four
// unsigned 32bits integers (in text form). xxtea reads all input before
// encrypt it so it isn't suitable for encrypt/decrypt long or infinite
// stream of numbers.
package main

import (
	"bufio"
	"github.com/ziutek/crypto/xxtea"
	"io"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err == nil {
		return
	}
	io.WriteString(os.Stderr, err.Error()+"\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "e" && os.Args[1] != "d" {
		io.WriteString(os.Stderr, "Usage: "+os.Args[0]+" {e|d} KEYFILE\n")
		os.Exit(1)
	}

	encrypt := (os.Args[1] == "e")
	keyfile := os.Args[2]

	// Read key

	f, err := os.Open(keyfile)
	checkErr(err)

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	var (
		key [4]uint32
		i   int
	)
	for i < len(key) && s.Scan() {
		u, err := strconv.ParseUint(s.Text(), 0, 32)
		checkErr(err)
		key[i] = uint32(u)
		i++
	}

	f.Close()
	checkErr(s.Err())
	if i != len(key) {
		io.WriteString(os.Stderr, "Key file doesn't contain enough numbers\n")
		os.Exit(1)
	}

	// Encrypt numbers from stdin

	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	var numbers []uint32
	for s.Scan() {
		u, err := strconv.ParseUint(s.Text(), 0, 32)
		checkErr(err)
		numbers = append(numbers, uint32(u))
	}
	checkErr(s.Err())

	if encrypt {
		xxtea.Encrypt(numbers, key)
	} else {
		xxtea.Decrypt(numbers, key)
	}

	for _, u := range numbers {
		_, err := io.WriteString(
			os.Stdout,
			"0x"+strconv.FormatUint(uint64(u), 16)+"\n",
		)
		checkErr(err)
	}
}
