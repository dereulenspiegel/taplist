<template>
  <div class="modal" v-bind:class="{'is-active': active}">
    <div class="modal-background"></div>
    <div class="modal-content is-clipped batch-select-modal">
      <h2 class="title">Select beer for Tap {{tapNumber}}</h2>
      <ul>
        <li v-for="batch in brewfatherBatches" v-bind:key="batch.id">
          {{batch.beer.name}}
        </li>
      </ul>
    </div>
    <button class="modal-close is-large" aria-label="close" v-on:click.prevent="cancel"></button>
  </div>
</template>

<script>

import BREWFATHER_BATCHES_QUERY from '../gql/brewfatherBatches.gql'

export default {
  name: 'ModalSelectBatch',
  props: {
    tapId: String,
    tapNumber: Number,
    active: Boolean
  },
  apollo: {
    brewfatherBatches: BREWFATHER_BATCHES_QUERY
  },
  methods: {
    cancel: function() {
      this.$emit('cancel')
    }
  }
}
</script>#

<style scoped>
.batch-select-modal {
  background-color: white;
  color: black;
}
</style>