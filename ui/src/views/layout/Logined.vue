<template>
  <el-container>
    <main-aside></main-aside>
    <el-container>
      <main-header></main-header>
      <main-container></main-container>
    </el-container>
  </el-container>
</template>

<script>
import MainAside from '@/components/home/Aside.vue'
import MainHeader from '@/components/home/Header.vue'
import MainContainer from '@/components/home/Main.vue'
export default {
  components: {
    MainAside,
    MainHeader,
    MainContainer,
  },
  data() {
    return {
      key: 'value'
    }
  }
}
</script>

<style lang="scss">
.el-header,
.el-footer {
  background-color: #b3c0d1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #d3dce6;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: #e9eef3;
  color: #333;
  text-align: center;
  line-height: 160px;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>
