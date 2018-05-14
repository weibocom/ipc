<template>
  <div>
    <el-tabs v-model="activeTabName" type="border-card">
      <el-tab-pane label="资源对比" name="resourceCompare">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-card class="page-card">
              <div slot="header">
                <span>源</span>
              </div>
              <div class="w-form-half">
                <el-form :model="leftCompareForm" label-width="90px" label-position="left" :rules="leftCompareFormRules" ref="leftCompareForm">
                  <el-form-item label="发布平台" prop="company">
                    <el-select class="form-item-content-half" v-model="leftCompareForm.company" placeholder="请选择发布平台">
                      <el-option v-for="company in companys" :key="company" :label="company" :value="company">
                      </el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="内容URL" prop="url">
                    <el-input class="form-item-content-half" v-model="leftCompareForm.url" placeholder="请输入URL" :disabled="leftCompareForm.company!=='weibo'" @blur="handleLeftURL(leftCompareForm.url)"></el-input>
                  </el-form-item>
                  <el-form-item label="用户ID" prop="uid">
                    <el-input class="form-item-content-half" v-model="leftCompareForm.uid" placeholder="请输入用户ID"></el-input>
                  </el-form-item>
                  <el-form-item label="消息ID" prop="mid">
                    <el-input class="form-item-content-half" v-model="leftCompareForm.mid" placeholder="请输入消息ID"></el-input>
                  </el-form-item>
                  <el-form-item label="内容" prop="compareInfo" v-show="false">
                    <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="leftCompareForm.compareInfo" placeholder="请输入对比内容"></el-input>
                  </el-form-item>
                </el-form>
              </div>
            </el-card>
          </el-col>
          <el-col :span="12">
            <el-card class="page-card">
              <div slot="header">
                <span>目标</span>
              </div>
              <div class="w-form-half">
                <el-form :model="rightCompareForm" label-width="90px" label-position="left" :rules="rightCompareFormRules" ref="rightCompareForm">
                  <el-form-item label="发布平台" prop="company">
                    <el-select class="form-item-content-half" v-model="rightCompareForm.company" placeholder="请选择发布平台">
                      <el-option v-for="company in companys" :key="company" :label="company" :value="company">
                      </el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="内容URL" prop="url">
                    <el-input class="form-item-content-half" v-model="rightCompareForm.url" placeholder="请输入URL" :disabled="rightCompareForm.company!=='weibo'" @blur="handleRightURL(rightCompareForm.url)"></el-input>
                  </el-form-item>
                  <el-form-item label="用户ID" prop="uid">
                    <el-input class="form-item-content-half" v-model="rightCompareForm.uid" placeholder="请输入用户ID"></el-input>
                  </el-form-item>
                  <el-form-item label="消息ID" prop="mid">
                    <el-input class="form-item-content-half" v-model="rightCompareForm.mid" placeholder="请输入消息ID"></el-input>
                  </el-form-item>
                  <el-form-item label="内容" prop="compareInfo" v-show="false">
                    <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="rightCompareForm.compareInfo" placeholder="请输入对比内容"></el-input>
                  </el-form-item>
                </el-form>
              </div>
            </el-card>
          </el-col>
        </el-row>
        <div style="margin-top: 20px; text-align: center;">
          <el-button class="form-item-content" type="primary" @click="submitResourceCompare" :loading="loading">对比</el-button>
        </div>
      </el-tab-pane>
      <el-tab-pane label="文本对比" name="frontend">
        <div class="page-card">
          <div class="w-form">
            <el-form :model="textCompareForm" label-width="90px" label-position="left" :rules="textCompareFormRules" ref="textCompareForm">
              <el-form-item label="发布平台" prop="company">
                <el-select class="form-item-content-half" v-model="textCompareForm.company" placeholder="请选择发布平台">
                  <el-option v-for="company in companys" :key="company" :label="company" :value="company">
                  </el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="内容URL" prop="url">
                <el-input class="form-item-content-half" v-model="textCompareForm.url" placeholder="请输入URL" @blur="handleURL(textCompareForm.url)"></el-input>
              </el-form-item>
              <el-form-item label="用户ID" prop="uid">
                <el-input class="form-item-content-half" v-model="textCompareForm.uid" placeholder="请输入用户ID"></el-input>
              </el-form-item>
              <el-form-item label="消息ID" prop="mid">
                <el-input class="form-item-content-half" v-model="textCompareForm.mid" placeholder="请输入消息ID"></el-input>
              </el-form-item>
              <el-form-item label="内容" prop="compareInfo">
                <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="textCompareForm.compareInfo" placeholder="请输入对比内容"></el-input>
              </el-form-item>
              <el-form-item size="medium">
                <el-button class="form-item-content-half" type="primary" @click="submitTextCompare" :loading="loading">对比</el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>
    <el-dialog title="对比结果" width="1080px" class="q-form-dialog" :visible.sync="twoResultDialogVisible">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-card class="page-card">
            <div slot="header">
              <span>源</span>
            </div>
            <el-form :model="leftResultForm" label-width="80px" label-position="left">
              <el-form-item label="发布平台" prop="company">
                <el-input class="form-item-content-half" v-model="leftResultForm.company" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="URL" prop="url">
                <el-input class="form-item-content-half" v-model="leftResultForm.url" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="用户ID" prop="uid">
                <el-input class="form-item-content-half" v-model="leftResultForm.uid" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="消息ID" prop="mid">
                <el-input class="form-item-content-half" v-model="leftResultForm.mid" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="消息内容" prop="content">
                <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="leftResultForm.content" :readonly="true"></el-input>
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card class="page-card">
            <div slot="header">
              <span>目标</span>
            </div>
            <el-form :model="rightResultForm" label-width="80px" label-position="left">
              <el-form-item label="发布平台" prop="company">
                <el-input class="form-item-content-half" v-model="rightResultForm.company" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="URL" prop="url">
                <el-input class="form-item-content-half" v-model="rightResultForm.url" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="用户ID" prop="uid">
                <el-input class="form-item-content-half" v-model="rightResultForm.uid" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="消息ID" prop="mid">
                <el-input class="form-item-content-half" v-model="rightResultForm.mid" :readonly="true"></el-input>
              </el-form-item>
              <el-form-item label="消息内容" prop="content">
                <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="rightResultForm.content" :readonly="true"></el-input>
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>
      </el-row>
      <div style="height: 40px; margin-top: 20px; text-align: center;">
        <span>相似度</span>
        <el-input style="margin-left: 10px; width: 120px;" v-model="towResultSimilarity" :readonly="true"></el-input>
      </div>
    </el-dialog>
    <el-dialog title="对比结果" width="680px" class="q-form-dialog" :visible.sync="resultDialogVisible">
      <el-form :model="resultForm" label-width="90px" label-position="left">
        <el-form-item label="URL" prop="url">
          <el-input class="form-item-content-half" v-model="resultForm.url" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="相似度" prop="similarity">
          <el-input class="form-item-content-half" v-model="resultForm.similarity" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="对比信息" prop="compareinfo">
          <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="resultForm.compareinfo" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="原文信息" prop="content">
          <el-input :autosize="{ minRows: 6 }" class="form-item-content-half" type="textarea" v-model="resultForm.content" :readonly="true"></el-input>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import API from '@/api/api.js'
import { FormMixin } from 'components/mixins'
import storage from '@/utils/storage'
import tool from '@/utils/function'
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
      activeTabName: 'resourceCompare',
      loading: false,
      companys: ['weibo', 'zhihu', 'facebook'],
      textCompareForm: {
        url: '',
        uid: '',
        mid: '',
        company: '',
        compareInfo: ''
      },
      textCompareFormRules: {
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ],
        compareInfo: [
          { required: true, message: '请输入对比内容', trigger: 'blur' }
        ]
      },
      leftCompareForm: {
        url: '',
        uid: '',
        mid: '',
        company: ''
      },
      leftCompareFormRules: {
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ]
      },
      rightCompareForm: {
        url: '',
        uid: '',
        mid: '',
        company: ''
      },
      rightCompareFormRules: {
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ]
      },
      twoResultDialogVisible: false,
      leftResultForm: {
        company: '',
        url: '',
        uid: '',
        mid: '',
        content: ''
      },
      rightResultForm: {
        company: '',
        url: '',
        uid: '',
        mid: '',
        content: ''
      },
      towResultSimilarity: '',
      resultDialogVisible: false,
      resultForm: {
        url: '',
        similarity: '',
        compareinfo: '',
        content: ''
      }
    }
  },
  created() {
    if (this.companys.length > 0) {
      this.textCompareForm.company = this.companys[0]
      this.leftCompareForm.company = this.companys[0]
      this.rightCompareForm.company = this.companys[0]
    }
  },
  computed: {},
  methods: {
    formatdata(data) {
      return data === undefined ? '' : data
    },
    // 比对
    submitResourceCompare() {
      this.validateForm('leftCompareForm')
        .then(valid => {
          this.validateForm('rightCompareForm')
            .then(valid => {
              let leftFormData = Object.assign({}, this.leftCompareForm)
              let rightFormData = Object.assign({}, this.rightCompareForm)
              let params = {
                srcUid: leftFormData.uid,
                srcMid: leftFormData.mid,
                srcCompany: leftFormData.company,
                dstUid: rightFormData.uid,
                dstMid: rightFormData.mid,
                dstCompany: rightFormData.company
              }

              this.loading = true
              API.compare_resource(params).then(
                res => {
                  if (
                    res.data.msg.toUpperCase() === 'OK' &&
                    res.data.data !== undefined
                  ) {
                    // similarity
                    this.towResultSimilarity = tool.checkData(
                      res.data.data.similarity
                    )
                    // left
                    this.leftResultForm.company = tool.checkData(
                      this.leftCompareForm.company
                    )
                    this.leftResultForm.url = tool.checkData(
                      this.leftCompareForm.url
                    )
                    this.leftResultForm.uid = tool.checkData(
                      this.leftCompareForm.uid
                    )
                    this.leftResultForm.mid = tool.checkData(
                      this.leftCompareForm.mid
                    )
                    this.leftResultForm.content = tool.checkData(
                      res.data.data.srcContent
                    )
                    // right
                    this.rightResultForm.company = tool.checkData(
                      this.rightCompareForm.company
                    )
                    this.rightResultForm.url = tool.checkData(
                      this.rightCompareForm.url
                    )
                    this.rightResultForm.uid = tool.checkData(
                      this.rightCompareForm.uid
                    )
                    this.rightResultForm.mid = tool.checkData(
                      this.rightCompareForm.mid
                    )
                    this.rightResultForm.content = tool.checkData(
                      res.data.data.dstContent
                    )
                    // show
                    this.twoResultDialogVisible = true
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
        })
        .catch(valid => {})
    },
    submitTextCompare() {
      this.validateForm('textCompareForm')
        .then(valid => {
          let formData = Object.assign({}, this.textCompareForm)
          let params = {
            dstUid: formData.uid,
            dstMid: formData.mid,
            dstCompany: formData.company,
            srcContent: formData.compareInfo
          }

          this.loading = true
          API.compare(params).then(
            res => {
              if (
                res.data.msg.toUpperCase() === 'OK' &&
                res.data.data !== undefined
              ) {
                this.resultForm.similarity = tool.checkData(
                  res.data.data.similarity
                )
                this.resultForm.url = tool.checkData(formData.url)
                this.resultForm.compareinfo = tool.checkData(
                  formData.compareInfo
                )
                this.resultForm.content = tool.checkData(res.data.data.dstContent)
                this.resultDialogVisible = true
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
    },
    handleURL(url) {
      if (url === '') {
        return
      }
      if (
        !validator.isURL(url, {
          protocols: ['http', 'https'],
          require_protocol: true,
          require_host: true,
          require_valid_protocol: true
        })
      ) {
        this.$warning('请输入合法的URL')
        return
      }
      // parse uid and mid from url, fill UID and MID input
      let obj = tool.uidmidFromURL(url)
      if (obj !== undefined && obj.uid !== undefined && obj.mid !== undefined) {
        this.textCompareForm.uid = obj.uid
        this.textCompareForm.mid = obj.mid
      } else {
        this.$warning('请输入完整的URL')
      }
    },
    handleLeftURL(url) {
      if (url === '') {
        return
      }
      if (
        !validator.isURL(url, {
          protocols: ['http', 'https'],
          require_protocol: true,
          require_host: true,
          require_valid_protocol: true
        })
      ) {
        this.$warning('请输入合法的URL')
        return
      }
      // parse uid and mid from url, fill UID and MID input
      let obj = tool.uidmidFromURL(url)
      if (obj !== undefined && obj.uid !== undefined && obj.mid !== undefined) {
        this.leftCompareForm.uid = obj.uid
        this.leftCompareForm.mid = obj.mid
      } else {
        this.$warning('请输入完整的URL')
      }
    },
    handleRightURL(url) {
      if (url === '') {
        return
      }
      if (
        !validator.isURL(url, {
          protocols: ['http', 'https'],
          require_protocol: true,
          require_host: true,
          require_valid_protocol: true
        })
      ) {
        this.$warning('请输入合法的URL')
        return
      }
      // parse uid and mid from url, fill UID and MID input
      let obj = tool.uidmidFromURL(url)
      if (obj !== undefined && obj.uid !== undefined && obj.mid !== undefined) {
        this.rightCompareForm.uid = obj.uid
        this.rightCompareForm.mid = obj.mid
      }
    }
  }
}
</script>

<style scoped>
.custom_label {
  width: 80px;
  height: 100%;
  background-color: rgb(238, 238, 238);
  /*margin-right: 12px;*/
  border-radius: 5px;
  padding-left: 12px;
  padding-right: 12px;
  color: rgba(0, 0, 0, 0.5);
}
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
