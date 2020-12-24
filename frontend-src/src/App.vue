<template>
  <section class="hero is-fullheight">
    
    <div class="hero-head">
      <h3 class="title">{{kegerator.name}}</h3>
    </div>
    <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-evenly is-align-content-stretch hero-body">
      <BeerTap v-for="tap in kegerator.taps" v-bind:key="tap.id" v-bind:tap="tap" @tapUpdated="refreshData"/>
    </div>
    <div class="hero-foot">
      <p>GitHub: dereulenspiegel/taplist</p>
    </div>
  </section>
</template>

<script>
import BeerTap from './components/BeerTap.vue'

import KEGERATOR_QUERY from './gql/kegerator.gql'

export default {
  name: 'App',
  apollo: {
    kegerator: {
      query: KEGERATOR_QUERY,
      pollInterval: 1000,
    }
  },
  components: {
    BeerTap
  },
  data: function(){
    return {
      kegerator: {
        name: 'Unknown',
        taps: []
      }
    }
  },
  methods: {
    refreshData: function() {
      this.$apollo.query({
        query: KEGERATOR_QUERY,
        update: data => this.kegerator = data
      })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  /*-moz-osx-font-smoothing: grayscale;*/
  text-align: center;
  color: white;
}
</style>
