package connectfour.model;

import java.util.ArrayList;
import java.util.List;

public class ImplListGrid implements Grid{

    public final List<List<Tokens>> grid;
    private Integer rowOfLastPutToken;
    private int rows;
    private int columns;

    public ImplListGrid(int nbCol, int nbLine){
        this.rows = nbLine;
        this.columns = nbCol;
        this.grid = new ArrayList<>();
        for(int i = 0; i < nbCol; i++){
            this.grid.add(new ArrayList<Tokens>());
        }
        this.init();
    }
    @Override
    public Tokens getToken(int x, int y) {

        if(x < 0 || y < 0 || x >= this.columns || y >= this.grid.get(x).size()){
            return null;
        }
        return this.grid.get(x).get(y);
    }

    @Override
    public void putToken(Tokens token, int x) throws ConnectException {
        if(x < 0 || x > this.columns - 1){
            throw new ConnectException("Ce chiffre ne correspond à aucune colonne");
        } else if (this.grid.get(x).size() >= this.rows) {
            throw new ConnectException("La colonne est pleine");
        }
        this.grid.get(x).add(token);
        rowOfLastPutToken = this.grid.get(x).size() - 1;
    }

    @Override
    public Integer getRowOfLastPutToken() {
        return rowOfLastPutToken;
    }

    @Override
    public void init() {
        rowOfLastPutToken = null;
    }
}
