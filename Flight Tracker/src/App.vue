<script>
  import HomeView from './views/HomeView.vue' 

  export default {
  name: 'App',
  components: {
    HomeView,
  }
}
</script>

<template>
  <div>
    <HomeView/>
  </div>
</template>

<style>
body{
    margin: 0;
    padding: 0;
}

::-webkit-scrollbar-track {
    box-shadow: inset 0 0 5px #23272A;
    border-radius: 10px;
}

::-webkit-scrollbar {
    width: 10px;
}

::-webkit-scrollbar-thumb {
    background: #23272A;
    border-radius: 10px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  font-weight: bolder;
  font-size: 20;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
