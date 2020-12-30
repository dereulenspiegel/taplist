<template>
  <div class="modal" v-bind:class="{'is-active': active}" v-observe-visibility="visibilityHasChanged">
    <div class="modal-background"></div>
    <div class="modal-content is-clipped batch-select-modal">
      <h2 class="title">Select batch for Tap {{tapNumber}}</h2>
      <ul v-if="loadingBatches < 1">
        <li v-for="batch in brewfatherBatches" v-bind:key="batch.id">
          <button class="button is-primary is-large is-fullwidth batch-button" v-on:click.prevent="selectBatch(batch)">
            Batch {{batch.number}}: {{batch.beer.name}} ({{batch.beer.style}})
          </button>
        </li>
        <li>
          <button class="button is-danger is-large is-fullwidth batch-button" v-on:click.prevent="setEmpty">
            Empty
          </button>
        </li>
      </ul>
      <b-loading :is-full-page="false" v-model="showSpinner" :can-cancel="false"></b-loading>
    </div>
    <button class="modal-close is-large" aria-label="close" v-on:click.prevent="cancel"></button>
  </div>
</template>

<script>

import BREWFATHER_BATCHES_QUERY from '../gql/brewfatherBatches.gql'
import gql from 'graphql-tag'
import errUtils from '../errors.js'

export default {
  name: 'ModalSelectBatch',
  props: {
    tapId: String,
    tapNumber: Number,
    active: Boolean
  },
  data: function() {
    return {
      loadingBatches: 0,
      updatingTap: 0
    }
  },
  apollo: {
    brewfatherBatches: {
      query: BREWFATHER_BATCHES_QUERY,
      error: function(error) {
        this.showError(error)
      },
      loadingKey: 'loadingBatches'
    }
  },
  computed: {
    showSpinner: function() {
      return this.loadingBatches > 0 || this.updatingTap > 0
    }
  },
  methods: {
    visibilityHasChanged: function(isVisible) {
      if(isVisible) {
        this.$apollo.queries.brewfatherBatches.refetch()
      }
    },
    showError: function(error) {
      var msg = errUtils.errorMsg(error)
      this.$notify.danger(msg)
    },
    cancel: function() {
      this.$emit('cancel')
    },
    setEmpty: function(){
      this.updatingTap++
      this.$apollo.mutate({
        mutation: gql`mutation($tapId: ID!, $tapData: TapData!){
          updateTap(id: $tapId, data: $tapData){ id }
        }`,
        variables: {
          tapId: this.tapId,
          tapData: {
            empty: true
          }
        },
        loadingKey: 'updatingTap'
      }).then( () => {
        this.updatingTap--
        this.$emit('selected')
      }).catch((error) => {
        this.updatingTap--
        this.showError(error)
        console.log("Error {}", error)
      })
    },
    selectBatch: function(batch) {
      this.updatingTap++
      this.$apollo.mutate({
        mutation: gql`mutation($tapId: ID!, $brewfatherId: ID!) {
          setBrewfatherBatchOnTap(tapId: $tapId, brewfatherBatchID: $brewfatherId) {
            id
          }
        }`,
        variables:{
          tapId: this.tapId,
          brewfatherId: batch.id
        },
        loadingKey: 'updatingTap'
      }).then( () => {
        this.updatingTap--
        this.$emit('selected', batch);
      }).catch((error) => {
        this.updatingTap--
        this.showError(error)
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