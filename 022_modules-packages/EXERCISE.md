# Übungen zum Thema "Module und Pakete"

Tipps:

* `go get` ohne Argumente bereinigt und aktualisiert die deklarierten Abhängigkeiten
* `go get <module-path>` fügt das Modul der `go.mod` hinzu und lädt es herunter

## a) Modul nutzen

Entwickeln Sie eine kleine Go-Anwendung, welche das Modul https://github.com/google/uuid nutzt.

Hierfür müssen Sie irgendwo einen neuen Ordner anlegen.

Ihre Anwendung benötigt eine `go.mod` Datei, welche Sie mittels `go mod init <module-path>` dort erzeugen lassen. Da wir
eine lokale Anwendung entwickeln, ist der Pfad nicht von Bedeutung.

Das benötigte UUID Modul deklarieren Sie in dieser Datei und führen dann `go get` aus.

Und/oder Sie führen `go get <module-path>` (mit oder ohne Version) aus, um das Modul importieren zu lassen.
Was verändert sich in der `go.mod` Datei?

Schauen Sie nun nach einem älteren Release des Moduls und nutzen dieses. Lassen Sie sich mögliche Upgrades
mit `go list -m -u all` anzeigen. Anschließend können Sie das Upgrade mit `go get -u` durchführen.

Schlußendlich können Sie natürlich das Package "uuid" in der `main.go` Datei importieren und nutzen (z.B.
mittels `fmt.Println(uuid.New())`

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

Schauen Sie sich den Code in `022_modules-packages/main.go` an -- sowie das experimentelle
Modul "pkg.go.dev/golang.org/x/exp" mit dem Package "slices": https://pkg.go.dev/golang.org/x/exp

Dort gibt es eine Funktion zum Vergleichen zweier Slices.

Verbessern Sie die bestehende `compareSlices()` Funktion, indem diese das `slices` Package nutzt.