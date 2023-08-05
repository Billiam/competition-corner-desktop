import { BrowserOpenURL, Environment } from '../wailsjs/runtime'
import {createApp} from 'vue'
import App from './App.vue'

createApp(App).mount('#app')

const preventDefault = e => e.preventDefault()
document.body.addEventListener('contextmenu', preventDefault)
document.body.addEventListener('drop', e => {
  e.preventDefault()
})
document.body.addEventListener('dragover', e => {
  e.preventDefault()
  e.stopPropagation()
})
Environment().then(({ buildType }) => {
  if (buildType === 'dev') {
    document.body.removeEventListener('contextmenu', preventDefault)
  }
})
document.body.addEventListener('click', (event) => {
  if (event.target.nodeName === 'A') {
    BrowserOpenURL(event.target.getAttribute('href'))
    event.preventDefault()
  }
})
