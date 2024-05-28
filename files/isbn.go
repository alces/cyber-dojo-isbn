package isbn

import (
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