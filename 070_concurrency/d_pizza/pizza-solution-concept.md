# Lösungskonzept für "Pizza" Aufgabe

## Pizzeria

Das Konzept einer Pizzeria besteht aus zwei Channels:

1. Bestellung einer Pizza (`pizzaOrderCh`)
2. Benachrichtigung einer fertigen Pizza (`pizzaDeliverCh`)

## Worker   

Eine Anzahl an "Worker" Goroutinen sind für die Bearbeitung von Zutaten und Pizzen zuständig.

Um das Handling von Zutaten und dem Ofen zu erleichtern, implementieren beide das `Preparer` interface.

Ein Worker muss also:

1. `Prepare` Instanzen bearbeiten
2. eine Pizza im Ofen backen können

Für beide Tätigkeiten gibt es jeweils einen Input (zu tun) und Output (erledigt) Channel. 

## Supervisor

Eine "Supervisor" Goroutine bearbeitet Pizza Bestellungen, indem sie

* sich bestellte Pizzen merkt
* die Zutaten für eine bestellte Pizza zur Bearbeitung an die Arbeiter übergibt
* über fertig bearbeitete Zutaten informiert wird
* Pizzen mit fertigen Zutaten zum Backen übergibt 

Der Supervisor arbeitet mindestens so lange bis `pizzaOrderCh` geschlossen wurde