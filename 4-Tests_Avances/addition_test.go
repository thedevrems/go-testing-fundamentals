package addition

import "testing"

// TestAdditionCases teste la fonction Addition avec plusieurs cas de test.
// La fonction vérifie différents les couples d'entrées et leur résultat attendu.
//
// Paramètres :
//   t (*testing.T) : Objet utilisé pour signaler les erreurs dans le test.
//
// Fonctionnement du test :
//   Les cas de test sont définis dans une structure composée d'entrées (a, b)
//   et du résultat attendu (expected). Pour chaque cas, la fonction Addition
//   est appelée avec a et b, et le résultat est comparé avec la valeur attendue.
//
//   Si le résultat ne correspond pas à la valeur attendue, une erreur est
//   enregistrée à l'aide de t.Errorf.
func TestAdditionCases(t *testing.T) {
	// Définition des différents cas de test avec leurs entrées et résultats attendus.
	cases := []struct {
		a, b     int // Les deux entiers à additionner.
		expected int // Le résultat attendu de l'addition.
	}{
		{2, 3, 5},       // Cas 1 : 2 + 3 = 5
		{-1, -1, -2},    // Cas 2 : -1 + -1 = -2
		{0, 0, 0},       // Cas 3 : 0 + 0 = 0
		{100, 200, 300}, // Cas 4 : 100 + 200 = 300
	}

	// Boucle sur chaque cas de test.
	for _, c := range cases {
		result := Addition(c.a, c.b) // On appelle la fonction Addition avec les valeurs a et b.

		// Si le résultat ne correspond pas à la valeur attendue, on enregistre une erreur.
		if result != c.expected {
			t.Errorf("Erreur pour Addition(%d, %d) : attendu %d, obtenu %d", c.a, c.b, c.expected, result)
		}
	}
}