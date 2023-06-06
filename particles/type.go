package particles

import "container/list"

// System définit un système de particules.
// Pour le moment il ne contient qu'une liste de particules, mais cela peut
// évoluer durant votre projet.

//Content est la liste du système comportant les particules tandis que SpawnRate est un float correspond au nombre de particule qui apparaitront chaque seconde
type System struct {
	Content *list.List
	SpawnRate float64
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.

//Ici, cette structure comporte les caractéristiques d'une particule.
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity, Base_Opacity           float64
	VitesseX, VitesseY              float64
	Lifetime						float64

}
