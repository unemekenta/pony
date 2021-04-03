<template lang="pug">
  .signUpPage
    h2 Sign up
    form.signUpPage-form
      input.signUpPage-form-name(type="text" placeholder="Name" v-model="name")
      input.signUpPage-form-email(type="text" placeholder="Email" v-model="email")
      input.signUpPage-form-password(type="password" placeholder="Password" v-model="password")
      button.signUpPage-form-button(@click="signUp") 登録
      p アカウントをすでにお持ちの方はこちら
      router-link(to="/signin") ログインする
</template>

<script>
// Firebase読み込み
import firebase from 'firebase/app'
import axios from 'axios'

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
        .then(
        )
        .catch(error => {
          alert(error.message)
        })
        firebase.auth().onAuthStateChanged(user => {
          const params = new URLSearchParams()
          params.append('name', this.name)
          params.append('uid', user.uid)
          console.log(user.uid)
          axios.post('/api/users', params)
            .then(function (response) {
              console.log(response)
            })
            .catch(function (error) {
              console.log(error)
            })
          this.$router.push('/')
        })
    }
  }
}
</script>

<style scoped>
@import "../assets/css/signUp.scss";
</style>
