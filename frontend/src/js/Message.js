import Vue from 'vue'

export default {
  data () {
    return {
      xMessage: {
        message: '',
        timeout: 2000,
        type: 'info',
        show: false
      }
    }
  },

  methods: {
    $message (msg) {
      Vue.set(this.xMessage, 'type', 'info')
      for (var k in msg) {
        Vue.set(this.xMessage, k, msg[k])
      }
      Vue.set(this.xMessage, 'show', true)
    }
  }
}
