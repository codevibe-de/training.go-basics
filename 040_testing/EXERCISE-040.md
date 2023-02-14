# Übungen zum Thema "Testing"

Tipp: Man kann Tests auf der Kommandozeile mit `go test` ausführen. Natürlich auch in der IDE - VSC hat
hierfür eine eigene Ansicht (Reagenzglas).

## Fibonacci

Schauen Sie sich die Fibonacci Implementierung in der Datei `fib.go` an.

Dort ist einiges nicht richtig.

Nutzen Sie die vorhandenen Tests, um die Implementierung zu fixen.

## Stack

Schreiben Sie Tests für Ihre Stack-Methoden, d.h. für die objektorientierte Variante.

In den Tests sollen alle Methoden geprüft werden, d.h. `Push()`, `Pop()` und `Peek()`.

Hierfür können Sie folgende Testarten nutzen/kombinieren:

* einfache Tests
* table-driven Tests
* Subtests
* Testen mit Beispielen (Hinweis: Die Namenskonvention hier wäre dann z.B. `func ExampleStack_Peek()`, siehe
  auch https://go.dev/blog/examples)

