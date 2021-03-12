<template lang="pug">
  .top
    .header
      router-link.header-allCommentsBtn(to="/")
        button 投稿一覧
      button.header-logoutBtn(@click="signOut") ログアウト
    .user
      img.user-icon(:src="userData.ImageURL")
      .user-name {{userData.Name}} さんのページ
    .main
      .main-myMessage
        h2.main-myMessage-title マイコメント
        .main-myMessage-contents(v-for="(myMessage, key) in this.myMessages" :key="key")
          img.main-myMessage-contents-userIcon(src="../assets/logo.png")
          p.main-myMessage-contents-txt {{myMessage.Content}}
          p.main-myMessage-contents-time {{myMessage.CreatedAt}}
    .messageForm
      form.messageForm-form(action="" @submit="sendMessage()")
        textarea.messageForm-form-input(v-model="messageContents" placeholder="メッセージを入力")
        button.messageForm-form-btn(type="submit") 送信
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
    firebase.auth().onAuthStateChanged(async user => {
      if (user) {
        var uid = user.uid
        let resUser = await axios.get('http://localhost:8000/api/users/' + uid)
        console.log('-------created', resUser.data)
        this.userData = resUser.data
        const UserId = resUser.data['ID']
        let resComments = await axios.get('http://localhost:8000/api/comments/' + UserId)
        console.log('-------comments', resComments.data)
        this.myMessages = resComments.data
        // const db = firebase.firestore()
        // db.collection('comments').get().then(snap => {
        //   const array = []
        //   snap.forEach(doc => {
        //     array.push(doc.data())
        //   })
        //   this.allMessages = array
        // })
        // const query = db.collection('comments').where('uid', '==', uid)
        // query.get().then(snap => {
        //   const array = []
        //   snap.forEach(doc => {
        //     // console.log(doc.data())
        //     array.push(doc.data())
        //   })
        //   this.myMessages = array
        // })
        // const fechUser = db.collection('users').where('uid', '==', uid)
        // fechUser.get().then(snap => {
        //   const userArray = []
        //   snap.forEach(doc => {
        //     userArray.push(doc.data())
        //   })
        //   this.userData = userArray
        //   this.iconUrl = this.userData[0].imageid
        // })
      } else {
        console.log('No user is signed in.')
      }
    })
  },
  methods: {
    signOut () {
      firebase.auth().signOut().then(() => {
        this.$router.push('/signin')
      })
    },
    getAllUser: async function () {
      let res = await axios.get('http://localhost:8000/api/users')
      console.log(res.data)
    },
    // getUserName (uid) {
    //   let userIconUrls = []
    //   const db = firebase.firestore()
    //   const comentUser = db.collection('users').where('uid', '==', uid)
    //   comentUser.get().then(snap => {
    //     snap.forEach(doc => {
    //       const userData = doc.data()
    //       const userIconUrl = userData.imageid
    //       console.log(typeof (userIconUrl))
    //       userIconUrls.push(userIconUrl)
    //     })
    //   })
    //   console.log(typeof (userIconUrls))
    //   return userIconUrls
    // },
    sendMessage () {
      // var date = new Date()
      // const db = firebase.firestore()
      firebase.auth().onAuthStateChanged(user => {
        if (user) {
        //   axios.post('http://localhost:8000/api/comments', {
        //     content: this.messageContents,
        //     user_id: this.userData.userID
        //   })
        //     .then(function (response) {
        //       console.log(response)
        //     })
        //     .catch(function (error) {
        //       console.log(error)
        //     })
        // } else {
        //   console.log('保存に失敗しました。')
        }
      })
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/myPage.scss";
</style>
