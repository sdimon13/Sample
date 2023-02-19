import Vue from 'vue'
import App from './App.vue'
import router from './router'
import BootstrapVue from 'bootstrap-vue'
import Axios from 'axios'
import VueTheMask from 'vue-the-mask'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import './styles/app.css'

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.prototype.$http = Axios
Vue.use(VueTheMask)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')


function show() {
  var p = document.getElementById('pwd');
  p.setAttribute('type', 'text');
}

function hide() {
  var p = document.getElementById('pwd');
  p.setAttribute('type', 'password');
}

var pwShown = 0;

document.getElementById("eye").addEventListener("click", function () {
  if (pwShown == 0) {
    pwShown = 1;
    show();
  } else {
    pwShown = 0;
    hide();
  }
}, false);
