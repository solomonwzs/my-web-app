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
        @click="addWord"></el-button>
    </el-row>

    <el-row>
      <el-table
        border
        :data="wordList">
        <el-table-column prop="word" label="word"></el-table-column>
        <el-table-column prop="means" label="means"></el-table-column>
        <el-table-column
          label="remove"
          width="100px">
          <template slot-scope="scope">
            <el-button
              icon="el-icon-remove"
              circle
              @click.native.prevent="delWord(scope.$index)">
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
  </div>
</template>

<script>
import FanyiBaidu from '@/js/FanyiBaidu.js'
console.log(process.env)

export default {
  mixins: [FanyiBaidu],

  data () {
    return {
      wordCurrentId: 0,
      inputWord: '',
      wordList: []
    }
  },

  methods: {
    addWord () {
      if (this.inputWord !== '') {
        this.get_trans(this.inputWord, 'en', 'zh', this.addWordInfo)
        // this.wordList.push({
        //   id: this.wordCurrentId,
        //   word: this.inputWord
        // })
        // this.inputWord = ''
        // this.wordCurrentId += 1
      } else {
        this.$message({
          type: 'error',
          message: 'please input a word'
        })
      }
    },

    delWord (idx) {
      console.log(idx)
      this.wordList.splice(idx, 1)
    },

    addWordInfo (res) {
      this.wordList.push({
        id: this.wordCurrentId,
        word: res.data.dict.word_name,
        means: res.data.dict.word_means
      })
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
</style>