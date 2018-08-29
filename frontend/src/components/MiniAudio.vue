<template>
  <div class="mini-audio">
    <audio ref="audio">
      <source :src="src">
    </audio>
  </div>
</template>

<script>
export default {
  props: ['src'],

  methods: {
    play () {
      this.$refs.audio.play()
    }
  },

  destroyed () {
    console.log('destroyed')
  }
}
</script>
