package connectfour.model;

import java.util.HashMap;
import java.util.Map;

public class ImplGrid implements Grid{

    public final Tokens[][] grid;
    private Integer rowOfLastPutToken;

    public ImplGrid(int nbLine, int nbCol){
        if(nbLine <0 || nbCol <0){
            throw new IllegalArgumentException();
        }
        grid = new Tokens[nbCol][nbLine];
        this.init();
    }

    @Override
    public Tokens getToken(int x, int y) {
        return grid[x][y];
    }

    @Override
    public void putToken(Tokens token, int x)  throws ConnectException{
        if(x <0 || x > Game.COLUMNS - 1){
            throw new ConnectException("Ce chiffre ne correspond à aucune colonne");
        }
        int y = 0;
        while (y < grid[x].length && grid[x][y] != null) {
            y++;
        }
        grid[x][y] = token;
        rowOfLastPutToken = y;
    }

    @Override
    public Integer getRowOfLastPutToken() {
        return rowOfLastPutToken;
    }

    @Override
    public void init() {
        this.rowOfLastPutToken = null;
    }
}
