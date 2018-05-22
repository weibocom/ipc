<template>
  <div class="full">
    <el-card class="full">
      <div class="full">
        <iframe :src="'http://10.13.216.127:3000/d/GxysaZVmk/ipc?panelId=3&orgId=2&tab=metrics&from=now-5m&to=now&kiosk'" style="width: 100%; height: 100%; border: none;"></iframe>
      </div>
    </el-card>
  </div>
</template>

<script>
import API from '@/api/api.js'
import { FormMixin } from 'components/mixins'

export default {
  mixins: [FormMixin],
  data() {
    return {}
  },
  created() {},
  computed: {},
  methods: {}
}
</script>

<style scoped>
</style>

<style>
.full {
  width: 100%;
  height: 100%;
}
.full .el-card__body {
  height: calc(100% - 40px);;
}
</style>
