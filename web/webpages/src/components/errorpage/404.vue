<template>
  <div>
    <el-card>
      <div class="errpage-container">
        <el-row>
          <el-col :span="12">
            <h1 class="text-404" style="margin-top: 50px;">404</h1>
            <h1 class="text-jumbo" style="margin-top: 20px;">页面不存在</h1>
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
  name: 'page404',
  data() {
    return {
      errgif: errgif + '?' + +new Date()
    }
  },
  methods: {}
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
.text-404 {
  font-size: 96px;
  font-weight: 700;
  color: #8492a6;
}
.text-jumbo {
  font-size: 32px;
  font-weight: 700;
  color: #484848;
}
</style>