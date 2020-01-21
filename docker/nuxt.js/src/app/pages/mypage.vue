<template>
  <v-container>
    <h1>こんにちは、{{this.$auth.$state.user.family_name}}さん!</h1>
    <v-btn color="normal" @click="logout">ログアウト</v-btn>
  </v-container>
</template>
<script>
export default {
  methods: {
    logout () {
      this.$auth.logout()
    }
  }
}
</script>
