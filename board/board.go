package board

// Board ist ein Spielbrett.
// Wir verwenden eine Slice von BoardRows, um das Spielbrett zu repräsentieren.
type Board []BoardRow

// MakeBoard erzeugt ein Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und den String, mit dem das Spielbrett gefüllt werden soll.
func MakeBoard(width, height int, fill string) Board {
	board := make(Board, height)
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice mit den gewünschten Zeilen zu füllen.
	   Erzeugen Sie für jede Zeile eine neue Slice mit der Funktion MakeBoardRow.
	*/
	for i := range board {
		board[i] = MakeBoardRow(width, fill)
	}
	return board
}

// MakeEmptyBoard erzeugt ein leeres Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und füllt das Spielbrett mit Leerzeichen.
func MakeEmptyBoard(width, height int) Board {
	/* Hinweis:
	   Verwenden Sie die Funktion MakeBoard.
	*/
	return MakeBoard(width, height, " ")
}

// GetRow gibt die Zeile an der übergebenen Position zurück.
func GetRow(b Board, y int) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Index-Notation, um die Zeile zu erhalten.
	*/
	// tag::solution[]
	return b[y]
}

// GetColumn gibt die Spalte an der übergebenen Position zurück.
func GetColumn(b Board, x int) BoardRow {
	column := make(BoardRow, len(b))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Index-Notation, um jeweils das x-te Element der Zeile zu erhalten.
	*/
	for i, row := range b {
		column[i] = row[x]
	}
	return column
}

// BoardRowContainsOnly prüft, ob die angegebene Zeile nur aus dem übergebenen String besteht.
func BoardRowContainsOnly(b Board, row int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetRow.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	return RowContainsOnly(GetRow(b, row), s)
}

// BoardColumnContainsOnly prüft, ob die angegebene Spalte nur aus dem übergebenen String besteht.
func BoardColumnContainsOnly(b Board, column int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetColumn.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	return RowContainsOnly(GetColumn(b, column), s)
}

// BoardContainsOnly prüft, ob das Spielbrett nur aus dem übergebenen String besteht.
func BoardContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	for row := range b {
		if !BoardRowContainsOnly(b, row, s) {
			return false
		}
	}
	return true
}
