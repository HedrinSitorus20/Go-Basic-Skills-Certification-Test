package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "unicode"
    
)



/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {
    str = strings.TrimSpace(str)
    
    var result strings.Builder
    for _, char := range str{
        if !unicode.IsDigit(char){
            result.WriteRune(char)
        }
    }
    // Reverse the string
    reversed := result.String()
    var reversedResult strings.Builder
    for i := len(reversed) - 1; i >= 0; i-- {
        reversedResult.WriteByte(reversed[i])
    }

    return reversedResult.String()
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    str := readLine(reader)

    result := ModifyString(str)

    fmt.Fprintf(writer, "%s\n", result)

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
