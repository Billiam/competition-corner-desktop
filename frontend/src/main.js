import { BrowserOpenURL } from '../wailsjs/runtime'
import {createApp} from 'vue'
import App from './App.vue'

document.body.addEventListener('contextmenu', e => e.preventDefault())
createApp(App).mount('#app')

document.body.addEventListener('click', (event) => {
  if (event.target.nodeName === 'A') {
    BrowserOpenURL(event.target.getAttribute('href'))
    event.preventDefault()
  }
})
