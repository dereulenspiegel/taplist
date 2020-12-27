<template>
  <section class="hero is-fullheight">
    
    <div class="hero-head">
      <h3 class="title">{{kegerator.name}}</h3>
    </div>
    <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-evenly is-align-content-stretch hero-body">
      <BeerTap v-for="tap in kegerator.taps" v-bind:key="tap.id" v-bind:tap="tap" @tapUpdated="refreshData"/>
    </div>
    <div class="hero-foot">
      <nav class="level is-mobile">
        <div class="level-left">
          <span class="level-item">GitHub: dereulenspiegel/taplist</span>
        </div>

        <div class="level-right">
          <button v-if="!isFullscreen" v-on:click.prevent="enableFullscreen" class="level-item button is-light">
            <span class="icon has-text-white is-large"><font-awesome-icon :icon="['fas', 'expand-arrows-alt']"></font-awesome-icon></span>
          </button>
        </div>
      </nav>
      
    </div>
  </section>
</template>

<script>
import NoSleep from 'nosleep.js';
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
      },
      noSleep: new NoSleep(),
    }
  },
  computed: {
    isFullscreen: function() {
      var fullscreen = document.fullScreen || document.mozFullScreen || document.webkitIsFullScreen
      if(!fullscreen) {
        this.noSleep.disable()
      }
      return fullscreen
    }
  },
  methods: {
    refreshData: function() {
      this.$apollo.query({
        query: KEGERATOR_QUERY,
        update: data => this.kegerator = data
      })
    },
    enableFullscreen: function() {
      this.noSleep.enable()
      var element = document.documentElement
      if(element.requestFullscreen) {
        element.requestFullscreen();
      } else if(element.msRequestFullscreen) {
        element.msRequestFullscreen();
      } else if(element.webkitRequestFullscreen) {
        element.webkitRequestFullscreen();
      }
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
