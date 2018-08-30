<template>
  <v-app>
    <v-snackbar
      v-model="xMessage.show"
      :timeout="xMessage.timeout"
      :color="xMessage.type"
      :top="true">
      {{ xMessage.message }}
    </v-snackbar>

    <v-container grid-list-md text-xs-center>
      <div>
        <v-layout row wrap>
          <v-flex xs9>
            <v-text-field
              label="Please input a word"
              :disabled='!inputEnable'
              clearable
              v-model="inputWord">
            </v-text-field>
          </v-flex>
          <v-flex xs3>
            <v-btn
              :disabled='!inputEnable'
              color="success"
              @click="addWord">
              add</v-btn>
          </v-flex>
        </v-layout>
      </div>

      <div v-if="showWordList">
        <div v-for="(word, idx) in wordList" v-bind:key="word.id">
          <v-layout class="row-word"
                    v-bind:class="{'row-word-dark': idx % 2 === 0}"
                    row wrap>
            <v-flex xs2>
              <v-card>
                {{ word.word }}
              </v-card>
            </v-flex>
            <v-flex xs6>
              <v-card>
                {{ word.means }}
              </v-card>
            </v-flex>
            <v-flex xs2>
              <mini-audio
                :small="true"
                :src="word.am.audio"
                :text="'US: [' + word.am.symbol + ']'"/>
              <mini-audio
                :small="true"
                :src="word.uk.audio"
                :text="'UK: [' + word.uk.symbol + ']'"/>
            </v-flex>
            <v-flex xs2>
              <v-btn
                small
                color="error"
                :disabled='!inputEnable'
                @click="delWord(idx)">
                remove</v-btn>
            </v-flex>
          </v-layout>
        </div>
        <v-layout row wrap>
          <v-flex xs12>
            <v-btn
              color="success"
              :disabled='!inputEnable'
              @click="generateExercise">
              generate exercise
            </v-btn>
            <v-btn
              color="error"
              :disabled='!inputEnable'
              @click="clearWordList">
              clear
            </v-btn>
          </v-flex>
        </v-layout>
      </div>

      <div v-if="showExercise">
        <exer-list
          :exercise="exercise"
          :wordList="wordList"
          :mode="mode">
        </exer-list>

        <v-layout row wrap v-if="mode === 'test-end'">
          <v-flex xs12>
            <v-card>{{ testResult }}</v-card>
          </v-flex>
        </v-layout>

        <v-layout row wrap>
          <v-flex xs12>
            <v-btn
              color="success"
              v-if="mode === 'test'"
              @click="submitExercise">
              submit
            </v-btn>
            <v-btn
              v-if="mode === 'test-end'"
              color="success"
              @click="redoExercise">
              redo
            </v-btn>
            <v-btn
              color="error"
              @click="closeExercise">
              close exercise
            </v-btn>
          </v-flex>
        </v-layout>
      </div>
    </v-container>
  </v-app>
</template>

<script>
import FanyiBaidu from '@/js/FanyiBaidu'
import Message from '@/js/Message'

import MiniAudio from '@/components/MiniAudio'
import ExerciseList from '@/components/ExerciseList'

export default {
  mixins: [FanyiBaidu, Message],

  data () {
    return {
      message: {
        text: '',
        timeout: 1000,
        show: false
      },
      wordCurrentId: 0,
      inputWord: '',
      wordList: [],
      wordSet: new Set(),
      exercise: [],
      mode: 'input'
    }
  },

  components: {
    'mini-audio': MiniAudio,
    'exer-list': ExerciseList
  },

  computed: {
    inputEnable () {
      return this.mode === 'input'
    },

    showWordList () {
      return this.wordList.length > 0
    },

    showExercise () {
      return this.mode.substr(0, 4) === 'test'
    },

    testResult () {
      var n = this.exercise.length
      var r = 0
      var e = 0
      for (var i in this.exercise) {
        var exer = this.exercise[i]
        if (exer.wid === exer.answer) {
          r += 1
        } else {
          e += 1
        }
      }
      return 'all: ' + n + ', right: ' + r + ', error: ' + e
    }
  },

  methods: {
    addWord () {
      if (this.inputWord === '' || this.inputWord === null) {
        this.$message({
          type: 'error',
          message: 'Please input a word'
        })
      } else if (this.wordSet.has(this.inputWord)){
        this.$message({
          type: 'error',
          message: 'Duplicate word'
        })
      } else {
        this.get_trans(this.inputWord, 'en', 'zh',
          this.addWordInfo, this.getWordInfoFail)
      }
    },

    getRandomInt(max) {
      return Math.floor(Math.random() * Math.floor(max));
    },

    getWordAudio(word) {
      if (word.am !== undefined && word.uk !== undefined) {
        var r = this.getRandomInt(2)
        return r === 0 ? word.am.audio : word.uk.audio
      } else if (word.am !== undefined) {
        return word.am.audio
      } else {
        return word.uk.audio
      }
    },

    clearWordList () {
      this.wordSet = new Set()
      this.wordList = []
    },

    closeExercise () {
      this.mode = 'input'
    },

    generateExercise () {
      this.exercise = []
      this.mode = 'test'

      var n = this.wordList.length * 4
      for (var i = 0; i < n; i++) {
        var j = this.getRandomInt(this.wordList.length)
        var word = this.wordList[j]

        var exer = {
          id: i,
          wid: word.id,
          word: word.word,
          audio: this.getWordAudio(word),
          answer: undefined,
          result: undefined
        }
        this.exercise.push(exer)
      }
    },

    submitExercise () {
      this.mode = 'test-end'
      for (var i in this.exercise) {
        var exer = this.exercise[i]
        this.$set(exer, 'result', exer.wid === exer.answer)
      }
    },

    redoExercise () {
      this.mode = 'test'
      for (var i in this.exercise) {
        var exer = this.exercise[i]
        this.$set(exer, 'answer', undefined)
        this.$set(exer, 'result', undefined)
      }
    },

    delWord (idx) {
      this.wordSet.delete(this.wordList[idx].word)
      this.$delete(this.wordList, idx)
    },

    getWordInfoFail (err) {
      console.log(err)
      this.$message({
        type: 'error',
        message: 'Get translate fail'
      })
    },

    addWordInfo (res) {
      var dict = res.data.dict
      var symbols = dict.symbols[0]

      var obj = {
        id: this.wordCurrentId,
        word: dict.word_name,
        means: dict.word_means.join('; '),
        symbols: dict.symbols[0]
      }

      if (symbols.ph_am !== undefined) {
        obj['am'] = {
          'symbol': symbols.ph_am,
          'audio': this.get_audio_url(dict.word_name, 'en')
        }
      }

      if (symbols.ph_am !== undefined) {
        obj['uk'] = {
          'symbol': symbols.ph_en,
          'audio': this.get_audio_url(dict.word_name, 'uk')
        }
      }

      this.wordList.push(obj)
      this.wordSet.add(obj.word)
      this.inputWord = ''
      this.wordCurrentId += 1
    }
  }
}
</script>

<style>
* {
  text-transform: none !important;
}

.row-word {
  margin: 10px 0;
  background: #eee;
}

.row-word-dark {
  background: #ddd;
}
.row-word .v-card {
  margin: 6px 0;
}

.row-word .answer-error {
  background: #fcc;
}

.row-word-dark .answer-error {
  background: #faa;
}

.row-word .answer-right {
  background: #cfc;
}

.row-word-dark .answer-right {
  background: #afa;
}
</style>
