<template>
  <div class="tile is-anecstor beertap" style="min-width: 350px">
    <modal-select-batch :tapId="tap.id" :tapNumber="tap.number" :active="modal.active" @cancel="modalCancel" @selected="batchSelected"/>
    <div class="tile is-horizontal" v-on:click.prevent="showModal">
      <div class="tile is-parent">
        <div class="tile is-child">
          <!-- Beer Image -->
          <BeerPint v-if="tap.beer" width="150px" v-bind:ebc="tap.beer.colorEbc"/>
          <BeerPint v-else width="150px" ebc="0" />
        </div>
      </div>

      <div v-if="tap.beer" class="tile is-vertical is-parent">
        <div class="tile is-child">
          <h1>Tap {{tap.number}}<h1> {{tap.beer.name}}
        </div>
        <div v-if="hasFacts" class="tile is-child">
          <span v-if="tap.beer.abv">ABV: {{tap.beer.abv}} %</span>
          <span v-if="tap.beer.ibu">IBU: {{tap.beer.ibu}} </span>
          <span v-if="tap.beer.buGuRatio">BU-GU-Ratio: {{tap.beer.buGuRatio}}</span>
        </div>
        <div v-if="hasSensorData" class="tile is-child">
          Beer sensor data
        </div>
      </div>
      <div v-else class="tile is-vertical is-parent">
        <div class="tile is-child">
          <h1>Tap {{tap.number}}</h1> Nothing available
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
  color: white;
  margin: 20px;
  background-color: rgb(39, 39, 39);
  box-shadow: 5px 10px #161616;
}
</style>
