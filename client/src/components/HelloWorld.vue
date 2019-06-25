<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: 'HelloWorld',
  data: function() {
    return {
      msg: String
    }
  },
  methods: {
    clear () {
      this.msg = ''
    }
  },
  created () {
    axios.get('https://7vr72spqg7.execute-api.ap-northeast-1.amazonaws.com/Prod/hello/yujiariyasu?age=3').then(res => {
      console.log(res)
      this.msg = res.data.personName
    }).catch(err => {
      console.error(err)
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
a {
  color: #42b983;
}
</style>
