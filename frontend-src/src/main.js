import Vue from 'vue'
import App from './App.vue'
import { createProvider } from './vue-apollo'
import Notify from 'vue2-notify'
import './fa.config'
require('@/app.scss')

Vue.config.productionTip = false

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

Vue.use(Notify,{
  itemClass: 'notification is-danger',
  position: 'top-right',
  enter: 'fadeIn',
  leave: 'fadeOut',
  closeButtonClass: false
})

new Vue({
  apolloProvider: createProvider({
    wsEndpoint: wsURL,
    httpEndpoint: httpURL
  }),
  render: h => h(App)
}).$mount('#app')
