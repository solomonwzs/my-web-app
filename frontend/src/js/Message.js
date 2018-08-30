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
      this.$set(this.xMessage, 'type', 'info')
      for (var k in msg) {
        this.$set(this.xMessage, k, msg[k])
      }
      this.$set(this.xMessage, 'show', true)
    }
  }
}
