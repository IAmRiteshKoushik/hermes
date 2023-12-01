import './assets/main.css'

import { createApp } from 'vue'
// The createApp API allows multiple Vue applications to co-exist on the 
// same page, each with its own scope for configuration and global assets
import App from './App.vue'

createApp(App).mount('#app')
