# Service d'Exécution

## Compilation manuel

```bash
EXEC_LANG=python go build -o srvexec_$EXEC_LANG -tags $EXEC_LANG .
```

## Création de l'image

Il faut d'abord compiler le binaire d'API pour le language de votre choix.

Puis, l'importer dans l'image du language.

```bash
./build.sh -l python
```

## Lancer le conteneur

```bash
docker run --rm --name srvexec-python -p 8080:8080 srvexec:python
```

## Tester

Le header `Content-Type` est important.

```bash
curl -X POST http://localhost:8080/exec -d '{"code": "def ...", "context": "import math"}' -H 'Content-Type: application/json'
```

