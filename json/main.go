package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var t = `[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]`

	data, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println(string(data))

	sum := sha256.Sum256([]byte("X"))
	log.Println(fmt.Sprintf("%x", sum))
}

// GetMachineID gets the unique ID of current machine.
func GetMachineID() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue("MachineGuid")
	if err != nil {
		return "", err
	}
	return s, nil
}
