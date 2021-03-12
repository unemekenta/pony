<template lang="pug">
.top
  .header
    router-link(to="/MyPage")
      button.header-myPageBtn マイページ
      button.header-logoutBtn(@click="signOut") ログアウト
  .main
    .main-allMessage
        h2.main-allMessage-title コメント一覧
        .main-allMessage-contents(v-for="(allMessage, key) in this.allMessages" :key="key")
          .main-allMessage-contents-userIcon
            img(v-bind:src="allMessage.CommentUserImage")
            p.main-allMessage-contents-userIcon-name {{allMessage.CommentUser}}
          p.main-allMessage-contents-time {{allMessage.CreatedAt}}
          p.main-allMessage-contents-txt {{allMessage.Content}}
</template>

<script>
import axios from 'axios'
import firebase from 'firebase/app'

export default {
  name: 'HelloWorld',
  data () {
    return {
      title: 'アクティビティ',
      allMessages: []
    }
  },
  async created () {
    let resAllComments = await axios.get('http://localhost:8000/api/comments')
    // console.log(resAllComments.data)
    for (const element of resAllComments.data) {
      // console.log(element.UserID)
      // console.log(element)
      let resCommentUser = await axios.get('http://localhost:8000/api/comment_users/' + element.UserID)
      // console.log(resCommentUser.data)
      element.CommentUser = resCommentUser.data.Name
      element.CommentUserImage = resCommentUser.data.ImageURL
    }
    this.allMessages = resAllComments.data
    console.log(this.allMessages)
  },
  methods: {
    signOut () {
      firebase.auth().signOut().then(() => {
        this.$router.push('/signin')
      })
    }
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/topList.scss";
</style>
