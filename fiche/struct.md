# Struct

Les structs sont utiliser pour ranger plusieurs variables liees entre elles. Permet de modeliser des objets ou des concepts concrets.

## Declarer une struct

```go
type Utilisateur struct {
    Nom     string
    Email   string
    Age     int
}
```

`Utilisateur` : struct

Ce struct a 3 champs, et chaque champ a un type.

## Creation

```go
u := Utilisateur{
    Nom: "Alice",
    Email: "alice@example.com",
    Age: 28,
}
```

Le code creer un utilisateur avec ces donnees.

## Acceder aux donnees du struct

```go
fmt.Println(u.Nom)  // Affiche "Alice"
```

## Utilisation

- Organisation des donnees : au lieux d'avoir des variables separes, on les regroupes.
- Lisibilite : il est facile de savoir ce que represente le struct
- Travail avec des fonctions : on peut passer un struct a une fonction pour manipuler tout l'objet
- Ajout de methodes : permet de faire de la POO
