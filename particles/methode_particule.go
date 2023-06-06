package particles

import("project-particles/config")


//Méthode permettant de savoir si une particule se situe dans la zone kill particules grâce aux positions X et Y de la particule.
//On renvoie True si c'est le cas et false si celle-ci est dans une zone autorisée aux particules.
func(p *Particle) OutOfScreen() bool{
	Kill_particule_WindowSizeX := 0.1*float64(config.General.WindowSizeX)
	Kill_particule_WindowSizeY := 0.1*float64(config.General.WindowSizeY)

	if p.PositionX > float64(config.General.WindowSizeX)+Kill_particule_WindowSizeX|| 
	p.PositionX < -Kill_particule_WindowSizeX|| 
	p.PositionY > float64(config.General.WindowSizeY)+Kill_particule_WindowSizeY|| 
	p.PositionY < -Kill_particule_WindowSizeY{
		return true
	}
	return false
}

//Cette fonction reprend la méthode précédente mais sous un autre angle.
// En effet, celle-ci est notamment utilisée pour les positions X et Y du pointeur de la sourie mais également pour les positions de Spawn.
//Ainsi, nous ne prenons pas en compte les positins d'une particule dans cette fonction.
func OutOfKillScreen(x,y float64) bool{
	if x > float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX|| 
	x < -config.General.Kill_particule_WindowSizeX|| 
	y > float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY|| 
	y < -config.General.Kill_particule_WindowSizeY{
		return true
	}
	return false
}


//Cette méthode permet de changer les vitesses des particules en les rendant soit négatives (si elles étaient positives avant) ou inversement.
//Néanmoins, la condition est que celle-ci s'enclenche que si la particule est en contact avec lesbordures de la fenetre ou de la KillParticule.
func (p *Particle)WallBounce(){
	posx,posy := p.PositionX,p.PositionY 
	switch{
	case posx < -config.General.Kill_particule_WindowSizeX || posx> float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX:
		p.VitesseX*=-1
	case posy < -config.General.Kill_particule_WindowSizeY || posy> float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY:
		p.VitesseY*=-1	
	}
}



//On choisit le typedeGénérateur que nous souhaitons parmi les choix proposés.
func (p *Particle) SetSpawn(){

	//Si le TypeGenerateur est -1 depuis config.json alors chaque particule aura une position différente sur les axes X et Y comprise entre 0 et la valeur défini dans WindowSizeX et WindowSizeY
	if config.General.TypeGenerateur == -1{
		p.PositionX = RandomBetweenFloat(-config.General.Kill_particule_WindowSizeX,float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX)
		p.PositionY = RandomBetweenFloat(-config.General.Kill_particule_WindowSizeY,float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY)
	}

	//Si le TypeGenerateur est 1 depuis config.json alors chaque particule aura une position de telle sorte à former un rectangle.
	if config.General.TypeGenerateur == 1{
		haut := float64(config.General.SpawnY)-config.General.RayonSpawnY
		bas := float64(config.General.SpawnY)+config.General.RayonSpawnY
		gauche :=float64(config.General.SpawnX)-config.General.RayonSpawnX
		droite := float64(config.General.SpawnX)+config.General.RayonSpawnX
		
		//Toutes ces conditions permettent de prendre en compte la zone killparticule quand les particules vont former un rectangle.
		if gauche < -config.General.Kill_particule_WindowSizeX{
			gauche = -config.General.Kill_particule_WindowSizeX
		}
		if droite> float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX{
			droite = float64(config.General.WindowSizeX)+config.General.Kill_particule_WindowSizeX
		}
		if haut < -config.General.Kill_particule_WindowSizeY{
			haut = -config.General.Kill_particule_WindowSizeY
		}
		if bas > float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY{
			bas = float64(config.General.WindowSizeY)+config.General.Kill_particule_WindowSizeY
		}

		//En mettant leurs ppositions en aléatoire, on arrive à remplir la figure que nous avons choisi
		p.PositionX =  RandomBetweenFloat(gauche,droite)
		p.PositionY =  RandomBetweenFloat(haut,bas)
	}


	//Si le TypeGenerateur est 1 depuis config.json alors chaque particule aura une position de telle sorte à former un cercle.
	if config.General.TypeGenerateur == 2{
		vecx,vecy:=RandomVecteur(RandomBetweenFloat(0,config.General.RayonSpawnX))
		p.PositionX+=vecx
		p.PositionY+=vecy
	}

}

//Cette méthode permet de choisir la durée de vie en fonction du lifetime ou bien  de l'opacité
func (p *Particle) UpdateOpacity(){

	if config.General.OpacityLifetime && config.General.Lifetime>0{
		p.Opacity-= p.Base_Opacity/config.General.Lifetime
	}else{
		p.Opacity+= config.General.ChangeOpacity
	}
	
}

