package board

import "fmt"

// PrintBoardRow gibt die Zeile des Spielbretts auf die Konsole aus.
// Die Funktion erwartet eine Zeile des Spielbretts.
func PrintBoardRow(row BoardRow) {
	output := "| "
	// TODO
	fmt.Println(output)
}

// Divider gibt den Trenn-String für die Zeilen des Spielbretts zurück.
// Die Funktion erwartet die Länge der Zeile.
// Der Trenn-String ist von der Form "+---+---+---+".
func RowDivider(length int) string {
	result := "+"
	// TODO
	return result
}

// PrintBoard gibt das Spielbrett auf die Konsole aus.
// Die Funktion erwartet das Spielbrett.
func PrintBoard(b Board) {
	divider := RowDivider(len(b[0]))
	fmt.Println(divider)
	// TODO
}
