# 6. Bonnes Pratiques pour les Tests Unitaires
- Isolation : Testez chaque unité indépendamment (ne dépendez pas de bases de données, d'API externes, etc.).
- Lisibilité : Les tests doivent être clairs et simples à comprendre.
- Nom significatif : Donnez des noms explicites aux fonctions de test (TestAddition au lieu de TestA).
- Cas limites : Pensez à tester les cas d'erreur ou extrêmes.
- Automatisation : Intégrez les tests dans votre CI/CD pour garantir leur exécution régulière.

# Définitions :

- CI (Continuous Integration) : Cela implique l'intégration régulière du code dans un dépôt centralisé (souvent plusieurs fois par jour). Chaque changement est automatiquement testé pour s'assurer qu'il n'introduit pas de bugs. L'objectif est de détecter rapidement les erreurs et d'assurer une base de code fonctionnelle à tout moment.

- CD (Continuous Delivery / Continuous Deployment) :

    - Continuous Delivery : Cela va un pas plus loin que l'intégration continue. Après l'intégration, le code est automatiquement préparé pour être déployé en production, mais le déploiement lui-même est souvent un processus manuel. L'idée est de rendre le déploiement plus rapide et moins risqué.
    - Continuous Deployment : Ici, chaque changement validé est automatiquement déployé en production sans intervention manuelle. Cela permet de livrer de nouvelles fonctionnalités plus rapidement et plus fréquemment.

En résumé, CI/CD vise à rendre le processus de développement plus rapide, fiable et automatisé, en garantissant que les tests sont effectués et que le code peut être livré en production de manière fluide et continue.


[Suite du cours](../6-Conclusion/conclusion.md)