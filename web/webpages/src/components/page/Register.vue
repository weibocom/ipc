<template>
  <div>
    <el-card class="page-card">
      <div class="w-form">
        <el-form :model="regForm" label-width="90px" label-position="left" :rules="regFormRules" ref="regForm">
          <el-form-item label="发布平台" prop="company">
            <el-select class="form-item-content" v-model="regForm.company" placeholder="请选择发布平台">
              <el-option v-for="company in companys" :key="company" :label="company" :value="company">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="用户ID" prop="uid">
            <el-input class="form-item-content" v-model="regForm.uid" placeholder="请输入用户ID"></el-input>
          </el-form-item>
          <el-form-item size="medium">
            <el-button class="form-item-content" type="primary" @click="submitRegister" :loading="loading">联盟链账号注册</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
  </div>
</template>

<script>
import API from '@/api/api.js'
import { FormMixin } from 'components/mixins'
import storage from '@/utils/storage'
import { STORAGE_KEY } from '@/utils/constants'
import IBYTE from '@/utils/ibyte'
import validator from 'validator'

export default {
  mixins: [FormMixin],
  data() {
    var validateUID = function(rules, value, callback) {
      if (value === '') {
        callback(new Error('请输入用户ID'))
      } else if (
        !validator.isInt(value, {
          allow_leading_zeroes: true,
          min: 0,
          max: 9999999999999
        })
      ) {
        callback(new Error('请输入正确的用户ID，数字类型，长度不超过13个字符'))
      } else {
        callback()
      }
    }
    return {
      loading: false,
      companys: ['weibo', 'zhihu', 'facebook'],
      regForm: {
        uid: '',
        company: ''
      },
      regFormRules: {
        company: [{ required: true, message: '请选择发布平台', trigger: 'change' }],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }]
      }
    }
  },
  created() {
    if (this.companys.length > 0) {
      this.regForm.company = this.companys[0]
    }
  },
  computed: {},
  methods: {
    // 新增
    submitRegister() {
      this.validateForm('regForm')
        .then(valid => {
          let formData = Object.assign({}, this.regForm)
          this.loading = true
          API.register(formData).then(
            res => {
              if (res.data.msg.toUpperCase() === 'OK') {
                this.$success('注册成功')
              } else {
                this.$error(res.data.msg)
              }
              this.loading = false
            },
            _ => {
              this.loading = false
            }
          )
        })
        .catch(valid => {})
    }
  }
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
