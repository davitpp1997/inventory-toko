<template>
  <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/barang">Data Barang</router-link>
  </div> -->
  <router-view />
</template>

<script>
export default {
  name: 'App',
}
</script>

<style>
@import 
url(https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css);

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>