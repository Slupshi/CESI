<script>
export default {
    props: ["Arrival", "Airport",]
}
</script>

<template>
    <div class="flightItem">
        <p> {{ Airport }} | Arrivée prévue pour : {{ Arrival }} </p>
    </div>
</template>

<style>
    .flightItem{
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        border: 1px solid white;
        border-radius: 1rem;
        background-color: #2C2F33;
        height: 40px;
        color: white;
    }
</style>