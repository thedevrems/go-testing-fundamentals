package addition

import "testing"

// TestAddition vérifie que la fonction Addition renvoie le résultat attendu.
// Ce test additionne 2 et 3, puis vérifie que la somme est bien égale à 5.
//
// Paramètres :
//   t (*testing.T) : Objet utilisé pour signaler les erreurs dans le test.
//
// Fonctionnement du test :
//   Si le résultat de l'addition n'est pas égal à la valeur attendue (5),
//   t.Errorf est appelé pour signaler l'erreur, mais l'exécution du test continue.
//   Si le test échoue, un message est généré indiquant la valeur attendue et la valeur obtenue.
func TestAddition(t *testing.T) {
	result := Addition(2, 3) // On appelle la fonction Addition avec 2 et 3 comme arguments.
	expected := 5            // On définit le résultat attendu.

	// On compare le résultat obtenu avec le résultat attendu.
	if result != expected {
		// Si le résultat ne correspond pas, on signale une erreur avec t.Errorf.
		// Ce message sera affiché dans le terminal si le test échoue.
		t.Errorf("Erreur : attendu %d, obtenu %d", expected, result)
	}
}

// Par ailleurs, il existe également t.FailNow
// t.FailNow : Arrête immédiatement le test si une erreur est détectée.
