package board

import "fmt"

// PrintBoardRow gibt die Zeile des Spielbretts auf die Konsole aus.
// Die Funktion erwartet eine Zeile des Spielbretts.
func PrintBoardRow(row BoardRow) {
	output := "| "
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice zu durchlaufen.
	   Hängen Sie an die Variable output jeweils das Element der Slice
	   und den String "|" an.
	*/
	// TODO
	fmt.Println(output)
}

// Divider gibt den Trenn-String für die Zeilen des Spielbretts zurück.
// Die Funktion erwartet die Länge der Zeile.
// Der Trenn-String ist von der Form "+---+---+---+".
func RowDivider(length int) string {
	result := "+"
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um den Trenn-String zu erzeugen.
	   Hängen Sie in jedem Schleifendurchlauf den String "---+" an.
	*/
	// TODO
	return result
}

// PrintBoard gibt das Spielbrett auf die Konsole aus.
// Die Funktion erwartet das Spielbrett.
func PrintBoard(b Board) {
	divider := RowDivider(len(b[0]))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen des Spielbretts zu durchlaufen.
	   Verwenden Sie die Funktion PrintBoardRow, um die Zeile auszugeben.
	   Geben Sie nach jeder Zeile den Trenn-String aus.
	*/
	fmt.Println(divider)
	// TODO
}
