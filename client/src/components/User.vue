<template>
  <div class="user">
    <h1>{{ userName + 'さんのページ' }}</h1>
    <div>
      <input type="text" name="Name">
      <button type="submit">タスク作成</button>
    </div>
    <ul>
      <li v-for="task_info in task_infos" v-bind:key="task_info.Name">
        {{ task_info.Name }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: 'User',
  data: function() {
    return {
      userName: String,
      task_infos: [],
    }
  },
  created () {
    this.userName = this.$route.params.userName;
    axios.get('https://7vr72spqg7.execute-api.ap-northeast-1.amazonaws.com/Prod/users/' + this.userName + '/tasks').then(res => {
      console.log(res);
      this.task_infos = res.data
    }).catch(err => {
      console.error(err);
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
input {
  padding: 10px;
  font-size: 19px;
  border-radius: 6px;
}
button {
  background: #0098e5;
  border: none;
  border-radius: 2px;
  color: #ffffff;
  position: relative;
  height: 36px;
  margin: 0;
  min-width: 64px;
  padding: 0 16px;
  display: inline-block;
  font-family: "Roboto", "Helvetica", "Arial", sans-serif;
  font-size: 14px;
  font-weight: 500;
  text-transform: uppercase;
  line-height: 1;
  letter-spacing: 0;
  overflow: hidden;
  will-change: box-shadow;
  transition: box-shadow 0.2s cubic-bezier(0.4, 0, 1, 1),background-color 0.2s cubic-bezier(0.4, 0, 0.2, 1),color 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  outline: none;
  text-align: center;
  line-height: 36px;
  vertical-align: middle;
}
a {
  color: #42b983;
}
</style>
