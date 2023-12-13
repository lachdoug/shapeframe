import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './vuetify'
import settings from './settings'

// Load settings from local Storage.
settings()

createApp(App).use(router).use(vuetify).mount('#app')
