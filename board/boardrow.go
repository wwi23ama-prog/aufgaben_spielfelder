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
	for i := range row {
		row[i] = fill
	}
	return row
}

// MakeEmptyBoardRow erzeugt eine leere Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile
// und füllt die Zeile mit Leerzeichen.
func MakeEmptyBoardRow(length int) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Funktion MakeBoardRow.
	*/
	return MakeBoardRow(length, " ")
}

// RowContainsOnly prüft, ob die Zeile nur aus dem übergebenen String besteht.
func RowContainsOnly(row BoardRow, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice zu durchlaufen.
	   Brechen Sie die Schleife ab, sobald ein Element ungleich dem übergebenen String ist.
	*/
	for _, v := range row {
		if v != s {
			return false
		}
	}
	return true
}

// IsEmpty prüft, ob die Zeile leer ist.
func IsEmpty(row BoardRow) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion ContainsOnly.
	*/
	return RowContainsOnly(row, " ")
}
