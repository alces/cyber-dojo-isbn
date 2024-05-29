package isbn

func isbn10Checksum(numbers []byte) byte {
    var result int
    
    for i := 0; i < 9; i ++ {
        result += int(numbers[i]) * (i + 1)
    }
    
    return byte(result % 11)
}

func isbn13Checksum(numbers []byte) byte {
    return 0
}