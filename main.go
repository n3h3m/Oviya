package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <file.ov>")
        os.Exit(1)
    }

    filePath := os.Args[1]
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Failed to open %s: %v\n", filePath, err)
        os.Exit(1)
    }
    defer file.Close()

    cCode := "#include <stdio.h>\n\nint main() {\n"
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if strings.HasPrefix(line, "print ") {
            str := strings.TrimPrefix(line, "print ")
            str = strings.Trim(str, "\"")
            cCode += fmt.Sprintf("    printf(\"%s\\n\");\n", str)
        }
    }

    cCode += "    return 0;\n}\n"

    outputPath := strings.TrimSuffix(filePath, ".ov") + ".c"
    err = os.WriteFile(outputPath, []byte(cCode), 0644)
    if err != nil {
        fmt.Printf("Failed to write C file: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("✅ Generated: %s\n", outputPath)
}
