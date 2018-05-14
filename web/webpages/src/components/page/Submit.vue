<template>
  <div>
    <el-card class="page-card">
      <div class="w-form">
        <el-form :model="publishForm" label-width="90px" label-position="left" :rules="publishFormRules" ref="publishForm">
          <el-form-item label="发布平台" prop="company">
            <el-select class="form-item-content" v-model="publishForm.company" placeholder="请选择发布平台">
              <el-option v-for="company in companys" :key="company" :label="company" :value="company">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="用户ID" prop="uid">
            <el-input class="form-item-content" v-model="publishForm.uid" placeholder="请输入用户ID"></el-input>
          </el-form-item>
          <el-form-item label="消息ID" prop="mid">
            <el-input class="form-item-content" v-model="publishForm.mid" placeholder="请输入消息ID"></el-input>
          </el-form-item>
          <el-form-item label="标题" prop="title">
            <el-input class="form-item-content" v-model="publishForm.title" placeholder="请输入标题"></el-input>
          </el-form-item>
          <el-form-item label="内容" prop="content">
            <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="publishForm.content" placeholder="请输入内容"></el-input>
          </el-form-item>
          <el-form-item size="medium">
            <el-button class="form-item-content" type="primary" @click="submitPublish" :loading="loading">提交</el-button>
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
    var validateMID = function(rules, value, callback) {
      if (value === '') {
        callback(new Error('请输入消息ID'))
      } else if (
        !validator.isInt(value, {
          allow_leading_zeroes: true,
          min: 0,
          max: 9999999999999999
        })
      ) {
        callback(new Error('请输入正确的消息ID，数字类型，长度不超过16个字符'))
      } else {
        callback()
      }
    }
    return {
      loading: false,
      companys: ['weibo', 'zhihu', 'facebook'],
      publishForm: {
        uid: '',
        mid: '',
        company: '',
        title: '',
        content: ''
      },
      publishFormRules: {
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [{ required: true, message: '请选择发布平台', trigger: 'change' }],
        title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
        content: [{ required: true, message: '请输入内容', trigger: 'blur' }]
      }
    }
  },
  created() {
    if (this.companys.length > 0) {
      this.publishForm.company = this.companys[0]
    }
  },
  computed: {},
  methods: {
    // 发布
    submitPublish() {
      this.validateForm('publishForm')
        .then(valid => {
          let formData = Object.assign({}, this.publishForm)
          formData.action = 'write'
          this.loading = true
          API.submit(formData).then(
            res => {
              if (res.data.msg.toUpperCase() === 'OK') {
                this.$success('发布成功')
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
