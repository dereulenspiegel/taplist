<template>
  <div class="box tile is-anecstor beertap" style="min-width: 250px">
    <modal-select-batch :tapId="tap.id" :tapNumber="tap.number" :active="modal.active" @cancel="modalCancel" @selected="batchSelected"/>
    <div class="tile is-vertical" v-on:click.prevent="showModal">
      <div class="tile is-parent">
        <div class="tile is-child" v-if="tap.beer">
          <h1 class="title">{{tap.beer.name}}</h1>
          <h2 class="subtitle" v-if="tap.beer.breweryName">{{tap.beer.breweryName}}</h2>
        </div>
        <div v-else class="tile is-child">
          <h1 class="title">Empty</h1>
          <h2 class="subtitle">Please refill</h2>
        </div>
      </div>

      <div class="tile is-horizontal">
        <div class="tile is-parent">
          <div class="tile is-child is-1">
            <span class="tapnumber">{{tap.number}}</span>
          </div>
          <div class="tile is-child is-4">
            <!-- Beer Image -->
            <BeerPint v-if="tap.beer" width="150px" v-bind:ebc="tap.beer.colorEbc"/>
            <BeerPint v-else width="150px" ebc="0" />
          </div>
        </div>
        
        <div class="tile is-parent is-vertical is-7" v-if="tap.beer">
          <div class="tile is-child" v-if="hasFacts">
            <span v-if="tap.beer.abv">ABV: {{tap.beer.abv}} %</span><br>
            <span v-if="tap.beer.ibu">IBU: {{tap.beer.ibu}} </span><br>
            <span v-if="tap.beer.buGuRatio">BU-GU-Ratio: {{tap.beer.buGuRatio}}</span><br>
          </div>

          <div class="tile is-child is-7" v-if="hasSensorData">
            Beer sensor data
          </div>
        </div>

        <div class="tile is-parent is-vertical" v-else>
          <div class="tile is-child">
            <!-- Insert sad smiley -->
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import BeerPint from './BeerPint.vue'
import ModalSelectBatch from './ModalSelectBatch.vue'

export default {
  name: 'BeerTap',
  components: {
    BeerPint,
    ModalSelectBatch
  },
  props: {
    tap: Object
  },
  data: function() {
    return {
      modal: {
        active: false,
      }
    }
  },
  computed: {
    hasFacts: function() {
      return this.tap.beer && (this.tap.beer.og || this.tap.beer.abv || this.tap.beer.buGuRation || this.tap.beer.ibu)
    },
    hasSensorData: function() {
      return this.tap.sensors && this.tap.sensors.length > 0
    }
  },
  methods: {
    showModal: function() {
      this.modal.active = true
    },
    modalCancel: function() {
      this.modal.active = false
    },
    batchSelected: function() {
      this.modal.active = false
      this.$emit('tapUpdated')
    }
  }
}
</script>

<style>
.beertap {
  color: black !important;
  margin: 20px;
}

.tapnumber {
  color: black;
  font-size: 6em !important;
  margin: 5px;
}
</style>
