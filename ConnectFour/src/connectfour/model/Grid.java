package connectfour.model;

/**
 * Une grille du jeu Connect Four.
 * Elle se caractérise par deux dimensions x et y.
 * A chaque position (x, y) est situé un jeton de type Tokens ou null.
 * 
 * Pour ajouter un jeton, on choisit la couleur (Tokens) et la colonne (x).
 * Le jeton "tombe" alors au plus petit y disponible (en (x,y) == null).
 * Une exception est levée si une colonne déborde, c'est-à-dire si on tente
 *  de mettre un jeton dans une colonne remplie.
 */
public interface Grid {
	
	/**
	 * Récupère le jeton présent en (x,y).
	 * @param x la colonne
	 * @param y la ligne
	 * @return le jeton présent, null si aucun
	 */
	Tokens getToken(int x, int y);
	
	/**
	 * Insère le jeton token dans la colonne x.
	 * Le jeton descend jusqu'à atteindre la dernière ligne y
	 *  où getToken(x, y) == null.
	 * @param token le jeton à insérer
	 * @param x la colonne cible
	 */
	void putToken(Tokens token, int x) throws ConnectException;
	
	/**
	 * Récupère la ligne dans laquelle se trouve le dernier jeton déposé.
	 * @return la ligne du dernier jeton, null si aucun jeton n'a été posé
	 */
	Integer getRowOfLastPutToken();
	
	/**
	 * Initialise la grille, en la vidant de tous ses jetons.
	 */
	void init();

}
