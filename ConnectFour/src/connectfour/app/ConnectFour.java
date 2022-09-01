package connectfour.app;

import connectfour.model.ImplGame;
import connectfour.view.ConsoleGameView;
import connectfour.view.GameView;

public class ConnectFour {
	
	public static void main(String[] args) {
		
		// Création du jeu
		// On instancie une vue, à laquelle on fournit un modèle (MVC)
		GameView gameView = new ConsoleGameView(new ImplGame());
		
		// Lancement du système de jeu
		gameView.play();
		
		// Fin de l'application
		System.exit(1);
		
	}

}
