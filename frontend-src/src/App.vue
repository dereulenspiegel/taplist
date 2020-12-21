<template>
  <section class="hero is-fullheight">
    <div class="hero-head">
      <h3 class="title">Kegerator</h3>
    </div>
    <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-evenly is-align-content-space-evenly hero-body">
      <BeerTap v-for="beer in beers" :key="beer.id" v-bind:beer="beer"/>
    </div>
    <div class="hero-foot">
      <p>GitHub: dereulenspiegel/taplist</p>
    </div>
  </section>
</template>

<script>
import BeerTap from './components/BeerTap.vue'

export default {
  name: 'App',
  components: {
    BeerTap
  },
  data: function(){
    return {
      beers: [
        {
          id: "1",
          name: 'Beer 1'
        },
        {
          id: "2",
          name: 'Beer 2'
        },
        {
          id: "3",
          name: 'Beer 3'
        },
        {
          id: "4",
          name: 'Beer 4'
        },
        {
          id: "1",
          name: 'Beer 1'
        },
        {
          id: "2",
          name: 'Beer 2'
        },
        {
          id: "3",
          name: 'Beer 3'
        },
        {
          id: "4",
          name: 'Beer 4'
        },
      ]
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
