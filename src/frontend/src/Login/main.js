import Vue from "vue"
import Login from "./Login.vue"
import Meta from 'vue-meta';
import "uikit/dist/css/uikit.min.css"
import "uikit/dist/css/uikit-core.min.css"
import "uikit/dist/js/uikit.min.js"
import "uikit/dist/js/uikit-core.min.js"
import "uikit/dist/js/uikit-icons.min.js"

Vue.use(Meta);

Vue.config.productionTip = false
new Vue({
    el: "#app",
    render: h => h(Login)
})

