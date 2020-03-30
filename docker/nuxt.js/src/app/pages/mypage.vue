<template>
  <v-container :list='list'>
    <h1>こんにちは、{{this.$auth.$state.user.family_name}}さん!</h1>
    <v-btn color="normal" @click="logout">ログアウト</v-btn>
    <li v-for='l in list'>
    {{ l }}
    </li>
  </v-container>
</template>
<script>
export default {
  methods: {
    logout () {
      this.$auth.logout()
    }
  },
  async asyncData (context) {
    const response = await context.app.$axios.$get('http://localhost:3000/api')
    console.log(response)
    return { list: response }
  }
}
</script>
