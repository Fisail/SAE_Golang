package particles

import (
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
)
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.


func (s *System) Update() {

	//Si le FollowMouse est activé alors le point de spawn des particules sera défini en fonctions des positions X et Y du curseur de la souris.
	if config.General.FollowMouseSpawn{
		config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
	}

	//Ce bout de code vous permet de faire apparaître un certain nombre de particules (ClickSpawnRate) grâce au clic gauche de la souris
	//Celui-ci est actif que si vous avez ms true l'extension "ClickMouseParticules" et que les positions du cursor de la souris se situent bien dans la zone "vivable"
	posx,posy :=ebiten.CursorPosition()
	oldspawnx,oldspawny:=config.General.SpawnX,config.General.SpawnY
	if config.General.ClickMouseParticules && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !OutOfKillScreen(float64(posx),float64(posy)){
		for i:=0.0;i<config.General.ClickSpawnRate;i++{
			config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
			s.newParticle()
		}
		config.General.SpawnX,config.General.SpawnY=oldspawnx,oldspawny
	}
	 


	//Si le clic droit de la suouris est appuyée alors le point de spawn des particules changera en fonction des positions du curseur de la souris.
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !OutOfKillScreen(float64(posx),float64(posy)){
		config.General.SpawnX,config.General.SpawnY= ebiten.CursorPosition()
	}



	//Chaque particule du système est transférée dans une variable "particule_individuelle".
	//Dans celle-ci, nous ajoutons une vitesse aléatoire (voir particle.go) à la position de la particule située dans "particule_individuelle".
	//Sachant que la vitesse est aléatoire, chaque particule comportera une position différente en fonction de la valeur de leur vitesseX et VitesseY.
	for element := s.Content.Front() ; element != nil; {

		p, _ := element.Value.(*Particle)

		//Si l'extension "Bounce" est activée alors cela déclenchera cette méthode qui inverse les vitesses.
		if config.General.Bounce{
			p.WallBounce()
		}

		//Nous ajoutons la vitesseX et vitesseY avec les positions de la particule(catactérisée par element) afin que celle-ci bouge.
		p.PositionX += p.VitesseX
		p.PositionY += p.VitesseY

		//Nous ajoutons la GravitéX et GravitéY avec les vitesses de la particule(catactérisée par element) afin que celle-ci s'oriente en fonction de la vitesse.
		p.VitesseX+= config.General.GravityX 
		p.VitesseY+= config.General.GravityY

		//La particule perd de sa vie à chaque appel de la fonction update.
		//La fonction UpdateOpacity permet de choisir la durée de vie entre le lifetime et l'opacité (en fonction du choix de l'utilisateur)
		p.Lifetime--
		p.UpdateOpacity()

		next := element.Next()

		
		// Si la particule est en dehors de la killzoneparticule  ou que celle-ci comporte une durée de vie initiale à 0 alors celle-ci sera supprimée de la liste du ssystème.
		//Cette suppression vaut également si la particule n'a plus de durée de vie ou si son opacité est à zero.
		//Ceci correspond ainsi à l'optimisation mémoire de notre projet.
		if (element.Value.(*Particle)).OutOfScreen()|| p.Lifetime <= 0 && config.General.Lifetime>0 || p.Opacity<=0{
			s.Content.Remove(element)
		}
		element = next
	}



	
	//Ce bout de code permet de générer un nombre de particules, défini par SpawRate dans config.json, par seconde.
	//La boucle permet ainsi de créer ce nombre de particule grâce à la méthode newParticle (présente dans particle.go)
	//Toutefois, afin d'éviter qu'une particule se situé dans une zone killparticule , leOutofKillScreen doit renvoyer une valeur false.
	s.SpawnRate+=config.General.SpawnRate
	if !OutOfKillScreen(float64(config.General.SpawnX),float64(config.General.SpawnY)){
		for i:=0; i< int(s.SpawnRate); i++{
			s.newParticle()
		}
	}
	s.SpawnRate-=float64(int(s.SpawnRate))

	
}


