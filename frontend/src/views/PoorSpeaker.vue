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
      <v-layout row wrap>
        <v-flex xs9>
          <v-text-field
            label="Please input a word"
            boxdd
            v-model="inputWord">
          </v-text-field>
        </v-flex>
        <v-flex xs3>
          <v-btn
            color="success"
            @click="addWord">
            add</v-btn>
        </v-flex>
      </v-layout>

      <div v-for="(word, idx) in wordList" v-bind:key="word.id">
        <v-layout row wrap class="row-word"
                           v-bind:class="{'row-word-dark': idx % 2 === 0}">
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
              @click="delWord(idx)">
              remove</v-btn>
          </v-flex>
        </v-layout>
      </div>
    </v-container>
  </v-app>

    <!-- <el-row> -->
    <!--   <div v-for="(word, idx) in wordList" v-bind:key="word.id"> -->
    <!--     <span>{{ word.word }}</span> -->
    <!--     <span>{{ idx }}</span> -->
    <!--     <el-button round -->
    <!--                plain -->
    <!--                type="danger" -->
    <!--                @click="delWord(idx)"> -->
    <!--       remove -->
    <!--     </el-button> -->
    <!--     <mini-audio :src="word.am.audio"/> -->
    <!--   </div> -->
      <!-- <el-table border :data="wordList"> -->
      <!--   <el-table-column prop="word" label="word"></el-table-column> -->
      <!--   <el-table-column prop="means" label="means"></el-table-column> -->

      <!--   <el-table-column label="symbols"> -->
      <!--     <template slot-scope="scope"> -->
      <!--       <div v-if="scope.row.am !== undefined" -->
      <!--            class="row-word-symbol"> -->
      <!--         <span>us: </span> -->
      <!--         <span>[{{ scope.row.am.symbol }}]</span> -->
      <!--         <mini-audio :src="scope.row.am.audio"/> -->
      <!--       </div> -->

      <!--       <div v-if="scope.row.uk !== undefined" -->
      <!--            class="row-word-symbol"> -->
      <!--         <span>uk: </span> -->
      <!--         <span>[{{ scope.row.uk.symbol }}]</span> -->
      <!--         <mini-audio :src="scope.row.uk.audio"/> -->
      <!--       </div> -->
      <!--     </template> -->
      <!--   </el-table-column> -->

      <!--   <el-table-column label="remove" -->
      <!--                    width="100px"> -->
      <!--     <template slot-scope="scope"> -->
      <!--       <el-button round -->
      <!--                  plain -->
      <!--                  type="danger" -->
      <!--                  @click="delWord(scope.$index)"> -->
      <!--         remove -->
      <!--       </el-button> -->
      <!--     </template> -->
      <!--   </el-table-column> -->
      <!-- </el-table> -->

      <!-- <div class="button-panel"> -->
      <!--   <el-button round -->
      <!--              plain -->
      <!--              type="primary" -->
      <!--              @click="generateExercise"> -->
      <!--     generate exercise -->
      <!--   </el-button> -->
      <!-- </div> -->
    <!-- </el-row> -->

    <!-- <el-row> -->
      <!-- <el-table -->
      <!--   border -->
      <!--   :data="exercise"> -->
      <!--   <el-table-column label="audio" width="100px"> -->
      <!--     <template slot-scope="scope"> -->
      <!--       <mini-audio :src="scope.row.audio"/> -->
      <!--     </template> -->
      <!--   </el-table-column> -->

      <!--   <el-table-column label="option"> -->
      <!--     <template slot-scope="scope"> -->
      <!--       <div class="answer" -->
      <!--            v-bind:class="{'answer-error': scope.row.err}"> -->
      <!--         <el-radio-group size="mini" v-model="scope.row.answer"> -->
      <!--           <el-radio v-for="word in wordList" -->
      <!--                     v-bind:key="word.id" -->
      <!--                     v-bind:label="word.id" -->
      <!--                     border> -->
      <!--             {{ word.word }} -->
      <!--           </el-radio> -->
      <!--         </el-radio-group> -->
      <!--       </div> -->
      <!--     </template> -->
      <!--   </el-table-column> -->
      <!-- </el-table> -->

      <!-- <div class="button-panel"> -->
      <!--   <el-button round -->
      <!--              plain -->
      <!--              type="primary" -->
      <!--              @click="submitExercise"> -->
      <!--     submit -->
      <!--   </el-button> -->
      <!-- </div> -->
    <!-- </el-row> -->

</template>

<script>
import FanyiBaidu from '@/js/FanyiBaidu'
import MiniAudio from '@/components/MiniAudio'
import Message from '@/js/Message'

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
      exercise: []
    }
  },

  components: {
    'mini-audio': MiniAudio
  },

  computed: {
    hasWords () {
      return this.wordList.length !== 0
    }
  },

  methods: {
    addWord () {
      if (this.inputWord === '') {
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

    generateExercise () {
      this.exercise = []

      var n = this.wordList.length * 4
      for (var i = 0; i < n; i++) {
        var j = this.getRandomInt(this.wordList.length)
        var word = this.wordList[j]

        var exer = {
          wid: word.id,
          audio: this.getWordAudio(word),
          answer: undefined,
          err: false
        }
        this.exercise.push(exer)
      }
    },

    submitExercise () {
      for (var i in this.exercise) {
        var exer = this.exercise[i]
        this.$set(exer, 'err', exer.wid !== exer.answer)
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
  margin: 5px 0;
  background: #eee;
}

.row-word-dark {
  background: #ddd;
}
.row-word .v-card {
  margin: 6px 0;
}
</style>
