package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.


//Cette méthode permet d'afficher la particule dans l'interface en mettant correctement les caractéristiques choisis par l'utilisateur.
func (g *game) Draw(screen *ebiten.Image) {
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok{
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}


	//Tout le code qui suit permet de mettre en forme l'interface en y incorporant les différentes extensions mais également leurs emplacements.
	// La boucle permet de faire en sorte de n'afficher qu'un menu à la fois et non les 3 en même temps.
	//Vous pouvez ainsi y retrouver les lettres associées aux différents paramètres !

	if g.cache{
		ebitenutil.DebugPrint(screen, "")
	} else if config.General.Debug {
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Nombre de TPS  : ", ebiten.ActualTPS()), config.General.WindowSizeX-250, 0)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Nombre de FPS  : ", ebiten.ActualFPS()),  config.General.WindowSizeX-250, 15)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombres de particules : ",g.system.Content.Len()),config.General.WindowSizeX-250,30)
		
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Type du générateur : ",config.General.TypeGenerateur , "    (B)"),0,60)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnY des particules : ",config.General.SpawnX, "    (C1)"),0,75)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnY des particules: ",config.General.SpawnY, "    (C2)"),0,90)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("RayonSpawnX : ",config.General.RayonSpawnX, "    (D1)"),0,105)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("RayonSpawnY : ",config.General.RayonSpawnY, "    (D2)"),0,120)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnRate : ",config.General.SpawnRate, "    (E)"),0,135)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ClickMouseParticles : ",config.General.ClickMouseParticules, "    (F)"),0,165)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ClickSpawnRate : ",config.General.ClickSpawnRate, "    (G)"),0,180)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("FollowMouseSpawn : ",config.General.FollowMouseSpawn, "    (I)"),0,195)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Bounce : ",config.General.Bounce, "    (J)"),0,210)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxVitesseX : ",config.General.MaxVitesseX, "    (K1)"),0,240)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxVitesseY : ",config.General.MaxVitesseY, "    (K2)"),0,255)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityX : ",config.General.GravityX, "    (L1)"),0,270)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityY : ",config.General.GravityY, "    (L2)"),0,285)


		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Kill_particule_WindowSizeX : ",config.General.Kill_particule_WindowSizeX, "    (X1)"),0,315)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Kill_particule_WindowSizeY : ",config.General.Kill_particule_WindowSizeY, "    (X2)"),0,330)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Lifetime : ",config.General.Lifetime, "    (N)"),0,360)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("OpacityLifetime : ",config.General.OpacityLifetime, "    (O)"),0,375)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ChangeOpacity : ",config.General.ChangeOpacity, "    (P)"),0,390)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorRed : ",config.General.MinColorRed, "    (Y1)"),0,420)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorRed : ",config.General.MaxColorRed, "    (Y2)"),0,435)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorGreen : ",config.General.MinColorGreen, "    (R1)"),0,465)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorGreen : ",config.General.MaxColorGreen, "    (R2)"),0,480)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinColorBlue : ",config.General.MinColorBlue, "    (S1)"),0,510)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxColorBlue : ",config.General.MaxColorBlue, "    (S2)"),0,525)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinScaleX : ",config.General.MinScaleX, "    (U1)"),0,555)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleX : ",config.General.MaxScaleX, "    (U2)"),0,570)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MinScaleY : ",config.General.MinScaleY, "    (V1)"),0,600)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("MaxScaleY : ",config.General.MaxScaleY, "    (V2)"),0,615)

	}else if g.ReadHelp{

		ebitenutil.DebugPrint(screen,"Le Clic gauche de la souris vous permet de faire spawner des particules en fonction du ClickSpawnRate ")
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Actuellement, lorsque vous cliquez / maintenez le clic gauche, vous faites apparaitre ", config.General.ClickSpawnRate, " particules par clic"), 0, 20)
		ebitenutil.DebugPrintAt(screen,"Le Clic droit de la souris vous permet de faire changer le point de spawn des particules en fonction des coordonnées de la souris ", 0, 60)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Actuellement, votre SpawnX de particules est situé à : ", config.General.SpawnX, " pixels"), 0, 80)
		ebitenutil.DebugPrintAt(screen,fmt.Sprint("Actuellement, votre SpawnY de particules est situé à : ", config.General.SpawnY, " pixels"), 0, 100)

		ebitenutil.DebugPrintAt(screen,"Pour ajouter une lettre dans la phrase : 'Entrez une lettre pour ...', il suffit simplement d'appuyer sur la lettre du clavier pour que celle-ci apparaisse", 0, 140)
		ebitenutil.DebugPrintAt(screen,"Pour enlever une lettre dans la phrase : 'Entrez une lettre pour ...', il suffit simplement d'appuyer sur la KeyBackspace situé à gauche de Ver Num pour l'enlever", 0, 160)

		ebitenutil.DebugPrintAt(screen,"la touche Espace vous permet de cacher le menu de configuration une fois dedans ", 0, 190)
		
		ebitenutil.DebugPrintAt(screen,	"Si vous mettez une valeur négative dans les champs Kill_particule alors la zone de propagation de particules diminuera en fonction de celle-ci", 0, 210)

	}else{
		ebitenutil.DebugPrint(screen,"Appuyez sur Echap pour voir/retirer le menu")
		ebitenutil.DebugPrintAt(screen,"Appuyez sur H pour afficher l'aide ainsi que le menu des commandes", 0, 30)
		ebitenutil.DebugPrintAt(screen,"Vous trouverez également des commentaires pour certaines propriétés directement dans le config.json", 0, 60)
		ebitenutil.DebugPrintAt(screen,"Amusez-vous bien !", 0, 75)
	}

	if g.ReadMode{
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Entrez la lettre associée au paramètre que vous souhaitez modifier : ",g.Value),config.General.WindowSizeX/2-200,30)
	}


}