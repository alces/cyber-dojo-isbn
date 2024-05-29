package isbn

import (
    "fmt"
)

func isbn10Checksum(numbers []byte) byte {
    var result int
    
    for i := 0; i < 9; i ++ {
        result += numbers[i] * (i + 1)
    }
    
    return byte(result % 11)
}    