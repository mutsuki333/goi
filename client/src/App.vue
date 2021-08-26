<template lang="pug">
v-app
  v-app-bar(app  color='white' dense )
    v-btn(text to='/about' color="primary") About

  v-main
    router-view

  v-footer(app padless)
    v-col.text-center(cols="12")
      p 版權所有©奇勤科技股份有限公司保留所有權利
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
};
</script>
