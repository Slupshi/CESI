package connectfour.model;

public class ImplGame implements Game{

    private Grid grid;
    private Tokens currentPlayer;
    private boolean isOver;
    private Tokens winner;
    private static final Tokens[] TOKENS_VALUES = Tokens.values();

    public ImplGame(){
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
        return isOver;
    }

    @Override
    public Tokens getWinner() {
        return this.winner;
    }

    @Override
    public void putToken(int column)  throws ConnectException{
        this.grid.putToken(this.currentPlayer, column);
        this.isOver = this.calculateOver(column);
        if(!isOver){
            this.getNextPlayer();
        }
    }

    @Override
    public void init() {
        this.grid = new ImplListGrid(Game.COLUMNS, Game.ROWS);
        this.currentPlayer = TOKENS_VALUES[(int)(Math.random()*TOKENS_VALUES.length)];
        this.isOver = false;
        this.winner = null;
    }

    public void getNextPlayer(){
        this.currentPlayer = TOKENS_VALUES[(this.currentPlayer.ordinal()+1) % TOKENS_VALUES.length];
    }

    public boolean calculateOver(int x){
        if(inspectNWSE(x, this.grid.getRowOfLastPutToken()) >= Game.REQUIRED_TOKENS){
            this.winner = currentPlayer;
            return true;
        } else if (inspectWestEast(x, grid.getRowOfLastPutToken()) >= Game.REQUIRED_TOKENS) {
            this.winner = currentPlayer;
            return true;
        } else if (inspectNESW(x, grid.getRowOfLastPutToken()) >= Game.REQUIRED_TOKENS) {
            this.winner = currentPlayer;
            return true;
        } else if (inspectSouth(x, grid.getRowOfLastPutToken()) >= Game.REQUIRED_TOKENS) {
            this.winner = currentPlayer;
            return true;
        }
        for(int i=0; i< Game.COLUMNS; i++)
        {
            if(this.grid.getToken(i, Game.ROWS - 1) == null){
                return false;
            }
        }
        return true;
    }

    private int inspectSouth(int x, int y){
        int foundInLine = 0;
        for(int i = 1; y - i >= 0 && getToken(x, y - i) == currentPlayer; i++){
            foundInLine++;
        }
        return foundInLine + 1;
    }

    private int inspectNWSE(int x, int y){
        int foundInLine = 0;
        for (int i = 1; x - i >= 0 && y + i < ROWS && getToken(x - i, y + i) == currentPlayer; i++) {
            foundInLine++;
        }
        for (int i = 1; x + i < COLUMNS && y - i >= 0 && getToken(x + i, y - i) == currentPlayer; i++) {
            foundInLine++;
        }
        return foundInLine + 1;
    }

    private int inspectWestEast(int x, int y){
        int foundInLine = 0;
        for (int i = 1; x - i >= 0 && getToken(x - i, y) == currentPlayer; i++) {
            foundInLine++;
        }
        for (int i = 1; x + i < COLUMNS && getToken(x + i, y) == currentPlayer; i++) {
            foundInLine++;
        }
        return foundInLine + 1;
    }

    private int inspectNESW(int x, int y){
        int foundInLine = 0;
        for (int i = 1; x - i >= 0 && y - i >= 0 && getToken(x - i, y - i) == currentPlayer; i++) {
            foundInLine++;
        }
        for (int i = 1; x + i < COLUMNS && y + i < ROWS && getToken(x + i, y + i) == currentPlayer; i++) {
            foundInLine++;
        }
        return foundInLine + 1;
    }

}
