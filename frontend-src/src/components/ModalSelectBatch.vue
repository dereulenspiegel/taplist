<template>
  <div class="modal" v-bind:class="{'is-active': active}">
    <div class="modal-background"></div>
    <div class="modal-content is-clipped batch-select-modal">
      <h2 class="title">Select batch for Tap {{tapNumber}}</h2>
      <ul>
        <li v-for="batch in brewfatherBatches" v-bind:key="batch.id">
          <button class="button is-primary is-large is-fullwidth batch-button" v-on:click.prevent="selectBatch(batch)">
            Batch {{batch.number}}: {{batch.beer.name}}
          </button>
        </li>
        <li>
          <button class="button is-danger is-large is-fullwidth batch-button" v-on:click.prevent="setEmpty">
            Empty
          </button>
        </li>
      </ul>
    </div>
    <button class="modal-close is-large" aria-label="close" v-on:click.prevent="cancel"></button>
  </div>
</template>

<script>

import BREWFATHER_BATCHES_QUERY from '../gql/brewfatherBatches.gql'
import gql from 'graphql-tag'

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
    },
    setEmpty: function(){
      this.$apollo.mutate({
        mutation: gql`mutation($tapId: ID!, $tapData: TapData!){
          updateTap(id: $tapId, data: $tapData){ id }
        }`,
        variables: {
          tapId: this.tapId,
          tapData: {
            empty: true
          }
        }
      }).then( () => {
        this.$emit('selected')
      }).catch((error) => {
        console.log("Error {}", error)
      })
    },
    selectBatch: function(batch) {
      this.$apollo.mutate({
        mutation: gql`mutation($tapId: ID!, $brewfatherId: ID!) {
          setBrewfatherBatchOnTap(tapId: $tapId, brewfatherBatchID: $brewfatherId) {
            id
          }
        }`,
        variables:{
          tapId: this.tapId,
          brewfatherId: batch.id
        }
      }).then( () => {
        this.$emit('selected', batch);
      }).catch((error) => {
        // TODO show error
        console.log("Error {}", error)
      })
    }
  }
}
</script>

<style scoped>
.batch-select-modal {
  background-color: white;
  color: black;
  border-radius: 5px;
  padding: 20px;
}

.batch-button {
  margin: 5px;
}
</style>