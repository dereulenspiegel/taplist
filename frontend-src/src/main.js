import Vue from 'vue'
import App from './App.vue'
import { createProvider } from './vue-apollo'
import Notify from 'vue2-notify'
import fullscreen from 'vue-fullscreen'
import VueObserveVisibility from 'vue-observe-visibility'
import {Loading} from 'buefy'
import 'buefy/dist/buefy.css'
import './fa.config'
require('@/app.scss')

Vue.config.productionTip = false

Vue.use(Loading)
Vue.use(VueObserveVisibility)
Vue.use(fullscreen)
Vue.use(Notify, {
  itemClass: 'notification',
  position: 'top-left',
  enter: 'fadeIn',
  leave: 'fadeOut',
  closeButtonClass: false
})

const types = {
  info: { itemClass: 'is-info', },
  error: { itemClass: 'is-danger' },
  warning: { itemClass: 'is-warning' },
  success: { itemClass: 'is-success', iconClass: 'fa fa-lg fa-check-circle' }
}

Vue.$notify.setTypes(types);

Vue.filter('number', function(value) {
  if (typeof value !== 'number'){
    return value
  }
  return value.toFixed(2) 
})

Vue.filter('int', function(value) {
  if (typeof value !== 'number'){
    return value
  }
  return Math.floor(value)
})

Vue.filter('precision', function(value, precision) {
  if (typeof value !== 'number'){
    return value
  }
  return value.toFixed(precision) 
})

var wsURL = ""

if(window.location.protocol === "https:") {
  wsURL = "wss://" + window.location.hostname
} else {
  wsURL = "ws://" +window.location.hostname
}

var httpURL = window.location.protocol + "//" + window.location.hostname

if(window.location.port){
  wsURL = wsURL + ":" + window.location.port
  httpURL = httpURL + ":" + window.location.port
}

wsURL = wsURL + "/query"
httpURL = httpURL + "/query"

new Vue({
  apolloProvider: createProvider({
    wsEndpoint: wsURL,
    httpEndpoint: httpURL,
    websocketsOnly: true
  }),
  render: h => h(App)
}).$mount('#app')
