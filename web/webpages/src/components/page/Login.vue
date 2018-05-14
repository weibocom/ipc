<template>
  <div class="login-wrap">
    <div class="ms-title">微博版权追溯平台</div>
    <div class="ms-login">
      <el-form :model="formLogin" :rules="ruleLogin" ref="formLogin" label-width="0px" class="demo-formLogin">
        <el-form-item prop="name">
          <el-input tabindex="1" v-model="formLogin.name" placeholder="用户名">
            <icon slot="prepend" class="input-prepend" name="user" />
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input tabindex="2" type="password" placeholder="密码" v-model="formLogin.password" @keyup.enter.native="handleLogin('formLogin')">
            <icon slot="prepend" class="input-prepend" name="lock" />
          </el-input>
        </el-form-item>
        <div class="login-btn">
          <el-button tabindex="3" type="primary" :loading="loading" @click="handleLogin('formLogin')">登录</el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import API from '@/api/api.js'
import { FormMixin } from 'components/mixins'
import JsEncrypt from 'jsencrypt/bin/jsencrypt'

export default {
  mixins: [FormMixin],
  data: function() {
    return {
      formLogin: {
        name: '',
        password: ''
      },
      ruleLogin: {
        name: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
      },
      loading: false,
      rsaPublicKey: ''
    }
  },
  created() {
    this.jse = new JSEncrypt()
    this.jse.setPublicKey(this.rsaPublicKey)
  },
  methods: {
    ...mapActions(['changeModalStatus', 'getProfile']),
    handleLogin(formName) {
      this.loading = true
      this.validateForm('formLogin').then(valid => {
        let formData = Object.assign({}, this.formLogin)
        if (formData.name === 'admin' && formData.password === 'admin') {
          this.loading = false
          this.changeModalStatus({ visible: false })
          this.getProfile(this.formLogin.name)
          this.$success('登录成功')
          this.$router.push({ name: 'register' })
        } else {
          this.$error('用户名或密码错误')
          this.loading = false
        }
      })
    }
  }
}
</script>

<style scoped>
.input-prepend {
  margin-top: 3px;
}
.login-wrap {
  position: relative;
  width: 100%;
  height: 100%;
}
.ms-title {
  position: absolute;
  top: 50%;
  width: 100%;
  margin-top: -230px;
  text-align: center;
  font-size: 30px;
  color: #fff;
}
.ms-login {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 300px;
  height: 160px;
  margin: -150px 0 0 -190px;
  padding: 40px;
  border-radius: 5px;
  background: #fff;
}
.login-btn {
  text-align: center;
}
.login-btn button {
  width: 100%;
  height: 36px;
}
.fa-icon {
  width: auto;
  height: 18px; /* 或任意其它字体大小相对值 */
}
</style>
