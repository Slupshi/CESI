package connectfour.view;

import java.io.BufferedReader;
import java.io.InputStreamReader;

import connectfour.model.ConnectException;
import connectfour.model.Game;
import connectfour.model.Tokens;

public class ConsoleGameView implements GameView {
	
	// ATTRIBUTES
	
	private final Game game;
	
	private final BufferedReader reader;
	private String userInput;
	private String errorMessage;
	private boolean exitRequest;
	private static final String KEYWORD_EXIT = "E";
	private static final String KEYWORD_RESTART = "R";
	
	// CONSTRUCTOR
	
	public ConsoleGameView(Game game) {
		if (game == null) {
			throw new IllegalArgumentException("game ne peut être null");
		}
		this.game = game;
		reader = new BufferedReader(new InputStreamReader(System.in));
	}
	
	// METHODS

	@Override
	public void play() {
		clearConsole();
		System.out.println("Exit : E, Restart : R");
		displayGrid();
		try {
			do {
				errorMessage = null;
				displayGameState();
				userInput = reader.readLine();
					try {
						switch(userInput){
							case KEYWORD_EXIT:
								this.exitRequest = true;
								break;
							case KEYWORD_RESTART:
								clearConsole();
								this.game.init();
								break;
							default:
								int column = Integer.parseInt(userInput);
								if (!game.isOver())
									game.putToken(column);
								break;
						}
						clearConsole();
					} catch (NumberFormatException e) {
						errorMessage = "Je n'ai pas compris cette réponse...";
					} catch (ConnectException ex){
						errorMessage = ex.getMessage();
					}
				if (errorMessage != null) {
					System.err.println(errorMessage);
				}
			} while (errorMessage != null);
		} catch (Exception e) {
			e.printStackTrace();
			exitRequest = true;
		}
		if (!exitRequest) {
			play();
		}
	}
	
	// TOOLS
	
	private void displayGameState() {
		if (game.isOver()) {
			Tokens winner = game.getWinner();
			if (winner == null) {
				System.out.println("La partie s'est terminée sur un match nul.");
			} else {
				System.out.println("La partie a été remportée par " + winner + " !");
			}
		} else {
			System.out.print("C'est au tour de [" + game.getCurrentPlayer() + "] ! [0-" + (Game.COLUMNS - 1) + "] : ");
		}
	}
	
	private void displayGrid() {
		StringBuffer output = new StringBuffer();
		Tokens token;
		for (int x = 0; x < Game.COLUMNS; x++) {
			output.append("   " + x + "  ");
		}
		output.append('\n');
		for (int y = Game.ROWS - 1; y >= 0; y--) {
			for (int x = 0; x < Game.COLUMNS; x++) {
				output.append("|     ");
			}
			output.append("|\n");
			for (int x = 0; x < Game.COLUMNS; x++) {
				output.append("|  ");
				token = game.getToken(x, y);
				if (token == null) {
					output.append(' ');
				} else {
					output.append(token);
				}
				output.append("  ");
			}
			output.append("|\n");
			for (int x = 0; x < Game.COLUMNS; x++) {
				output.append("|_____");
			}
			output.append("|\n");
		}
		System.out.println(output.toString());
	}
	
	private void clearConsole() {
		for (int i = 0; i < 50; i++) {
			System.out.println();
		}
	}

}
