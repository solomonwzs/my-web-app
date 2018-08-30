<template>
  <div class="mini-audio">
    <audio ref="audio">
      <source :src="src">
    </audio>
    <v-btn
      :small="small === true"
      color="info"
      @click="play">
      {{ text }}
    </v-btn>
  </div>
</template>

<script>
export default {
  props: ['src', 'text', 'small'],

  methods: {
    play () {
      this.$refs.audio.play()
    }
  },

  destroyed () {
  }
}
</script>
