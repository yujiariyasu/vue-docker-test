<template>
  <div class="user">
    <img alt="Vue logo" src="../assets/logo.png">
    <User/>
  </div>
</template>

<script>
// @ is an alias to /src
import User from '@/components/User.vue'

export default {
  name: 'user',
  components: {
    User
  }
}
</script>
