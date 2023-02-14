# Übungen zum Thema "Module und Pakete"

Tipps:
* `go get` ohne Argumente bereinigt und aktualisiert die deklarierten Abhängigkeiten
* `go get <module-path>` fügt das Modul der `go.mod` hinzu und lädt es herunter

## a) Modul nutzen

Erstellen Sie eine kleine Go-Anwendung, welche das Modul https://github.com/google/uuid nutzt. Dies können Sie in diesem
Verzeichnis tun.

Ihre Anwendung benötigt eine `go.mod` Datei, welche Sie mittels `go mod init` anlegen können.

Das benötigte Modul deklarieren Sie in dieser Datei und führen dann `go get` aus.

Und/oder Sie führen `go get <module-path>` aus, um das Modul importieren zu lassen. Was verändert sich in der `go.mod`
Datei?

Schauen Sie nun nach einem älteren Release des Moduls und nutzen dieses.

## b) Eigenes Modul erstellen

Erstellen Sie ein Github Repository und checken dieses aus.

Dort erstellen Sie nun eine `go.mod` Datei mit dem korrekten Module-Path.

Legen Sie zwei Packages in Ihrem Modul an.

Pushen Sie Ihren Code.

Literatur: https://www.digitalocean.com/community/tutorials/how-to-distribute-go-modules

## c) Eigenes Modul nutzen

Aktualisieren Sie nun Ihre kleine Go-Anwendung aus Übung a) indem Sie Ihr eigenes Modul dort hinzufügen und im Code
verwenden.

## d) Experimentelles Package nutzen

Schauen Sie sich den Code in `main.go` an -- sowie das experimentelle Package https://pkg.go.dev/golang.org/x/exp/slices.

Dort gibt es eine Funktion zum Vergleichen zweier Slices.

Verbessern Sie die bestehende `compareSlices()` Funktion, indem diese das `slices` Package nutzt.