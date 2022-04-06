<script>
import Navbar from '../components/Navbar.vue'
import FlightsListArrival from '../components/FlightsListArrival.vue'
import FlightsListStart from '../components/FlightsListStart.vue'

export default {
  name: 'HomeView',
  components: {
    Navbar,
    FlightsListStart,
    FlightsListArrival,
  }
}
</script>


<template>
  <div class="container">
    <div class="NAVBAR">
      <Navbar/>
    </div>
    <div class="FLIGHTS">
      <FlightsListStart class="verticalList"/>
      <FlightsListArrival class="verticalList"/>
    </div>
</div>
</template>

<style scoped>
  .container {  display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 0.2fr 1.8fr;
    gap: 0px 0px;
    grid-auto-flow: row;
    grid-template-areas:
      "NAVBAR"
      "FLIGHTS";
      background-color: #99AAB5;
  }

  .NAVBAR { 
    grid-area: NAVBAR;
    background-color: #23272A;
    color: white;
    font-size: 20px;
    }

  .FLIGHTS { 
    grid-area: FLIGHTS; 
    display: grid;
    grid-template-rows: 1fr;
    grid-template-columns: 1fr 1fr;
    }

    .verticalList{
      display: flex;
      flex-direction: column;
      margin: 5px;
      border-left: 3px solid black;
      border-right: 3px solid black;
      padding : 5px;
      background-color: transparent;
    }

</style>
