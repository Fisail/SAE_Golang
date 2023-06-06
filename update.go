package main

import("github.com/hajimehoshi/ebiten/v2/inpututil"
		"github.com/hajimehoshi/ebiten/v2"
		"project-particles/config"
		)
// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	g.system.Update()

	//Si la touche Echap est appuyé alors le menu des configuration apparaitra, vous permettant ainsi de mchoisir ce que vous souhaitez mettre ou enlever.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape){
		g.ReadMode = true
		config.General.Debug = !config.General.Debug
		g.ReadHelp = false
	}
	//Si la touche H est appuyé alors le menu d'aide ainsi que des commandes apparaitra.
	//Cela cahcera également le menu de configuration ainsi que la page d'accueil .
	if inpututil.IsKeyJustPressed(ebiten.KeyH){
		g.ReadHelp = !g.ReadHelp
		g.ReadMode = false
		config.General.Debug = false
	}

	//Si la touche Espacee est appuyé alors le menu des configuration sera caché, vous permettant ainsi de mieux voir le résultat de vos particules.
	//Réappuyez dessus et le menu sera de nouveau visible.
	if inpututil.IsKeyJustPressed(ebiten.KeySpace){
		g.cache = !g.cache
		if g.cache{
			g.ReadMode = false
		} else {
			g.ReadMode = true
		}
	}

	g.Ajout_Lettre_Champs()
	
	return nil
}


//Cette méthode permet de rendre les lettres que vous atperez visible dans le champs de texte du choix d'une lettre
//Celle-ci vous donne également la possibilité de supprimer la lettre si vous appuyez sur la touche défini et qu'au moins une lettre soit affiché.
func (g *game) Ecriture_Lettre()string{
	if g.ReadMode{
		g.Value += NumPressed()
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(g.Value)>0{
			g.Value = g.Value[:len(g.Value)-1]
		}
	}
	return g.Value
}

//Cette méthode permet de pouvoir changer les paramètres que l'on souhaite en fonction de la lettre insérée.
//Chaque fonctionnalité comporte une lettre qui lui est associée (elles sont à côté de celles-ci).
func (g*game) Ajout_Lettre_Champs(){
	g.Ecriture_Lettre()
	switch{

	//Cette confition comporte une limite car nous ne voulons afficher que des chiffres compris entre -1 et 2.
	case g.Value == "B":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			if config.General.TypeGenerateur < 2{
				config.General.TypeGenerateur += 1
			}
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			if config.General.TypeGenerateur > -1{
				config.General.TypeGenerateur -= 1
			}
		}

	case g.Value == "C1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.SpawnX+= 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.SpawnX -= 1
		}

	case g.Value == "C2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.SpawnY+= 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.SpawnY -= 1
		}


	case g.Value == "D1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.RayonSpawnX += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.RayonSpawnX -= 0.1
		}


	case g.Value == "D2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.RayonSpawnY += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.RayonSpawnY -= 0.1
		}

	case g.Value == "E":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.SpawnRate += 1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.SpawnRate -= 1
		}



	case g.Value == "F":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.ClickMouseParticules = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.ClickMouseParticules = false
		}

	case g.Value == "G":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.ClickSpawnRate += 0.1
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.ClickSpawnRate -= 0.1
		}

	case g.Value == "I":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.FollowMouseSpawn = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.FollowMouseSpawn = false
		}

	case g.Value == "J":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Bounce = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Bounce = false
		}

	case g.Value == "K1":
		if ebiten.IsKeyPressed(ebiten.KeyKPAdd){
			config.General.MaxVitesseX += 0.2
		}else if ebiten.IsKeyPressed(ebiten.KeyKPSubtract){
			config.General.MaxVitesseX -= 0.2
		}

	case g.Value == "K2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxVitesseY += 0.2
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MaxVitesseY -= 0.2
		}
	case g.Value == "L1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.GravityX += 0.2
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.GravityX -= 0.2
		}

	case g.Value == "L2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.GravityY += 0.2
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.GravityY -= 0.2
		}

	case g.Value == "X1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Kill_particule_WindowSizeX += 5
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Kill_particule_WindowSizeX -=  5
		}

	case g.Value == "X2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Kill_particule_WindowSizeY += 5
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Kill_particule_WindowSizeY -=  5
		}

	case g.Value == "N":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.Lifetime += 10
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.Lifetime -= 10
		}

	case g.Value == "O":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.OpacityLifetime = true
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.OpacityLifetime = false
		}

	case g.Value == "P":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.ChangeOpacity += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.ChangeOpacity -= 0.1
		}

	case g.Value == "Y1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorRed += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MinColorRed -= 0.1
		}

	case g.Value == "Y2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorRed += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MaxColorRed -= 0.1
		}

	case g.Value == "R1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorGreen += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MinColorGreen -= 0.1
		}

	case g.Value == "R2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorGreen += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MaxColorGreen -= 0.1
		}

	case g.Value == "S1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinColorBlue += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MinColorBlue -= 0.1
		}

	case g.Value == "S2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxColorBlue += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MaxColorBlue -= 0.1
		}

	case g.Value == "U1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinScaleX += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MinScaleX -= 0.1
		}

	case g.Value == "U2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MaxScaleX += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MaxScaleX -= 0.1
		}

	case g.Value == "V1":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
			config.General.MinScaleY += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MinScaleY -= 0.1
		}

	case g.Value == "V2":
		if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd){
				config.General.MaxScaleY += 0.1
		}else if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract){
			config.General.MaxScaleY -= 0.1
		}
	}
}
	


//Cette fonction permet de renvoyer un string de la lettre que vous avez choisi dans g.Value(champs dans lequel vous choississez votre lettre)
func NumPressed() string{
	switch{
	case inpututil.IsKeyJustPressed(ebiten.KeyB):
		return "B"
	case inpututil.IsKeyJustPressed(ebiten.KeyC):
		return "C"
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		return "D"
	case inpututil.IsKeyJustPressed(ebiten.KeyE):
		return "E"
	case inpututil.IsKeyJustPressed(ebiten.KeyF):
		return "F"
	case inpututil.IsKeyJustPressed(ebiten.KeyG):
		return "G"
	case inpututil.IsKeyJustPressed(ebiten.KeyH):
		return "H"
	case inpututil.IsKeyJustPressed(ebiten.KeyI):
		return "I"
	case inpututil.IsKeyJustPressed(ebiten.KeyJ):
		return "J"
	case inpututil.IsKeyJustPressed(ebiten.KeyK):
		return "K"
	case inpututil.IsKeyJustPressed(ebiten.KeyL):
		return "L"
	case inpututil.IsKeyJustPressed(ebiten.KeyX):
		return "X"
	case inpututil.IsKeyJustPressed(ebiten.KeyN):
		return "N"
	case inpututil.IsKeyJustPressed(ebiten.KeyO):
		return "O"
	case inpututil.IsKeyJustPressed(ebiten.KeyP):
		return "P"
	case inpututil.IsKeyJustPressed(ebiten.KeyY):
		return "Y"
	case inpututil.IsKeyJustPressed(ebiten.KeyR):
		return "R"
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		return "S"
	case inpututil.IsKeyJustPressed(ebiten.KeyT):
		return "T"
	case inpututil.IsKeyJustPressed(ebiten.KeyU):
		return "U"
	case inpututil.IsKeyJustPressed(ebiten.KeyV):
		return "V"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad1):
		return "1"
	case inpututil.IsKeyJustPressed(ebiten.KeyNumpad2):
		return "2"
	}
	return ""
}

