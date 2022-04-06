<script>
export default {
    props: [],
    data(){
        return{
            time : 1617227200,
            // time : Date.now()
        }
    },
    methods :{
        DateConverter(unixTime){
        let date = new Date(unixTime * 1000);
        let hours = date.getHours();
        let minutes = "0" + date.getMinutes();
        let day = date.getDay();
        let month = date.getMonth();
        let year = date.getFullYear();
        let formattedTime = `${day}/${month}/${year} | ` + hours + ':' + minutes.substr(-2);
        return formattedTime;
        },
    },
}
</script>

<template>

    <div class="nav">
        <div class="flightTypeNav">
            <p class="flightTypePara">
                Départs
            </p>
            <p class="flightTypePara">
                Arrivées
            </p>
        </div>
        <div class="airportName">
            Aéroport Roissy Charles de Gaulle, Paris, France | {{ this.DateConverter(this.time) }}
        </div>        
    </div>
</template>

<style>
    .nav{
        margin: 10px 0;
        display: flex;
        flex-direction: column;
    }

    .flightTypeNav{
        display: flex;
        flex-direction: row;
        justify-content: space-around;
    }

    .flightTypePara{
        border: 1px solid white;
        padding: 5px;
    }

    .airportName{
        font-size: 15px;
    }
</style>