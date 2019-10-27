import Vue from "vue"
import Meta from 'vue-meta';
import MyVieos from "./MyVieos.vue"
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

Vue.use(Meta);
Vue.use(VueMaterial)

Vue.config.productionTip = false
new Vue({
    el: "#app",
    render: h => h(MyVieos)
})

