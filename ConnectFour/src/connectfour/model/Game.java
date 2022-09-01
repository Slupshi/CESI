package connectfour.model;

/**
 * Mécanique du jeu Connect Four.
 * Chaque joueur pose à son tour un jeton dans une colonne de la grille.
 * Le jeu s'arrête lorsque quatre jetons de même couleur
 *  sont alignés en ligne, en colonne ou en diagonale
 *  ou lorsque la grille est remplie (auquel cas il y a match nul).
 * A la création, le jeu est initialisé.
 */
public interface Game {
	
	/** Nombre de colonnes dans la grille */
	static final int COLUMNS = 7;
	
	/** Nombre de lignes dans la grille */
	static final int ROWS = 6;
	
	/** Nombre de jetons à aligner pour gagner */
	static final int REQUIRED_TOKENS = 4;
	
	/**
	 * Récupère le jeton présent en (x,y).
	 * @param x la colonne
	 * @param y la ligne
	 * @return le jeton présent, null si aucun
	 */
	Tokens getToken(int x, int y);
	
	/**
	 * Récupère le joueur dont c'est le tour de jouer.
	 * @return le jeton du joueur qui doit jouer
	 */
	Tokens getCurrentPlayer();
	
	/**
	 * Indique si le jeu est terminé ou non.
	 * @return true si l'un des joueurs a gagné ou si la grille est complète
	 */
	boolean isOver();
	
	/**
	 * Récupère le vainqueur de la partie.
	 * @return le jeton du joueur qui a remporté la partie
	 *  si elle est terminée et qu'il n'y a pas eu de match nul.
	 */
	Tokens getWinner();
	
	/**
	 * Insère un jeton pour le joueur courant dans la colonne x.
	 * @param column la colonne cible
	 */
	void putToken(int column) throws ConnectException;
	
	/**
	 * Initialise le jeu.
	 * Vide la grille et tire au sort le joueur qui commence.
	 */
	void init();

}
