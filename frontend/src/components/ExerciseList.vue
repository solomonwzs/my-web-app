<template>
  <div>
    <div v-for="(exer, idx) in exercise" v-bind:key="exer.id">
      <div class="row-word"
           v-bind:class="{'row-word-dark': idx % 2 === 0}">
        <v-layout v-bind:class="{
        'answer-right': exer.result === true,
        'answer-error': exer.result === false
        }"
                  row wrap>
          <v-flex xs2>
            <mini-audio
              :small="true"
              :src="exer.audio"
              :text="'play'"/>
          </v-flex>

          <v-flex xs10>
            <v-radio-group v-model="exer.answer" row>
              <v-radio
                v-for="word in wordList"
                v-bind:key="word.id"
                :label="word.word"
                :disabled="mode !== 'test'"
                :value="word.id">
              </v-radio>
            </v-radio-group>
          </v-flex>
        </v-layout>
      </div>
    </div>
  </div>
</template>

<script>
import MiniAudio from '@/components/MiniAudio'

export default {
  components: {
    'mini-audio': MiniAudio
  },

  props: ['exercise', 'wordList', 'mode']
}
</script>
