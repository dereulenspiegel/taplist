import Vue from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

Vue.component('font-awesome-icon', FontAwesomeIcon);

import { faUser as fasUser, faBuilding, faExpandArrowsAlt } from '@fortawesome/free-solid-svg-icons'
import { faUser as farUser, faCheckCircle } from '@fortawesome/free-regular-svg-icons'

library.add(
  fasUser,
  farUser,
  faBuilding,
  faCheckCircle,
  faExpandArrowsAlt
);