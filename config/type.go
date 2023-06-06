package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              	 		string
	WindowSizeX, WindowSizeY 	 		int
	ParticleImage            	 		string
	Debug                    	 		bool
	InitNumParticles         	 		int
	RotationParticule                   float64
	TypeGenerateur						int
	SpawnX, SpawnY           	 		int
	RayonSpawnX,RayonSpawnY				float64
	MaxVitesseX,MaxVitesseY  	 		float64
	SpawnRate                	 		float64
	GravityX,GravityY        	 		float64
	Kill_particule_WindowSizeX   		float64
	Kill_particule_WindowSizeY   		float64
	OpacityLifetime  		 	 		bool
	Lifetime,ChangeOpacity       		float64 
	MinColorRed,MaxColorRed  	 		float64
	MinColorBlue,MaxColorBlue    		float64
	MinColorGreen,MaxColorGreen  		float64
	MinScaleX, MaxScaleX         		float64
	MinScaleY, MaxScaleY         		float64
	ClickMouseParticules,FollowMouseSpawn	bool
	ClickSpawnRate               		float64
	EffectExplosion              		bool
	Bounce						 		bool
}

var General Config
