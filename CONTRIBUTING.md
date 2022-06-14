# Environnement

* Installer go version 1.18 ou supérieur
* Installer Air pour le hot-reloading du serveur
```
go install github.com/cosmtrek/air@latest
```
* Installer les dépendances du projet
```
go get
```

# Développement

## Sans hot reloading

LANGAGE étant une étiquette de build pour ne compiler qu'aver le support d'un environnement (language + paquets optionnels). Par design, un binaire ne peut supporter qu'un langage à la fois.

```
./run.sh bin <LANGAGE>
```

## Avec le hot reloading

```
./run.sh dev <LANGAGE>
```

# Build

## Sans docker

Permet plusieurs executable à la fois.

```
./build.sh bin -l <LANGAGE> [-l <AUTRELANGAGE>]
```

## Avec docker

```
./build.sh container -l <LANGAGE> [-l <AUTRELANGAGE>]
```