<img style="width:100%;" src="/github-banner.png">

# Travaux Pratique 4 - Rédaction de user stories et déroulement d’un sprint

Le TP3 est un programme permettant d'intéragir avec plusieurs API et fournit une API pour une interface Web.
Ce programme utilise le langage GOLANG, l'installation ci-dessus est nécessaire à son fonctionnement.
Le projet a été lié à l'outils d'intégration continue Travis ainsi que Heroku.

## Version <img style="width:10%;" src="https://travis-ci.com/glo2003/team10.svg?token=sfos4hew5nBWSWgqfxrp&branch=master">

`1.0` Interaction avec l'API de Github/Création du ReadMe.

`2.0` Integration Continue avec Travis.

`3.0` Déploiement en continue et métriques associées au projet.

`4.0` Ajout de nouveau contenu sur l'application

## Installation

Avant de pouvoir utiliser le programme, l'installation de plusieurs packages est nécessaire.

**Starter kit go**
```bash
github.com/glo2003/starter-kit-go
```
lien: <https://github.com/glo2003/starter-kit-go>

**Oauth2**
```bash
go get golang.org/x/oauth2

```
lien: <https://github.com/golang/oauth2>

**Récuperer les sources**
```bash
git clone https://github.com/glo2003/team10.git
```

## Utilisation

* Faire l'installation
* Ouvrir une invite de commande
* Set votre token GitHub
```bash
set GITHUB_TOKEN=votre_token
```
* Lancer le serveur:
```bash
go run /server.go
```
* Lancer une page dans votre navigateur web et entrer dans la barre de recherche: *localhost:PORT*
* La page / permet d'avoir un aperçu des fichiers dans le dépôt.
* La page /projects permet d'avoir un aperçu des données reçues au format JSON.

## Nouvelles applications

* Il est désormait possible de rentrer son token GitHub depuis l'interface de l'application la commande Set n'est donc plus indispensable.
* Nouveauté sur le status des répertoires. En plus des informations précédentes(issue, contributor et langage), des nouvelles ont fait leur apparition. Vous pouvez maintenant savoir les branches disponibles sur un répertoire ainsi que sont status Travis sans être obliger de passer par GitHub.


# Contributions

Instruction pour l'utilisation du Git.
* Pour la mise à jour de document, les sections mises à jour doivent être spécifiées dans la partie version.
* Pour les mises à jour de la partie documentation, les commits doivent se faire sur la branche /documentation tant que la documentation n'est pas dans sa version final.
* Pour les mises à jour dans la partie application, les commits doivent se faire sur la branche /application tant que le code n'est pas fonctionnel ou est dans sa version final.
* Pour les mises à jour dans la partie API, les commits doivent se faire sur la branche /api tant que le code n'est pas fonctionnel ou est dans sa version final.
* Après commit, vérifier l'intégration avec le retour du logiciel Travis.

Les messages de commit doivent contenir les fichiers modifiés, les fonctionnalités modifiées.
Les fusions(merges) se font quand les fonctionnalités implémentées sont fonctionnelles et correctes pour le fonctionnement entier de l'application.

Si un bug est décelé dans une partie précédemment push, une branche Bug peut être créé afin de corriger ce bug et tester l'intégriter du programme avant d'être merge.

## Vérification du style de code

Le projet intègre la vérification du style de code. Pour plus d'information sur le style de code du langage Go, vous pouvez consulter cette [documentation](https://golang.org/doc/effective_go.html).

La vérification du style de code se fait avec la commande `go fmt`, pour plus d'information consulter cette [page](http://blog.golang.org/go-fmt-your-code).

Vous pouvez suivre le build de travis grâce au sticker à coté de la section Version.
=======

# go-getting-started

Les applications Go peuvent facilement être déployer sur Heroku.

Cette application supporte [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go).

## Deployment address

Travis deploie automatiquement sur Heroku quand la compilation passe.

```sh
https://git.heroku.com/team10-heroku.git 
```

## Documentation

Pour plus d'informations sur l'utilisation de Go avec Heroku, allez voir l'article suivant :

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
