<template>
  <div class="sidebar">
    <el-menu :default-active="onRoutes" class="el-menu-vertical-demo" background-color="#424f63" text-color="#fff" active-text-color="#ffd04b" unique-opened router>
      <template v-for="item in items">
        <template v-if="item.subs">
          <el-submenu :key="item.index" :index="item.index">
            <template slot="title">
              <icon class="menu-logo" :name="item.icon" />{{ item.title }}</template>
            <el-menu-item class="sidebar-sub-item" v-for="(subItem,i) in item.subs" :key="i" :index="subItem.index"><span class="text">{{ subItem.title }}</span>
            </el-menu-item>
          </el-submenu>
        </template>
        <template v-else>
          <el-menu-item :key="item.index" :index="item.index">
            <icon class="menu-logo" :name="item.icon" />{{ item.title }}
          </el-menu-item>
        </template>
      </template>
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          icon: 'user-plus',
          index: 'account',
          title: '账号申请'
        },
        {
          icon: 'search',
          index: 'search',
          title: '内容查询'
        },
        {
          icon: 'history',
          index: 'compare',
          title: '版权追溯'
        },
        {
          icon: 'plus',
          index: 'submit',
          title: '内容发布'
        },
        {
          icon: 'chart-area',
          index: 'monitor',
          title: '数据监控'
        }
      ]
    }
  },
  computed: {
    onRoutes() {
      return this.$route.path.replace('/', '')
    }
  }
}
</script>

<style scoped>
.sidebar {
  display: block;
  position: absolute;
  width: 181px;
  left: 0;
  top: 50px;
  bottom: 0;
  background: #545c64;
}
.sidebar > ul {
  height: 100%;
}
/* .sidebar .sidebar-sub-item {
  min-width: 180px;
  max-width: 180px;
} */
.sidebar .sidebar-sub-item .text {
  margin-left: 10px;
}
.sidebar .menu-logo {
  margin-right: 10px;
  margin-bottom: 3px;
  width: 18px;
  height: 18px;
  font-size: 16px;
}
</style>
