# Übungen zum Thema "Concurrency"

## a) Channels: Lesen und Schreiben

Starten Sie eine Goroutine, die in einen Channel einen Wert schreibt (z.B. ein `int`).

Dieser Wert soll in der `main()` Funktion ausgelesen und ausgegeben werden.

## b) Gemeinsamer Start

Lassen Sie drei Goroutinen starten, die eine zufällige Anzahl von 1 - 5 Sekunden brauchen, bis diese einsatzbereit sind.
Dazu können Sie die fertige Funktion `startup()` nutzen.

Sobald jede Goroutine bereit ist, soll sie auf ein gemeinsames Startsignal (Channel!) warten und dann ein "X" ausgeben.

In der `main()` Methode schreiben Sie in den Channel, um jeder Goroutine das Startsignal zu geben.