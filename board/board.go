package board

// Board ist ein Spielbrett.
// Wir verwenden eine Slice von BoardRows, um das Spielbrett zu repräsentieren.
type Board []BoardRow

// MakeBoard erzeugt ein Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und den String, mit dem das Spielbrett gefüllt werden soll.
func MakeBoard(width, height int, fill string) Board {
	board := make(Board, height)
	// TODO
	return board
}

// MakeEmptyBoard erzeugt ein leeres Spielbrett.
// Die Funktion erwartet die Breite und Höhe des Spielbretts
// und füllt das Spielbrett mit Leerzeichen.
func MakeEmptyBoard(width, height int) Board {
	// TODO
	return Board{}
}

// GetRow gibt die Zeile an der übergebenen Position zurück.
func GetRow(b Board, y int) BoardRow {
	// TODO
	return b[0]
}

// GetDiagRight gibt die Diagonale von links oben nach rechts unten zurück.
func GetDiagRight(b Board) BoardRow {
	diag := make(BoardRow, len(b))
	// TODO
	return diag
}

// GetDiagLeft gibt die Diagonale von rechts oben nach links unten zurück.
func GetDiagLeft(b Board) BoardRow {
	diag := make(BoardRow, len(b))
	// TODO
	return diag
}

// GetColumn gibt die Spalte an der übergebenen Position zurück.
func GetColumn(b Board, x int) BoardRow {
	column := make(BoardRow, len(b))
	// TODO
	return column
}

// BoardRowContainsOnly prüft, ob die angegebene Zeile nur aus dem übergebenen String besteht.
func BoardRowContainsOnly(b Board, row int, s string) bool {
	// TODO
	return false
}

// BoardColumnContainsOnly prüft, ob die angegebene Spalte nur aus dem übergebenen String besteht.
func BoardColumnContainsOnly(b Board, column int, s string) bool {
	// TODO
	return false
}

// BoardDiagRightContainsOnly prüft, ob die Diagonale von links
// oben nach rechts unten nur aus dem übergebenen String besteht.
func BoardDiagRightContainsOnly(b Board, s string) bool {
	// TODO
	return false
}

// BoardDiagLeftContainsOnly prüft, ob die Diagonale von rechts
// oben nach links unten nur aus dem übergebenen String besteht.
func BoardDiagLeftContainsOnly(b Board, s string) bool {
	// TODO
	return false
}

// BoardContainsOnly prüft, ob das Spielbrett nur aus dem übergebenen String besteht.
func BoardContainsOnly(b Board, s string) bool {
	// TODO
	return false
}
