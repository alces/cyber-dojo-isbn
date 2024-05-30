package isbn

func Verify(isbn string) bool {
    return verifyChecksum(cleanUp(isbn))
}
