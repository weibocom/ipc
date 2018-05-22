<template>
  <div>
    <el-card class="page-card">
      <div class="errpage-container">
        <el-button @click="refresh" class="pan-back-btn">刷新</el-button>
        <el-row>
          <el-col :span="12">
            <h1 class="text-jumbo" style="margin-top: 50px;">网络异常</h1>
            <h1 class="text-jumbo" style="margin-top: 20px;">请稍后重试</h1>
          </el-col>
          <el-col :span="12">
            <img :src="errgif" width="313" height="428" alt="Girl has dropped her ice cream.">
          </el-col>
        </el-row>
      </div>
    </el-card>
  </div>
</template>

<script>
import errgif from '@/assets/errorpage/503.gif'
export default {
  name: 'page401',
  data() {
    return {
      errgif: errgif + '?' + +new Date(),
      timerid: ''
    }
  },
  beforeMount() {
    var self = this
    this.timerid = setInterval(refreshpage, 10000)
    function refreshpage() {
      self.refresh()
    }
  },
  beforeDestroy() {
    clearInterval(this.timerid)
  },
  methods: {
    refresh() {
      this.$router.go(-1)
    }
  }
}
</script>

<style scoped>
.errpage-container {
  width: 800px;
  margin: 100px auto;
}
.errpage-container .pan-back-btn {
  background: #008489;
  color: #fff;
}
.text-jumbo {
  font-size: 32px;
  font-weight: 700;
  color: #484848;
}
</style>