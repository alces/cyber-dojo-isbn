package isbn

import (
    "strings"
)

func Verify(isbn string) bool {
    return len(isbn) >= 10
}

func cleanUp(isbn string) string {
    if strings.Contains(isbn, "-") {
        return strings.ReplaceAll(isbn, "-", "")
    }
    
    return isbn
}