package particles

import (
	"project-particles/config"
	"math/rand")

func (s *System) newParticle() {

	//newParticule comporte les caractéristique d'une particule tels que sa position, sa couleur, sa taille afin d'afficher correctement les caractéristiques demandées sur l'interface.
	newParticule := Particle{
		PositionX: float64(config.General.SpawnX),
		PositionY: float64(config.General.SpawnY),

		ScaleX: RandomBetweenFloat(config.General.MinScaleX, config.General.MaxScaleX), 
		ScaleY: RandomBetweenFloat(config.General.MinScaleY, config.General.MaxScaleY),

		ColorRed: RandomBetweenFloat(config.General.MinColorRed, config.General.MaxColorRed),
		ColorGreen: RandomBetweenFloat(config.General.MinColorGreen, config.General.MaxColorGreen),
		ColorBlue: RandomBetweenFloat(config.General.MinColorBlue, config.General.MaxColorBlue),

		Opacity: rand.Float64(),
		Base_Opacity : 0,

		VitesseX: float64(rand.Float64()*config.General.MaxVitesseX-config.General.MaxVitesseX/2),
		VitesseY: float64(rand.Float64()*config.General.MaxVitesseY-config.General.MaxVitesseY/2),

		Lifetime : config.General.Lifetime,

	}
	newParticule.Base_Opacity = newParticule.Opacity

	newParticule.SetSpawn() //CHoix de la forme du générateur de particules

	s.Content.PushFront(&newParticule) // On insère cette nouvelle particule dans la liste du système afin que celui-ci puisse l'afficher.
}








