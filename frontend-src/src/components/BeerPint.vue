<template>
  <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" x="0px" y="0px" v-bind:width="width" viewBox="0 0 100 100" enable-background="new 0 0 100 100" xml:space="preserve">
    <path d="M50,5.229c19.883,0,24.31,3.301,25.096,4.109l-7.441,81.638C66.221,93.484,56.561,94.705,50,94.771  c-6.561-0.063-16.22-1.283-17.653-3.794L24.905,9.337C25.688,8.53,30.108,5.229,50,5.229 M50,3.229  c-19.104,0-24.822,2.955-26.53,4.716l-0.64,0.659l0.083,0.915l7.441,81.639l0.04,0.433l0.215,0.378  c2.575,4.51,17.659,4.785,19.371,4.802c1.751-0.018,16.835-0.297,19.411-4.802l0.216-0.377l0.04-0.434L77.088,9.52l0.083-0.916  l-0.642-0.66C74.817,6.183,69.095,3.229,50,3.229L50,3.229z"></path>
    <path v-if="notEmpty" class="glass-liquid" v-bind:fill="hexColor" d="M50,15.041c-4.076,0-22.011,1.767-22.011,1.767l6.249,68.614C35.869,89.225,50,89.36,50,89.36s14.129-0.136,15.759-3.938  l6.252-68.614C72.011,16.808,70.652,15.041,50,15.041z"></path>
  </svg>
</template>

<script>
import {srmToColor} from '@brewlab/calc';

export default {
  name: 'BeerPint',
  props: {
    ebc: Number,
    width: String
  },
  computed: {
    hexColor: function() {
      var srm = this.ebc / 1.97
      return srmToColor(srm)
    },
    notEmpty: function() {
      return this.ebc > 0
    }
  }
}
</script>