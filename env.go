// Package env is used for reading key pair values from a .env file. The goal is to provide an easy location to store sensitive data 
// which can be kept out of source control. .env file format is API_KEY=KEYHERE

package env 

import (
    "os"
    "io/ioutil"
    "strings"
    "bufio"
    "errors"
)

/**
 * Get
 * 
 * Method takes a name string for the lookup term. The get method will then
 * scan the root .env file for the key and if it finds it, it will return 
 * the value associated with the pair, if not it will return an error.
 * 
 * @param name string     
 * 
 * @return string, error
 */
func Get(name string) (string, error) {
    f, err := os.Open(".env")
    if err != nil {
        return "", err
    }

    defer f.Close()

    bs, err := ioutil.ReadAll(f)
     if err != nil {
        return "", err
    }

    fileContent := string(bs)

    scanner := bufio.NewScanner(strings.NewReader(fileContent))

    m := map[string]string{}
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), "=")
        if len(parts) == 2 {
            m[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
        }
    }
    if err := scanner.Err(); err != nil {
        return "", err
    }

    if val, ok := m[name]; ok {
        return val, nil
    } 

    return "", errors.New("Key not found in environment file")
}