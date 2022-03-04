//#region Crepes
const DefaultIngredients = {flour: 63, eggs : 1, milk : 0.2, sugar : 0.5, butter : 13};
const Ingredients = [flour = 63, eggs = 1, milk = 0.2, sugar = 0.5, butter = 13];

function Crepes(){
    let finalIngredient= []; 
    let invite = document.getElementById("peopleNumberInput").value;    
    Ingredients.forEach(
    function Calcul(value){
        value = value*invite;
        finalIngredient.push(value);    
    });

    document.getElementById("flourLabel").innerText = "Farine : " + finalIngredient[0] + " g"
    document.getElementById("eggLabel").innerText = "Oeuf : " + finalIngredient[1] + " unité(s)"
    document.getElementById("milkLabel").innerText = "Lait : " + finalIngredient[2].toFixed(1) + " L"
    document.getElementById("sugarLabel").innerText = "Sucre : " + finalIngredient[3] + " c à s"
    document.getElementById("butterLabel").innerText = "Beurre : " + finalIngredient[4] + " g"

    // console.info("Farine : " + finalIngredient[0] + " g");
    // console.info("Oeuf(s) : " + finalIngredient[1] + " unité(s)");
    // console.info("Lait : " + finalIngredient[2].toFixed(1) + " L");
    // console.info("Sucre : " + finalIngredient[3] + " c à s");
    // console.info("Beurre : " + finalIngredient[4] + " g");
}
//#endregion Crepes