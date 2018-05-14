<template>
  <div class="wrapper">
    <v-head></v-head>
    <v-sidebar></v-sidebar>
    <div class="content">
      <transition name="move" mode="out-in">
        <router-view></router-view>
      </transition>
    </div>
  </div>
</template>

<script>
import vHead from './Header.vue'
import vSidebar from './Sidebar.vue'
export default {
  components: {
    vHead,
    vSidebar
  }
}
</script>

<style>
.el-card__header {
  background-color: rgb(238, 238, 238);
  color: rgba(0, 0, 0, 0.4);
  font-family: 'Roboto', 'sans-serif';
}
.el-table__header th {
  background-color: rgb(238, 238, 238);
  color: rgba(0, 0, 0, 0.4);
  font-family: 'Roboto', 'sans-serif';
}
.w-form-half .el-form-item__label {
  background-color: rgb(238, 238, 238);
  /*margin-right: 12px;*/
  border-radius: 5px;
  padding-left: 12px;
  padding-right: 12px;
  color: rgba(0, 0, 0, 0.5);
}
.w-form-half .el-form-item__error {
  margin-left: 10px;
}
.w-form {
  margin: 0 auto;
  width: 640px;
}
.w-form .el-form-item__label {
  background-color: rgb(238, 238, 238);
  /*margin-right: 12px;*/
  border-radius: 5px;
  padding-left: 12px;
  padding-right: 12px;
  color: rgba(0, 0, 0, 0.5);
}
.w-form .el-form-item__error {
  margin-left: 10px;
}
.q-form-dialog .el-form-item__label {
  background-color: rgb(238, 238, 238);
  /*margin-right: 12px;*/
  border-radius: 5px;
  padding-left: 12px;
  padding-right: 12px;
  color: rgba(0, 0, 0, 0.5);
}
.q-form-dialog .el-form-item__error {
  margin-left: 10px;
}
.form-item-content {
  margin-left: 10px;
  width: 540px;
}
.form-item-content-half {
  margin-left: 10px;
  width: 100%;
}
</style>
