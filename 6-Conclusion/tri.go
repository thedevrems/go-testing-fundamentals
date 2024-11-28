package tri

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.
*/

func tri(tinit []int) (tfin []int) {
	// Si tinit est nil, on retourne un tableau vide
	if tinit == nil {
		return []int{}
	}
	// Créer un tableau tfin de la même taille que tinit
	tfin = make([]int, len(tinit))

	// Copier les éléments de tinit vers tfin
	for i := 0; i < len(tinit); i++ {
		tfin[i] = tinit[i]
	}

	// Tri à bulles
	n := len(tfin)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if tfin[j] > tfin[j+1] {
				// Échanger les éléments
				tfin[j], tfin[j+1] = tfin[j+1], tfin[j]
			}
		}
	}

	// Retourner le tableau trié
	return tfin
}
