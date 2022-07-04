<template>
  <div>
    <button @click="getData">ボタン</button>
    <p>{{ msg }}</p>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',
  data () {
    return {
        msg: "null"
    };
  },
  methods: {
    getData() {
      this.$axios.get('http://localhost:12345/api/todo/').then(res => {
        // リクエスト成功時
        this.msg = res.data;
      }).catch()
    },
  },
})
</script>
