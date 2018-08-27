import Axios from 'axios'

const url = process.env.VUE_APP_BACKEND + '/_backend/'

const trans_url = 'http://fanyi.baidu.com/basetrans'
const audio_url = 'http://fanyi.baidu.com/gettts'
const user_agent = 'Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Mobile Safari/537.36'

export default {
  methods: {
    get_audio_url (text, lan) {
      return audio_url + '?spd=3&source=web&lan=' + lan + '&text=' + text
    },

    get_trans (word, from, to, onCb, errCb) {
      var headers = {
        'X-Forward': trans_url,
        'User-Agent': user_agent
      }
      var post_data = 'query=' + word + '&from='+ from + '&to='+ to

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
