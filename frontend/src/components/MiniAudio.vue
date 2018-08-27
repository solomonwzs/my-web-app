<template>
  <div class="mini-audio">
    <audio ref="audio">
      <source :src="src">
    </audio>
    <el-button
      icon="el-icon-caret-right"
      @click="play"
      circle/>
  </div>
</template>

<script>
export default {
  props: ['src'],

  methods: {
    play () {
      this.$refs.audio.play()
    }
  }
}
</script>

<style>
.mini-audio {
  float: left;
}
</style>
