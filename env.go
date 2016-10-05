// Package env is used for reading key pair values from a .env file. The goal is to provide an easy location to store sensitive data 
// which can be kept out of source control. .env file format is API_KEY=KEYHERE

package env 

import (
    "os"
    "io/ioutil"
    "strings"
    "bufio"
)

func Env(name string) string {
    f, err := os.Open(".env")
    if err != nil {
        return ""
    }

    defer f.Close()

    bs, err := ioutil.ReadAll(f)
     if err != nil {
        return ""
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
        return ""
    }

    if val, ok := m[name]; ok {
        return val
    } 

    return ""
}