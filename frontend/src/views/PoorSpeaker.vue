<template>
  <div>
    <el-row>
      <el-input
        placeholder="please input a word"
        v-model="inputWord"
        clearable>
      </el-input>
      <el-button
        icon="el-icon-circle-plus"
        circle
        @click="addWord"/>
    </el-row>

    <el-row>
      <el-table
        border
        :data="wordList">
        <el-table-column prop="word" label="word"></el-table-column>
        <el-table-column prop="means" label="means"></el-table-column>

        <el-table-column label="symbols">
          <template slot-scope="scope">
            <div v-if="scope.row.am !== undefined" class="row-word-symbol">
              <span>us: </span>
              <span>[{{ scope.row.am.symbol }}]</span>
              <mini-audio :src="scope.row.am.audio"/>
            </div>

            <div v-if="scope.row.uk !== undefined" class="row-word-symbol">
              <span>uk: </span>
              <span>[{{ scope.row.uk.symbol }}]</span>
              <mini-audio :src="scope.row.uk.audio"/>
            </div>
          </template>
        </el-table-column>

        <el-table-column
          label="remove"
          width="100px">
          <template slot-scope="scope">
            <el-button
              icon="el-icon-remove"
              circle
              @click.native.prevent="delWord(scope.$index)"/>
          </template>
        </el-table-column>
      </el-table>

      <div class="button-panel">
        <el-button
          @click="generateExercise"
          round>
          generate exercise
        </el-button>
      </div>
    </el-row>

  </div>
</template>

<script>
import FanyiBaidu from '@/js/FanyiBaidu'
import MiniAudio from '@/components/MiniAudio'

export default {
  mixins: [FanyiBaidu],

  data () {
    return {
      wordCurrentId: 0,
      inputWord: '',
      wordList: [],
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
      if (this.inputWord !== '') {
        this.get_trans(this.inputWord, 'en', 'zh',
          this.addWordInfo, this.getWordInfoFail)
      } else {
        this.$message({
          type: 'error',
          message: 'please input a word'
        })
      }
    },

    generateExercise () {
      console.log('hello')
    },

    wordMeans (means) {
      return means.join()
    },

    delWord (idx) {
      this.wordList.splice(idx, 1)
    },

    getWordInfoFail (err) {
      console.log(err)
      this.$message({
        type: 'error',
        message: 'get translate fail'
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
      this.inputWord = ''
      this.wordCurrentId += 1
    }
  }
}
</script>

<style>
.el-row {
  margin: 10px 0;
}

.el-row .el-input {
  width: 50%;
}

.el-row .row-word-symbol {
  width: 100%;
  float: left;
}

.el-row .row-word-symbol span {
  margin: 0 5px;
  padding: 5px 0;
  float: left;
}

.el-row .button-panel {
  margin: 5px 0;
}
</style>
