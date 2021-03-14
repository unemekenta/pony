<template lang="pug">
  .top
    .header
      button.header-logoutBtn(@click="signOut") ログアウト
    .user
      img.user-icon(:src="userData.ImageURL")
      .user-name {{userData.Name}} さんのページ
    .main
      .main-myMessage
        h2.main-myMessage-title マイコメント
        .main-myMessage-contents(v-for="(myMessage, key) in this.myMessages" :key="key")
          .main-myMessage-contents-top
            img.main-myMessage-contents-top-userIcon(src="../assets/logo.png")
            button.main-myMessage-contents-top-deleteBtn(@click="deleteMessage(myMessage.ID)") 削除
          p.main-myMessage-contents-txt {{myMessage.Content}}
          p.main-myMessage-contents-time {{myMessage.CreatedAt}}
    .messageForm
      form.messageForm-form(action="")
        textarea.messageForm-form-input(v-model="messageContents" placeholder="メッセージを入力")
        button.messageForm-form-btn(@click="sendMessage()") 送信
</template>

<script>
import firebase from 'firebase/app'
import axios from 'axios'

export default {
  name: 'MyPage',
  data () {
    return {
      db: null,
      myMessages: [],
      userData: '',
      messageContents: ''
    }
  },
  created () {
    this.getData()
  },
  methods: {
    signOut () {
      firebase.auth().signOut().then(() => {
        this.$router.push('/signin')
      })
    },
    async getAllUser () {
      let res = await axios.get('/api/users')
      console.log(res.data)
    },
    async getData () {
      firebase.auth().onAuthStateChanged(async user => {
        if (user) {
          var uid = user.uid
          let resUser = await axios.get('/api/users/' + uid)
          this.userData = resUser.data
          const UserId = resUser.data['ID']
          let resComments = await axios.get('/api/comments/' + UserId)
          this.myMessages = resComments.data
        } else {
          console.log('No user is signed in.')
        }
      })
    },
    sendMessage () {
      firebase.auth().onAuthStateChanged(async user => {
        if (user) {
          const params = new URLSearchParams();
          params.append('content', this.messageContents)
          params.append('user_id', this.userData.ID)
          await axios.post('/api/comments', params)
            .then(response => {
              this.messageContents = "";
              console.log(response)
            })
            .catch(error => {
              console.log(error)
            })
        } else {
          console.log('送信に失敗しました。')
        }
      })
      this.getData()
    },
    deleteMessage(myMessageID) {
      firebase.auth().onAuthStateChanged(async user => {
        if (user) {
          const params = myMessageID
          console.log(myMessageID)
          await axios.delete('/api/comments/' + params)
            .then(response => {
              this.messageContents = "";
              console.log(response)
            })
            .catch(error => {
              console.log(error)
            })
        } else {
          console.log('削除に失敗しました。')
        }
      })
      this.getData()
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/myPage.scss";
</style>
