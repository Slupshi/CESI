//#region Crepes
const DefaultIngredients = {flour: 63, eggs : 1, milk : 0.2, sugar : 0.5, butter : 13};
const Ingredients = [flour = 63, eggs = 1, milk = 0.2, sugar = 0.5, butter = 13];

function Crepes()
{
    let finalIngredient= []; 
    let invite = document.getElementById("peopleNumberInput").value;    
    Ingredients.forEach(
    function Calcul(value){
        value = value*invite;
        finalIngredient.push(value);    
    });

    let sCrepe = " Crêpe";
    switch(invite){
        case "" :
            sCrepe = "0 Crêpe";
            break;
        case "0":
            sCrepe = " Crêpe";
            break;
        case "1":
            sCrepe = " Crêpe";
            break;
        default:
            sCrepe = " Crêpes";
            break;
    };

    let sOeuf = " Crêpe";
    switch(finalIngredient[1]){        
        case 0:
            sOeuf = '';
            break;
        case 1:
            sOeuf = '';
            break;
        default:
            sOeuf = 's';
            break;
    };

    document.getElementById("flourLabel").innerText = "Farine : " + finalIngredient[0] + " g";
    document.getElementById("eggLabel").innerText = "Oeuf : " + finalIngredient[1] + " unité" + sOeuf;
    document.getElementById("milkLabel").innerText = "Lait : " + finalIngredient[2].toFixed(1) + " L";
    document.getElementById("sugarLabel").innerText = "Sucre : " + finalIngredient[3] + " c à s";
    document.getElementById("butterLabel").innerText = "Beurre : " + finalIngredient[4] + " g";
    document.getElementById("lastIngredientLabel").innerText = "Pour " + invite + sCrepe;

    // console.info("Farine : " + finalIngredient[0] + " g");
    // console.info("Oeuf(s) : " + finalIngredient[1] + " unité(s)");
    // console.info("Lait : " + finalIngredient[2].toFixed(1) + " L");
    // console.info("Sucre : " + finalIngredient[3] + " c à s");
    // console.info("Beurre : " + finalIngredient[4] + " g");
}


function InputUp()
{
    document.getElementById("peopleNumberInput").value++;  
}

function InputDown(){
    if(document.getElementById("peopleNumberInput").value > 0)
    {
        document.getElementById("peopleNumberInput").value--;  
    }
}

//#endregion Crepes