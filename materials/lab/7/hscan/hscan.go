package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup = make(map[string]string)
var md5lookup = make(map[string]string)
var shaM sync.Map
var md5M sync.Map

func GuessSingle(sourceHash string, filename string) { //this length chech isn't correct

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if len(password) == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
		} else {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func solveMD5(password string, wga *sync.WaitGroup) {
	log.Printf("Launched an MD5 routine")

	temp := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	md5M.Store(temp, password)
	log.Printf("%s MD5 routine --CLOSED", temp)
	wga.Done()
}

func solve256(password string, wgb *sync.WaitGroup) {
	temp := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	shaM.Store(temp, password)
	wgb.Done()
}

func GenHashMaps(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(f)
	var password string

	var wg sync.WaitGroup

	for scanner.Scan() {
		password = scanner.Text()

		wg.Add(2)
		go solveMD5(password, &wg)
		go solve256(password, &wg)

	}

	wg.Wait()
	f.Close()
	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
}

func GetSHA(hash string) (string, error) {
	//password, ok := shaM.Load("hash") //shalookup[hash]

	return fmt.Sprint(shaM.Load(hash)), nil

	//return "", errors.New("password does not exist")

}

//TODO
func GetMD5(hash string) (string, error) {
	return fmt.Sprint(md5M.Load(hash)), nil
	//return "", errors.New("not implemented")
}
