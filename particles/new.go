package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.




func NewSystem() System {

	s := System{Content: list.New()}
	rand.Seed(time.Now().UnixNano()) //Ceci crée une valeur en fonction du temps. Sachant qu'une durée est unique par jour, la valeur la sera également et nous obtiendrons une valeur aléatoire non-répétitive.


	//Cette boucle permet d'afficher le nombre de particules renseigné à InitNumParticles dans config.json directement sur l'interface.
	for i:=0; i < config.General.InitNumParticles; i++{
		s.newParticle()
	}
	return s// On retourne le système et plus précisemment sa liste afin d'obtenir la liste comportant chaque particle.
}

