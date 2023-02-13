# Übungen zum Thema "Methoden und Interfaces"

## a) Methoden

### Bookshelf

Erstellen Sie ein `Bookshelf` Struct (Klasse), welches folgende Methoden anbietet:

* `add(b book)`
* `forIsbn(isbn string) book`
* `all() []books`

Erzeugen Sie eine Instanz dieses Bookshelfs.

Fügen Sie dort die in `main.go` vordefinierten Bücher ein und fragen diese dann einzeln bzw. gesamthaft ab.

### Stack (Bonusaufgabe)

Refaktorieren Sie Ihre `Stack` Implementierung aus dem Abschnitt "020/Slices" in eine objektorientierte Variante.

Machen Sie die Implementierung in einer Datei `stack\stack.go` (d.h. in Unterverzeichnis) mit Package "stack".

Dies erfordert allerdings auch, dass Sie den Struct-Bezeichner und alle Methodennamen mit einem **Großbuchstaben**
beginnen lassen, damit diese Deklarationen exportiert werden.

Achtung, es gibt keinen Konstruktor in Go. Falls Sie einen benötigen, so muss eine `NewStack()` Funktion
herhalten.

## c) Interfaces & Embedding

### Formatter & Parser

Erstellen Sie ein Interface `Formatter`, welches ein `int` in einen `string` formatieren kann.

Erstellen Sie außerdem ein Interface `Parser`, welches einen `string` in ein `int` parsen kann.

Erstellen Sie die folgenden Structs, welche durch ihre vorhandenen Methoden diese Interfaces implizit implementieren:

1. `BinaryFormatter` -- macht aus einer Zahl wie z.B. 42 den String "101010"
2. `BinaryParser` -- macht aus einem String wie z.B. "10" die Zahl 2

Entkommentieren Sie die `check()` Methode und rufen Sie diese mit Ihren Structs auf.

### LoggingParser

Erstellen Sie ein struct `LoggingParser`, welches mittels Embedding wie ein `Parser` auftreten kann.

Dieser `LoggingParser` soll jeden zu parsenden String ausgeben.

Ersetzen Sie in dem Aufruf an die `check()` Methode den alten Parser mit dem neuen.
