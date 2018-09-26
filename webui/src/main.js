// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import {
  Vuetify,
  VApp,
  VNavigationDrawer,
  VFooter,
  VList,
  VBtn,
  VIcon,
  VGrid,
  VToolbar,
  transitions,
  VImg,
  VCard,
  VExpansionPanel,
  VDivider,
  VCheckbox,
  VDatePicker,
  VPagination,
  VDataTable,
  VSelect,
  VChip,
  VTextField,
  VDialog
} from 'vuetify'
import '../node_modules/vuetify/src/stylus/app.styl'

Vue.use(Vuetify, {
  components: {
    VApp,
    VNavigationDrawer,
    VFooter,
    VList,
    VBtn,
    VIcon,
    VGrid,
    VToolbar,
    transitions,
    VImg,
    VCard,
    VExpansionPanel,
    VDivider,
    VCheckbox,
    VDatePicker,
    VPagination,
    VDataTable,
    VSelect,
    VChip,
    VTextField,
    VDialog
  },
  theme: {
    primary: '#64B5F6',
    secondary: '#000000',
    accent: '#9c27b0',
    error: '#f44336',
    warning: '#ffeb3b',
    info: '#2196f3',
    success: '#4caf50'
  }
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
