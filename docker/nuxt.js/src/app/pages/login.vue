<template>
  <v-container>
    <v-form>
      <v-btn color="primary" @click="loginWithAuthZero">Auth0でログイン
      </v-btn>
    </v-form>
  </v-container>
</template>

<script>
export default {
  methods: {
    loginWithAuthZero () {
      this.$auth.loginWith('auth0')
    }
  }
}
</script>
