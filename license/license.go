package license

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/xumuk666/lk"
	"log"
	"os"
	"time"
)

// the public key b32 encoded from the private key using: lkgen pub my_private_key_file`.
// It should be hardcoded somewhere in your app.
const publicKeyBase32 = "ATQYAQR2U5XYBPABZ3Q6HBZXSUFNMRVKQJFDUKJZLYXACNCMPAT62ETQ2N6LUOVQRZYSWFXBKMF6GVURDTIXSZDPWPXETAOU7ZJRTGR6OTPXOD4PQND5HVKTGGFNKIHWMJDIWWBLISUJRQJGJCOAJYYKJK7Q===="

func Verify() bool {

	file, err := os.Open("license.key")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	licenseB32, e := Readln(r)
	if e != nil {
		return false
	}



	// Unmarshal the public key.
	publicKey, err := lk.PublicKeyFromB32String(publicKeyBase32)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Unmarshal the customer license.

	license, err := lk.LicenseFromB32String(licenseB32)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// validate the license signature.
	if ok, err := license.Verify(publicKey); err != nil {
		log.Fatal(err)
		return false
	} else if !ok {
		log.Fatal("Invalid license signature")
		return false
	}

	result := struct {
		Email string    `json:"email"`
		End   time.Time `json:"end"`
	}{}

	// unmarshal the document.
	if err := json.Unmarshal(license.Data, &result); err != nil {
		log.Fatal(err)
	}

	// Now you just have to check that the end date is after time.Now() then you can continue!
	if result.End.Before(time.Now()) {
		log.Fatal("License expired on: %s", result.End.Format("2006-01-02"))
	} else {
		fmt.Printf(`Licensed by %s until %s`, result.Email, result.End.Format("2006-01-02"))
		return true
	}
	return false
}


func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}