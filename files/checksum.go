package isbn

func isbn10Checksum(numbers []byte) (result byte) {    
    for i := 0; i < 9; i ++ {
        result += numbers[i] * byte(i + 1)
    }
    
    return result % 11
}    