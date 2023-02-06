# Übungen

## a) Variablen & Konstanten

Sie können für diese Übungen die Datei `020-syntax/a_variables-and-constants/main.go` nutzen.

### Variablen

Deklarieren Sie auf verschiedene Weisen Variablen.

Welche Zuweisungen von einer zu anderen gehen, welche nicht?

### Konstanten

Deklarieren Sie nicht-gruppierte und gruppierte Konstanten.

Was passiert, wenn diese gemischt in einem Ausdruck genutzt werden?

### Iota

Schreiben Sie eine Gruppe an Konstanten, welche die Werte `KB`, `MB` und `GB` deklariert.

Nutzen Sie hierfür `iota` zur automatischen Berechnung.

Tipp: mit `1 << 10` shiften Sie eine bestehende Zahl um 10 bits nach vorne

## b) Pointers

Nutzen Sie die Datei `020-syntax/b_pointers/main.go`.

### miniSort

Verändern Sie die Funktion `miniSort` und deren Aufruf derart, dass die gewünschte Ausgabe
erfolgt.

## c) Types

### Umwandlung cm in m

Deklarieren Sie eigene Typen für `centimeter` und `meter`. Wo muss die Deklaration erfolgen?

Schreiben Sie eine Funktion, die Zentimeter in Meter umrechnet.

Geben Sie das Ergebnis aus, sodass der gewünschte Output erfolgt.

## 

## d) Structs

Sie möchten folgendes JSON mit Structs verarbeiten:

````json
{
  "produktId": "P-123",
  "NAME": "Pizza Salami",
  "Preis": {
    "betrag-je-stueck": 12.99,
    "währung": "EUR"
  } 
}
````

Erstellen Sie die dafür nötigen Structs. Diese sollen nur englische Attribute verwenden
(also "productId" statt "produktId") und in Camel-Case notiert werden ("unitPrice" statt "betrag-je-stueck").

Legen Sie dann eine Struct Instanz mit den obigen Daten an.

## e) Funktionen

### Fibonacci

Schreiben Sie eine Funktion mit folgendem Input/Output Verhalten:

````go
numbers, count := fibonacci(50)
fmt.Println(numbers) // "1 1 2 3 5 8 13 21 34"
fmt.Println(count)  // 9
````

### Summierung mit Offset

Schreiben Sie eine Summierungsfunktion, die aus einer beliebigen Anzahl an `int` Werten eine
Gesamtsumme berechnet.

Über einen zusätzlichen Offset-Wert soll diese Summe verändert werden können.

### Statistiken

Ähnlich der Anzeige von Statistiken in Excel, wenn mehrere Zellen selektiert sind, benötigen wir
eine Funktion, die für eine beliebige Anzahl an `int` Werten die folgenden Statistiken berechnet:
- Anzahl Werte
- kleinster
- größter
- Summe
- Durchschnitt

## f) First Class Citizen Functions

Deklarieren Sie einen Funktionstyp mit folgender Signatur:
````go
funct(int, int) int
````

Implementieren Sie diese Funktion in mehreren Varianten:
* Addition
* Subtraktion
* Multiplikation

Rufen Sie nun ihre Funktion (auch gerne verschachtelt) auf und lassen
das Ergebnis ausgeben.

## ?) Arrays

Implementieren Sie eine Methode, die auf einem `int` Array 

## ?) Slices

Implementieren Sie einen **Stack** von `string` Werten mithilfe eines Slice.

Dieser soll folgende Funktionen implementieren:
````go
func Push(stack []string, v string) {}
func Peek(stack []string) (string) {}
func Pop(stack []string) (string) {}
````