package board

// BoardRow ist eine Zeile des Spielbretts.
// Wir verwenden Strings, um die Spielfiguren zu repräsentieren.
// Daher ist BoardRow ein Slice von Strings.
type BoardRow []string

// MakeBoardRow erzeugt eine Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile und den String,
// mit dem die Zeile gefüllt werden soll.
func MakeBoardRow(length int, fill string) BoardRow {
	row := make(BoardRow, length)
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice mit dem gewünschten String zu füllen.
	*/
	// TODO
	return row
}

// MakeEmptyBoardRow erzeugt eine leere Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile
// und füllt die Zeile mit Leerzeichen.
func MakeEmptyBoardRow(length int) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Funktion MakeBoardRow.
	*/
	// TODO
	return []string{}
}

// RowContainsOnly prüft, ob die Zeile nur aus dem übergebenen String besteht.
func RowContainsOnly(row BoardRow, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice zu durchlaufen.
	   Brechen Sie die Schleife ab, sobald ein Element ungleich dem übergebenen String ist.
	*/
	// TODO
	return false
}

// IsEmpty prüft, ob die Zeile leer ist.
func IsEmpty(row BoardRow) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion ContainsOnly.
	*/
	// TODO
	return false
}
