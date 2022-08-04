import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import VueGoogleMaps from '@fawmi/vue-google-maps'

axios.defaults.baseURL = 'http://localhost:8888/api/'
axios.defaults.withCredentials = true

createApp(App).use(store).use(router).use(VueGoogleMaps, {
    load: {
        key: 'AIzaSyCb7FqJb6W_n39blFQMfna9u1wUcDD-ABk',
    },
}).mount('#app')
