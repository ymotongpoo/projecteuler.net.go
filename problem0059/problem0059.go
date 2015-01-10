// Each character on a computer is assigned a unique code and the preferred standard is ASCII
// (American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.
//
// A modern encryption method is to take a text file, convert the bytes to ASCII,
// then XOR each byte with a given value, taken from a secret key. The advantage with the XOR function is that using
// the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.
//
// For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes.
// The user would keep the encrypted message and the encryption key in different locations, and without both "halves",
// it is impossible to decrypt the message.
//
// Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key.
// If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message.
// The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.
//
// Your task has been made easy, as the encryption key consists of three lower case characters.
// Using cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes,
// and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// CipherFile is assigned cipher
const CipherFile = "p059_cipher.txt"

func readCipher(reader io.Reader) ([]int, error) {
	var err error
	br := bufio.NewReader(reader)
	codes := []int{}
	for err == nil {
		token, err := br.ReadBytes(',')
		if err == io.EOF {
			break
		}
		i, err := strconv.ParseInt(string(token[:len(token)-1]), 0, 0)
		if err != nil {
			return nil, err
		}
		codes = append(codes, int(i))
	}
	return codes, nil
}

func decrypt(key []byte, crypted []byte) ([]byte, error) {
	if len(key) > 3 {
		return nil, fmt.Errorf("key is too long (must be 3 chars): current %v (%vchars)", key, len(key))
	}
	decrypted := make([]byte, len(crypted))
	i := 0
	for len(crypted) > 0 {
		buf := crypted[:3]
		buf[0] = crypted[0] ^ key[0]
		buf[1] = crypted[1] ^ key[1]
		buf[2] = crypted[2] ^ key[2]
		copy(decrypted[i:i+3], buf)
		i += 3
		crypted = crypted[3:]
	}
	return decrypted, nil
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(cwd, CipherFile)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	nums, err := readCipher(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
}
