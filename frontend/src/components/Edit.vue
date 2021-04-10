<template lang="pug">
  .editPage
    h2 登録情報編集
    form.signUpPage-form
      input.signUpPage-form-name(type="text" placeholder="Name" v-model="editedName")
      //- input(type="file" name="photo" @change="fileChange" accept="image/*")
      //- button(@click="upload") アップロード
      //- .center(v-if="image_url")
      //-   img(:src="image_url" width="80%")
      button.signUpPage-form-button(@click="edit") 更新
      //- button(@click="getUploadedImages") 画像を探す
</template>

<script>
// Firebase読み込み
import firebase from 'firebase/app'
import axios from 'axios'

export default {
  name: 'Edit',
  data () {
    return {
      name: '',
      image: '',
      imageName: 'icon01.png',
      image_url: '',
      uid : '',
      userId : '',
      editedName: '',
    }
  },
  created () {
    firebase.auth().onAuthStateChanged(async user => {
      if (user) {
        this.uid = user.uid
        this.getData()
      } else {
        console.log('No user is signed in.')
      }
    })
    this.getData()
    // this.getUploadedImages()
  },
  methods: {
    async getData () {
      let resUser = await axios.get('/api/users/' + this.uid)
      this.userData = resUser.data
      this.userId = resUser.data['ID']
    },

    // todo プロフィール画像の機能関連

    // fileChange (e) {
    //   this.image = e.target.files[0]
    //   this.imageName = e.target.files[0].name
    // },
    // upload () {
    //   // // ストレージオブジェクト作成
    //   var storageRef = firebase.storage().ref()
    //   // // ファイルのパスを設定
    //   var mountainsRef = storageRef.child(`images/${this.userId}/${this.image.name}`)
    //   // // ファイルを適用してファイルアップロード開始
    //   var uploadTask = mountainsRef.put(this.image)
    //   // // // ステータスを監視
    //   uploadTask.on(
    //     'state_changed',
    //     (snapshot) => {
    //       console.log('snapshot', snapshot)
    //     },
    //     (error) => {
    //       console.log('err', error)
    //     },
    //     () => {
    //       // Handle successful uploads on complete
    //       uploadTask.snapshot.ref.getDownloadURL().then(function(downloadURL) {
    //         console.log('File available at', downloadURL)
    //         this.image_url = downloadURL
    //       })
    //     }
    //  )
    // },
    // getUploadedImages () {
    //   console.log('取得中')
    //   var storageRef = firebase.storage().ref();
    //   var ImagesRef = storageRef.child(`images/${this.userId}/${this.imageName}`);
    //   ImagesRef.getDownloadURL().then(url => {
    //     console.log('imageURL is '+url)
    //   })
    // },
    edit () {
      firebase.auth().onAuthStateChanged( async user => {
        let resUser = await axios.get('/api/users/' + user.uid)
        this.userData = resUser.data
        this.name = resUser.data['name']
        const UserId = resUser.data['ID']
        const params = new URLSearchParams()
        params.append('name', this.editedName)
        // params.append('image_url', this.image_url)
        console.log('-------params' + params)
        axios.post('api/users/edit/'+ UserId, params)
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
