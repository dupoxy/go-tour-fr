# Go Tour

A Tour of Go est une introduction au langage de programmation Go. Visite
https://go-tour-fr.appspot.com/ pour commencer la visite.

## Télécharger / installer

Pour installer la visite depuis la source, commencez par
[installer Go](https://golang.org/doc/install) puis exécutez:

	$ go get github.com/dupoxy/go-tour-fr

Cela placera un binaire `go-tour-fr` dans le
répertoire `bin` de votre [espace de travail](https://golang.org/doc/code.html#Workspaces).
Le programme de visite peut être exécuté hors ligne.

## Contribuer

Nous apprécions vos contributions.
Pour rendre le processus aussi transparent que possible, nous demandons ce qui suit:

* Allez-y et Forkez le projet et apportez vos modifications. Nous encourageons les pull requests pour discuter des changements.

pour plus de detail voir [CONTRIBUTING](CONTRIBUTING.md)

Pour exécuter le programme de visite localement:

```sh
go run .
```

Votre navigateur devrait maintenant s'ouvrir. Sinon, veuillez visiter [http://localhost:3999/](http://localhost:3999).

## Déployer

1.	Afin de déployer go-tour-fr.appspot.com, lancez la commande suivante :

	```
	GO111MODULE=on gcloud app deploy --no-promote
	```

	Cela va créer une nouvelle version, visible depuis le
	[projet GCP go-tour-fr](https://console.cloud.google.com/appengine/versions?hl=fr&project=go-tour-fr).

2.	Vérifiez que la version déployée a l'air correcte (cliquez sur le lien de la
	version dans GCP).

3.	Si tout est bon, cliquez sur « migrer le traffic » pour déplacer 100% du
	traffic de https://go-tour-fr.appspot.com vers la nouvelle version.

4.	Vous avez terminé.

## Licence

Sauf indication contraire, les fichiers sources de `go-tour` sont distribués
sous la licence de type BSD qui se trouve dans le fichier LICENSE.
