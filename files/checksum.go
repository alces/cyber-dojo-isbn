package isbn

func isbn10Checksum(numbers []byte) byte {
    var result int
    
    for i := 0; i < 9; i ++ {
        result += int(numbers[i]) * (i + 1)
    }
    
    return byte(result % 11)
}

func isbn13Checksum(numbers []byte) byte {
    var result int
    
    for i := 0; i < 12; i ++ {
        coef := 1
        
        if i % 2 == 1 {
            coef = 3
        }
        
        result += int(numbers[i]) * coef           
    }
    
    return byte(10 - (result % 10))
}