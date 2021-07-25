<template lang="pug">
.allWrapper
  .l-container
    .top
      .header
        p HOME
      .main
        .main-menu01
          .menu-menu01-content
            .menu-menu01-content-head
              | カテゴリ
            .menu-menu01-content-item(v-for="(allCatrgory, key) in this.allCatrgories" :key="key")
              li {{allCatrgory.Content}}

        .main-allMessage
            h2.main-allMessage-title コメント一覧
            .main-allMessage-contents(v-for="(allMessage, key) in this.allMessages" :key="key")
              items(:allMessage="allMessage")
        .main-menu02
          .menu-menu02-content
            .menu-menu02-content-head
              | メニュー
            .menu-menu02-content-items
              .menu-menu02-content-items-item
                router-link( to="/" ) HOME
              .menu-menu02-content-items-item
                router-link( to="/MyPage" ) マイページ
              .menu-menu02-content-items-item(@click="signOut") ログアウト

</template>

<script>
import axios from 'axios'
import firebase from 'firebase/app'
import Items from '/app/frontend/src/organisms/Items.vue'

export default {
  name: 'TopList',
  components: {
    Items,
  },
  data () {
    return {
      allMessages: [],
      allCatrgories: [],
    }
  },
  async created () {
    let resAllComments = await axios.get('/api/comments');
    for (const element of resAllComments.data) {
      let resCommentUser = await axios.get('/api/comment_users/' + element.UserID);
      element.CommentUser = resCommentUser.data.Name;
      element.CommentUserImage = resCommentUser.data.ImageURL;
    }
    this.allMessages = resAllComments.data;
    console.log(this.allMessages);
    this.getCatrgories();
  },
  methods: {
    signOut () {
      firebase.auth().signOut().then(() => {
        this.$router.push('/signin');
      })
    },
    async getCatrgories () {
      let resAllCatrgories = await axios.get('/api/categories');
      this.allCatrgories = resAllCatrgories.data;
    },
    createDate (createdAt) {
      let formattedDate = createdAt.split(/T|:|-/);
      formattedDate.pop()
      let dateString = formattedDate[0] + '年' + formattedDate[1]+ '月' + formattedDate[2] + '日' + formattedDate[3] + '時' + formattedDate[4] + '分';
      return dateString;
    },
  },
  computed: {
  }
}
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
  @import "../assets/css/topList.scss";
</style>
