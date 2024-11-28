# 3. Exemple Simple : Tester une Fonction
Dans ce dossier, vous pouvez retrouver un exemple de cas de test.
Dans notre cas, voici la fonction [addition.go](./addition.go) avec son fichier de test nommé [addition_test.go](./addition_test.go)

Attention : Veuillez toujours respecter cette structure au niveau des noms de fichiers.

# Comment éxécuter des tests ?
## Pour exécuter les tests :

1. Ouvrez un terminal dans le dossier contenant votre code.
2. Les commandes importantes :
Exécution de tous les tests
```bash
go test
```
Exécution du test TestAddition
```bash
go test -run TestAddition
```

3. Vous verrez un résultat indiquant si les tests ont réussi ou échoué.

## Exécution avec plus de détails
Pour des détails supplémentaires :
```bash
go test -v
```
### go test -v (mode verbose) :
- La commande -v (verbose) donne une sortie plus détaillée.
- Pour chaque test :
    - Le nom de chaque test est affiché.
    - Son état (PASS, FAIL, ou SKIP) est indiqué.
    - Les messages t.Log et autres informations de diagnostic (si présents) sont également affichés.
- Exemple :
    ``` bash
    === RUN   TestAddition
    --- PASS: TestAddition (0.00s)
    === RUN   TestSubtraction
    --- PASS: TestSubtraction (0.00s)
    ok   my/package   0.123s
    ```
### Quand utiliser -v ?
- Pour déboguer ou comprendre le comportement des tests.
- Pour voir des tests individuels et leurs sorties détaillées.


[Suite du cours](../4-Tests_Avances/test_avances.md)