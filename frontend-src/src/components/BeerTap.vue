<template>
  <div class="tile is-anecstor" style="min-width: 350px">
    <div class="tile is-horizontal">
      <div class="tile is-parent">
        <div class="tile is-child">
          <!-- Beer Image -->
          <p>Here comes a pint SVG</p>
        </div>
      </div>

      <div class="tile is-vertical is-parent">
        <div class="tile is-child">
          {{beer.name}}
        </div>
        <div v-if="hasFacts" class="tile is-child">
          Beer Facts 1
        </div>
        <div v-if="hasSensorData" class="tile is-child">
          Beer sensor data
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BeerTap',
  props: {
    beer: Object
  },
  computed: {
    hasFacts: function() {
      return this.beer.og || this.beer.abv || this.beer.buGuRation || this.beer.ibu
    },
    hasSensorData: function() {
      return this.beer.sensors && this.beer.sensors.length > 0
    }
  }
}
</script>
