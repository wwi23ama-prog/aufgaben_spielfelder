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
	// TODO
	return board
}

// MakeEmptyBoard erzeugt ein leeres Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und füllt das Spielbrett mit Leerzeichen.
func MakeEmptyBoard(width, height int) Board {
	/* Hinweis:
	   Verwenden Sie die Funktion MakeBoard.
	*/
	// TODO
	return Board{}
}

// GetRow gibt die Zeile an der übergebenen Position zurück.
func GetRow(b Board, y int) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Index-Notation, um die Zeile zu erhalten.
	*/
	// TODO
	return b[0]
}

// GetColumn gibt die Spalte an der übergebenen Position zurück.
func GetColumn(b Board, x int) BoardRow {
	column := make(BoardRow, len(b))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Index-Notation, um jeweils das x-te Element der Zeile zu erhalten.
	*/
	// TODO
	return column
}

// GetDiagRight gibt die Diagonale von links oben nach rechts unten zurück.
func GetDiagRight(b Board) BoardRow {
	diag := make(BoardRow, len(b))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Da Zeile 0 oben und Spalte 0 links ist, brauchen Sie die Elemente (i,i),
	   also aus jeder Zeile das i-te Element.
	*/
	// TODO
	return diag
}

// GetDiagLeft gibt die Diagonale von rechts oben nach links unten zurück.
func GetDiagLeft(b Board) BoardRow {
	diag := make(BoardRow, len(b))
	/* Hinweis:
	Gehen Sie analog zu GetDiagRight vor.
	Überlegen Sie sich, welches Element Sie dieses Mal aus jeder Zeile benötigen.
	*/
	// TODO
	return diag
}

// BoardRowContainsOnly prüft, ob die angegebene Zeile nur aus dem übergebenen String besteht.
func BoardRowContainsOnly(b Board, row int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetRow.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// TODO
	return false
}

// BoardColumnContainsOnly prüft, ob die angegebene Spalte nur aus dem übergebenen String besteht.
func BoardColumnContainsOnly(b Board, column int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetColumn.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// TODO
	return false
}

// BoardContainsOnly prüft, ob das Spielbrett nur aus dem übergebenen String besteht.
func BoardContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// TODO
	return false
}
