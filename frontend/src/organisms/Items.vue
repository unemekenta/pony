<template lang="pug">
  .wrapper
    .main-allMessage-contents-userIcon
      img(v-bind:src="allMessage.CommentUserImage")
      p.main-allMessage-contents-userIcon-name {{allMessage.CommentUser}}
    p.main-allMessage-contents-txt {{allMessage.Content}}
    .main-allMessage-contents-category(v-for="catrgory in catrgoriesOfComments")
      li {{catrgory.Content}}
    p.main-allMessage-contents-time {{createDate(allMessage.CreatedAt)}}

</template>

<script>
import axios from 'axios'

export default {
  name: 'Items',
  data () {
    return {
      catrgoriesOfComments: [],
    }
  },
  props: {
    allMessage: Object,
  },
  created() {
    this.getCatrgoriesOfComments(this.allMessage.ID);
  },
  methods: {
    async getCatrgoriesOfComments (commentsID) {
      const comments_id = commentsID
      await axios.get('api/attachedCategories/' + comments_id)
      .then(response => {
        this.catrgoriesOfComments = response.data
      // catchでエラー時の挙動を定義する
      }).catch(err => {
          console.log('err:', err);
      });
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
