# Go Tour

A Tour of Go est une introduction au langage de programmation Go. Visite
https://go-tour-fr.appspot.com/ pour commencer la visite.

## Download/Install

Pour installer la visite depuis la source, commencez par
[installer Go](https://golang.org/doc/install) puis exécutez:

	$ go get github.com/dupoxy/go-tour-fr

Cela placera un binaire `go-tour-fr` dans le
répertoire `bin` de votre [espace de travail](https://golang.org/doc/code.html#Workspaces).
Le programme de visite peut être exécuté hors ligne.

## Contributing

Nous apprécions vos contributions.
Pour rendre le processus aussi transparent que possible, nous demandons ce qui suit:

* Allez-y et Forkez le projet et apportez vos modifications. Nous encourageons les pull requests pour discuter des changements.

pour plus de detail voir [CONTRIBUTING](CONTRIBUTING.md)

Pour exécuter le programme de visite localement:

```sh
go run .
```

Votre navigateur devrait maintenant s'ouvrir. Sinon, veuillez visiter [http://localhost:3999/](http://localhost:3999).

## Deploying

1.	To deploy go-tour-fr.appspot.com, run:

	```
	GO111MODULE=on gcloud app deploy --no-promote
	```

	This will create a new version, which can be viewed within the
	[go-tour-fr GCP project](https://console.cloud.google.com/appengine/versions?hl=fr&project=go-tour-fr).

2.	Check that the deployed version looks OK (click the version link in GCP).

3.	If all is well, click "Migrate Traffic" to move 100% of the tour.golang.org
	traffic to the new version.

4.	You're done.

## License

Unless otherwise noted, the go-tour source files are distributed
under the BSD-style license found in the LICENSE file.
