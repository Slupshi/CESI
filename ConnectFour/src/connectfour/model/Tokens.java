package connectfour.model;

public enum Tokens {
	RED('X'), BLUE('O');
	
	private final char symbol;
	
	private Tokens(char symbol) {
		this.symbol = symbol;
	}
	
	@Override
	public String toString() {
		return "" + symbol;
	}
}
