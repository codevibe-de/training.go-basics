# Übungen zum Thema "Methoden und Interfaces"

## a) Methoden

### Bookshelf

Erstellen Sie ein `Bookshelf` Struct, welches folgende Methoden anbietet:

* `add(b book)`
* `forIsbn(isbn string) book`
* `all() []books`

Erzeugen Sie eine Instanz dieses Structs.

Fügen Sie dort die in `main.go` vordefinierten Bücher ein und fragen diese dann einzeln bzw. gesamthaft ab.

### Stack (Bonusaufgabe)

Kopieren Sie Ihre Stack-Implementierung aus dem Verzeichnis `015-syntax/o_slices` hierher in eine Datei
`020_methods-interfaces/stack/stack.go`.

Ändern Sie den Package-Namen auf "stack".

Refaktorieren Sie dann den Code in eine objektorientierte Variante.

Die Ablage in einem neuen Package erfordert, dass Sie den Struct-Bezeichner und alle Methodennamen mit einem *
*Großbuchstaben** beginnen lassen, damit die Deklarationen exportiert werden.

Achtung, es gibt keinen Konstruktor in Go. Falls Sie einen benötigen, so muss eine `NewStack()` Funktion
herhalten.

## c) Interfaces: Formatter & Parser

*Diese Übung bitte mit reinem "Interface implementieren" lösen - nicht mit Embeddings.*

Erstellen Sie ein Interface `Formatter`, welches ein `int` in einen `string` formatieren kann.

Erstellen Sie außerdem ein Interface `Parser`, welches einen `string` in ein `int` parsen kann.

Erstellen Sie die folgenden Structs, welche durch ihre vorhandenen Methoden diese Interfaces implizit implementieren:

1. `BinaryFormatter` -- macht aus einer Zahl wie z.B. 42 den String "101010"
2. `BinaryParser` -- macht aus einem String wie z.B. "10" die Zahl 2

Entkommentieren Sie in `020_methods-interfaces/c_interfaces/main.go` die `check()` Methode und rufen Sie diese mit Ihren
Structs auf.

## d) Embeddings: LoggingParser

*In dieser Übung bitte mit Embeddings arbeiten.*

Erstellen Sie ein struct `LoggingParser`, welches mittels Embedding wie ein `Parser` (aus der vorherigen Übung)
auftreten kann.

Dieser `LoggingParser` soll jeden zu parsenden String ausgeben.

Ersetzen Sie in dem Aufruf an die `check()` Methode den alten Parser mit dem neuen.
