package isbn

import (
    "regexp"
    "strings"
)

func Verify(isbn string) bool {
    return len(isbn) >= 10
}

func cleanUp(isbn string) string {
    for _, c := range []string{"-", " "} {
        if strings.Contains(isbn, c) {
            return strings.ReplaceAll(isbn, c, "")
        }
    }
    
    return isbn
}

func isISBN10(isbn string) bool {
    match, err := regexp.MatchString(`^\d{9}[\dX]$`, isbn)
    return err == nil && match
}

func isISBN13(isbn string) bool {
    match, err := regexp.MatchString(`^\d{13}$`, isbn)
    return err == nil && match
}

func stringToNumbers(isbn string) []byte {
    result := make([]byte, len(isbn))
    
    for i, c := range isbn {
        result[i] = byte(c) - 48
    }
    
    return result
}