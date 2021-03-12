<template lang="pug">
  .signUpPage
    h2 Sign up
    .signUpPage-form
      input.signUpPage-form-email(type="text" placeholder="Email" v-model="email")
      input.signUpPage-form-password(type="password" placeholder="Password" v-model="password")
      button.signUpPage-form-button(@click="signUp") Register
      p アカウントをすでにお持ちの方はこちら
      router-link(to="/signin") ログインする
</template>

<script>
// Firebase読み込み
import firebase from 'firebase/app'

export default {
  name: 'Signup',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signUp () {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
        .then(user => {
          alert('Create account: ', user.email)
        })
        .catch(error => {
          alert(error.message)
        })
    }
  }
}
</script>

<style scoped>
@import "../assets/css/signUp.scss";
</style>
