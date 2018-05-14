<template>
  <div>
    <el-card class="page-card">
      <div class="w-form" style="width: 100%; height: 560px;">
        <iframe :src="'http://10.13.216.127:3000/dashboard-solo/db/test?orgId=2&panelId=1'" style="width: 100%; height: 100%; border: none;"></iframe>
      </div>
      <div class="w-form" style="width: 100%; height: 560px; margin-top: 20px;">
        <iframe :src="'http://10.13.216.127:3000/dashboard-solo/db/test?orgId=2&panelId=2'" style="width: 100%; height: 100%; border: none;"></iframe>
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
.page-card .el-card__header {
  background-color: rgb(238, 238, 238);
  color: rgba(0, 0, 0, 0.4);
  font-family: 'Roboto', 'sans-serif';
}
.page-card .el-table__header th {
  background-color: #fff;
}
</style>
