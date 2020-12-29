<template>
  <div class="box beertap">
    <modal-select-batch :tapId="tap.id" :tapNumber="tap.number" :active="modal.active" @cancel="modalCancel" @selected="batchSelected"/>
    <div class="is-flex is-flex-direction-column is-flex-wrap-nowrap is-justify-content-flex-start" v-on:click.prevent="showModal">
      <div v-if="tap.beer">
        <h1 class="title">{{tap.beer.name}}</h1>
        <h2 class="subtitle" v-if="tap.beer.breweryName">{{tap.beer.breweryName}}</h2>
      </div>
      <div v-else>
        <h1 class="title">Empty</h1>
        <h2 class="subtitle">Please refill</h2>
      </div>
      
      <div class="columns is-gapless is-vcentered">
        <div class="column is-narrow">
          <span class="tapnumber">{{tap.number}}</span>
        </div>
        <div class="column is-narrow">
          <!-- Beer Image -->
          <BeerPint v-if="tap.beer" width="100px" v-bind:ebc="tap.beer.colorEbc"/>
          <BeerPint v-else width="100px" ebc="0" />
        </div>
        <div class="column is-narrow beer-data" v-if="hasFacts">
          <p class="content">
            <span v-if="tap.beer.abv">ABV: {{tap.beer.abv}} %</span><br>
            <span v-if="tap.beer.ibu">IBU: {{tap.beer.ibu}} </span><br>
            <span v-if="tap.beer.buGuRatio">BU-GU-Ratio: {{tap.beer.buGuRatio}}</span><br>
          </p>
        </div>
        <div class="column" v-else>
          <!-- Insert sad smiley -->
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

.beer-data {
  font-size: 1.0em;
}
</style>
