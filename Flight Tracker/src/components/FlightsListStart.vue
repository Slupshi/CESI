<script>
import FlightsListStartItem from './FlightsListStartItem.vue'

export default {
    name : "FlightsList",
    components: {
    FlightsListStartItem,
    }, 
    data(){
        return {
            arrivals : [],
        }
    },
    methods : {
        async SetArrivals(){
            this.arrivals = await GetArrivals();
        },

        DateConverter(unixTime){
        let date = new Date(unixTime * 1000);
        let hours = date.getHours();
        let minutes = "0" + date.getMinutes();
        let formattedTime = hours + ':' + minutes.substr(-2);
        return formattedTime;
        },

        async GetAirportName(icao){
            // Cadeau de l'api_key
            let api_key = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI0IiwianRpIjoiMjZlM2FjYmNlMjhmZDNkOTFiYTRiNzM3NGE2NjBjODc1YTNkMTIzMTYyMzZmOTIxYzlkYTdiMmQzMTczMGNhZTY1YzQ5YmZiZTM2OWE3NGEiLCJpYXQiOjE2NDkyODAzNjIsIm5iZiI6MTY0OTI4MDM2MiwiZXhwIjoxNjgwODE2MzYxLCJzdWIiOiIyMzA0Iiwic2NvcGVzIjpbXX0.KZjgGx7R5nTWV_hd2n51GIOG-97jfk-U1hBluTsT0cIsVMGdVV3Bnp3aTrRKYIEcjeANoUoTcs0jgUaB5VDuuw";
            let url = `https://app.goflightlabs.com/airports?access_key=${api_key}&search=${icao}`;
            let response = await fetch(url, {mode : 'no-cors'});
            let data = await response.json();
            
            console.log(data);
        },

    },
    created(){
        this.SetArrivals();
    }

}

async function GetArrivals(){   
    let begin = 1617227200; // Date.now() return une 404 donc on met une date aléatoire
    let end = 1617230800;
    let pw = "pcZd2ZQMv4NUMP@";
    let username = "Slupshi";
    let icao = "LFPG"; // Roissy Charles de Gaulle
    var headers = new Headers();
    headers.append('Authorization', 'Basic ' + btoa(username + ':' + pw));
    let url = `https://opensky-network.org/api/flights/departure?airport=${icao}&begin=${begin}&end=${end}`; 
    let response = await fetch(url, {headers : headers});  
    let data = await response.json();
    return data;
}

</script>

<template>
    <div>
        <p v-for="item in arrivals" :key="item.icao24">
                <!-- Tentative désespéré de récupérer le nom de la ville de l'aéroport -->
                <!-- <FlightsListStartItem :Start="this.DateConverter(item.firstSeen)" :Airport="this.GetAirportName(item.estArrivalAirport)"/> -->
                <FlightsListStartItem :Start="this.DateConverter(item.firstSeen)" :Airport="item.estArrivalAirport"/>
        </p> 
    </div>
</template>

<style>

</style>