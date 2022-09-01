package connectfour.model;

public class ImplGame implements Game{

    private final Grid grid;
    private Tokens currentPlayer;
    private static final Tokens[] TOKENS_VALUES = Tokens.values();

    public ImplGame(){
        this.grid = new ImplGrid(Game.ROWS,Game.COLUMNS);
        this.init();
    }
    @Override
    public Tokens getToken(int x, int y) {
        return this.grid.getToken(x, y);
    }

    @Override
    public Tokens getCurrentPlayer() {
        return this.currentPlayer;
    }

    @Override
    public boolean isOver() {
        return false;
    }

    @Override
    public Tokens getWinner() {
        return null;
    }

    @Override
    public void putToken(int column)  throws ConnectException{
        this.grid.putToken(this.currentPlayer, column);
        this.getNextPlayer();
    }

    @Override
    public void init() {
        this.currentPlayer = TOKENS_VALUES[(int)(Math.random()*TOKENS_VALUES.length)];
    }

    public void getNextPlayer(){
        this.currentPlayer = TOKENS_VALUES[(this.currentPlayer.ordinal()+1) % TOKENS_VALUES.length];
    }
}
