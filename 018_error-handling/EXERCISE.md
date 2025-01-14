# Übungen zum Thema "Error Handling"

Öffnen Sie eine lokale Datei in Go und lesen Sie deren Inhalt als Slice von Zeilen aus.

````go
func openAndReadFile(filename string) ([]string, error) {
}
````

Dabei hilft:

- `os.Open()`
- `scanner := bufio.NewScanner(file)`

Falls Fehler auftreten, sollen diese in einem neuen Fehler `verpackt zurückgegeben werden.

Denken Sie an ein `defer`, um das File schließen zu lassen.