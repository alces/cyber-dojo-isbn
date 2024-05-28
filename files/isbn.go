package isbn

func Verify(isbn string) bool {
    return len(isbn) >= 10
}

func cleanUp(isbn string) string {
    return ""
}