package isbn

func Verify(isbn string) bool {
    return len(isbn) >= 10
}
