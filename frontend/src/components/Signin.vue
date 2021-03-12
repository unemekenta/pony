<template lang="pug">
  .signInPage
    h2 Sign in
    .signInPage-form
      input.signInPage-form-email(type="text" placeholder="Email" v-model="email")
      input.signInPage-form-password(type="password" placeholder="Password" v-model="password")
      button.signInPage-form-button(@click="signIn") ログイン
      p アカウントをまだお持ちでない方はこちら
      router-link(to="/signup") アカウントを作成する
</template>

<script>
// Firebase読み込み
import firebase from 'firebase/app'
// import 'firebase/firestore'

export default {
  name: 'Signin',
  data: function () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    signIn () {
      firebase.auth().signInWithEmailAndPassword(this.email, this.password).then(
        user => {
          console.log(user)
          alert('ログインしました')
          this.$router.push('/')
        },
        err => {
          alert(err.message)
        }
      )
    }
  }
}
</script>

<style scoped>
  @import "../assets/css/signIn.scss";
</style>
