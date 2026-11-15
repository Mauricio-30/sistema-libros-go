package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lector = bufio.NewReader(os.Stdin)

// LeerTexto retorna el texto y un bool → RETORNO MÚLTIPLE.
func LeerTexto(mensaje string) (string, bool) {
	fmt.Print(mensaje)
	texto, err := lector.ReadString('\n')
	if err != nil {
		return "", false
	}
	return strings.TrimSpace(texto), true
}

// LeerEntero retorna el número y un bool → RETORNO MÚLTIPLE.
func LeerEntero(mensaje string) (int, bool) {
	texto, ok := LeerTexto(mensaje)
	if !ok {
		return 0, false
	}
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, false
	}
	return numero, true
}
