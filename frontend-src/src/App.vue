<template>
  <section class="hero is-fullheight">
    <div class="hero-head">
      <h3 class="title">{{kegerator.name}}</h3>
    </div>
    <div class="is-flex is-flex-direction-row is-flex-wrap-wrap is-justify-content-space-around is-align-items-center is-align-content-center hero-body">
      <BeerTap v-for="tap in kegerator.taps" v-bind:key="tap.id" v-bind:tap="tap" @tapUpdated="refreshData"/>
    </div>
    <div class="hero-foot">
      <nav class="level is-mobile">
        <div class="level-left">
          <span class="level-item">GitHub: dereulenspiegel/taplist (Version {{kegerator.serverVersion}})</span>
        </div>

        <div class="level-right">
          <button v-if="!fullscreen" v-on:click.prevent="enableFullscreen" class="level-item button is-light">
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
import errUtils from './errors.js'

import KEGERATOR_QUERY from './gql/kegerator.gql'

export default {
  name: 'App',
  apollo: {
    kegerator: {
      query: KEGERATOR_QUERY,
      pollInterval: 1000,
      error: function(error) {
        this.showError(error)
      },
      update: function(data) {
        if(this.kegerator.serverVersion && this.kegerator.serverVersion != data.kegerator.serverVersion) {
          window.location.reload(true); 
        }
        return data.kegerator
      }
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
      fullscreen: false
    }
  },
  methods: {
    refreshData: function() {
      this.$apollo.queries.kegerator.refetch()
    },
    enableFullscreen: function() {
      this.$fullscreen.toggle(document.documentElement,{
        wrap: false,
        callback: this.fullscreenChange
      })
    },
    fullscreenChange: function(fullscreen) {
      if(fullscreen) {
        this.noSleep.enable()
      } else {
        this.noSleep.disable()
      }
      this.fullscreen = fullscreen
    },
    showError: function(error) {
      var msg = errUtils.errorMsg(error)
      this.$notify.danger(msg)
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
