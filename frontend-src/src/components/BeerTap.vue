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
          {{tap.number}}
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
    tap: Object
  },
  computed: {
    hasFacts: function() {
      return this.tap.beer.og || this.tap.beer.abv || this.tap.beer.buGuRation || this.tap.beer.ibu
    },
    hasSensorData: function() {
      return this.tap.sensors && this.tap.sensors.length > 0
    }
  }
}
</script>
