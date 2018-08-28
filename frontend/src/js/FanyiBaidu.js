import Axios from 'axios'

const url = process.env.VUE_APP_BACKEND + '/_backend/'

const trans_url = 'http://fanyi.baidu.com/basetrans'
const audio_url = 'http://fanyi.baidu.com/gettts'

export default {
  methods: {
    get_audio_url (text, lan) {
      return audio_url + '?spd=3&source=web&lan=' + lan + '&text=' + text
    },

    get_trans (word, from, to, onCb, errCb) {
      var headers = {
        'X-Forward': trans_url
      }
      var lowerCase = word.toLowerCase()
      var post_data = 'query=' + lowerCase + '&from='+ from + '&to='+ to

      Axios.post(url, post_data, {
        headers: headers,
        withCredentials: true
      }).then(res => {
        if (onCb !== undefined) {
          onCb(res)
        }
      }).catch(err => {
        if (errCb !== undefined) {
          errCb(err)
        }
      })
    }
  }
}
