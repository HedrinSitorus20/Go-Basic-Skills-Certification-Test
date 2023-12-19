package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
    // Initialize a counter for set bits
    count := int32(0)

    // Iterate through each bit and count set bits
    for i := 0; i < 32; i++ {
        // Check if the current bit is set (1)
        if num&(1<<i) != 0 {
            count++
        }
    }

    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    numTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    num := uint32(numTemp)

    result := countBits(num)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
