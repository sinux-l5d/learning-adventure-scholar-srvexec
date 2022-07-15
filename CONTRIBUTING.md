# Structure

## Environment vs langage

Un environment utilise un langage dans un environnement d'exécution spécifique.
Si l'on veut simplement executer le code d'un langage, on va créer un environnement générique pour ce langage.

Par contre, si l'on peut exécuter du code python pour un cours de mathématiques, on va créer un environnement python-science contenant des paquets comme *numpy* ou *scipy*. Comme on ajoute des paquets, il se peut que la façon d'exécuter du code soit différente, on doit donc créer un fichier d'environnement en Go. Celui-ci peut simplement rajouté les imports des paquets nécessaires.

En résumé : 
* Nouveau langage : un fichier dockerfile, un fichier de langage et un fichier d'environment générique.
* Nouvel environnement : un fichier dockerfile et un fichier d'environment spécifique.

## Implémenter un nouveau langage

Remplacez LANGAGE par le nom du langage que vous souhaitez implémenter.

1. Créez un fichier `LANGAGE.go` dans `languages/` contenant les fonctions généralement utiles pour ce langage (e.g.: en python, Indent pour indenter du code).
2. Créer un fichier `LANGAGE-generic.go` dans `environments/` contenant le code minimal pour executer du code de ce langage. Il doit a minima définir une variable global MainEnvironment qui contient le nom de l'environnement (LANGAGE-generic) et la fonction à laquelle passer le code et l'exercice.

## Implémenter un nouvel environnement

Remplacez ENVIRONMENT par le nom de l'environnement que vous souhaitez implémenter.

1. Créez un fichier `ENVIRONMENT.go` dans `environments/` contenant une variable global MainEnvironment qui contient le nom de l'environnement et la fonction à laquelle passer le code et l'exercice.

## Convention de nommage des tags

Lorsque l'on build un environnement python_generic, on ne veux pas mettre dans le binaire le code qui permet d'exécuter du Java. 

Pour ça, on utilise des tags lors de la compilation.

Pour les fichiers dans `languages/`, on utilise le tag `libLANGUAGE` (e.g.: `libpython`)

Pour les fichiers dans `environments/`, on utilise le tag `LANGUAGE_ENVIRONMENT` (e.g.: `python_science`, `python_generic`)

En implémenant un nouveau language, il faut écrire un fichier `LANGUAGE_generic.go` qui contient ce qu'il faut pour exécuter du code de ce langage. Il servira de référence pour d'autre environnements.

Ainsi, lorsque l'on veut build l'environment python-science, le script `build.sh` va inclure les tags `python_science` ET `libpython`.

# Environnement de developpement

* Installer go version 1.18 ou supérieur
* Installer Air pour le hot-reloading du serveur
```
go install github.com/cosmtrek/air@latest
```
* Installer les dépendances du projet (en étant dans le repertoire du projet)
```
go mod download
```

# Build

## Sans docker

Permet plusieurs executable à la fois.

```
./build.sh bin -l <ENVIRONMENT> [-l <AUTREENVIRONMENT>]
```

## Avec docker

```
./build.sh container -l <ENVIRONMENT> [-l <AUTREENVIRONMENT>]
```

# Développement

## Sans hot reloading

ENVIRONMENT étant une étiquette de build pour ne compiler qu'avec le support d'un environnement (langage + paquets optionnels). Par design, un binaire ne peut supporter qu'un langage à la fois.

```
./run.sh bin <ENVIRONMENT>
```

## Avec le hot reloading

```
./run.sh dev <ENVIRONMENT>
```
