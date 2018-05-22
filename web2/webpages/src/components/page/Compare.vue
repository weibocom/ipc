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
                    <el-select class="form-item-content" v-model="leftCompareForm.company" placeholder="请选择发布平台">
                      <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
                      </el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="查询类型" prop="type">
                    <el-select class="form-item-content" v-model="resourcecomparetype" placeholder="请选择查询类型">
                      <el-option v-for="restype in comparetypes" :key="restype.value" :label="restype.label" :value="restype.value">
                      </el-option>
                    </el-select>
                  </el-form-item>
                  <div v-if="resourcecomparetype === 'user'">
                    <el-form-item label="内容URL" prop="url">
                      <el-input class="form-item-content" v-model="leftCompareForm.url" placeholder="请输入URL" @blur="handleLeftURL(leftCompareForm.url)"></el-input>
                    </el-form-item>
                    <el-form-item label="用户ID" prop="uid">
                      <el-input class="form-item-content" v-model="leftCompareForm.uid" placeholder="请输入用户ID"></el-input>
                    </el-form-item>
                    <el-form-item label="消息ID" prop="mid">
                      <el-input class="form-item-content" v-model="leftCompareForm.mid" placeholder="请输入消息ID"></el-input>
                    </el-form-item>
                    <el-form-item label="内容" prop="compareInfo" v-show="false">
                      <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="leftCompareForm.compareInfo" placeholder="请输入对比内容"></el-input>
                    </el-form-item>
                  </div>
                  <div v-if="resourcecomparetype === 'dna'">
                    <el-form-item label="哈希值" prop="dna">
                      <el-input class="form-item-content" v-model="leftCompareForm.dna" placeholder="请输入哈希值"></el-input>
                    </el-form-item>
                  </div>
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
                    <el-select class="form-item-content" v-model="rightCompareForm.company" placeholder="请选择发布平台">
                      <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
                      </el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="查询类型" prop="type">
                    <el-select class="form-item-content" v-model="resourcecomparetype" placeholder="请选择查询类型">
                      <el-option v-for="restype in comparetypes" :key="restype.value" :label="restype.label" :value="restype.value">
                      </el-option>
                    </el-select>
                  </el-form-item>
                  <div v-if="resourcecomparetype === 'user'">
                    <el-form-item label="内容URL" prop="url">
                      <el-input class="form-item-content" v-model="rightCompareForm.url" placeholder="请输入URL" @blur="handleRightURL(rightCompareForm.url)"></el-input>
                    </el-form-item>
                    <el-form-item label="用户ID" prop="uid">
                      <el-input class="form-item-content" v-model="rightCompareForm.uid" placeholder="请输入用户ID"></el-input>
                    </el-form-item>
                    <el-form-item label="消息ID" prop="mid">
                      <el-input class="form-item-content" v-model="rightCompareForm.mid" placeholder="请输入消息ID"></el-input>
                    </el-form-item>
                    <el-form-item label="内容" prop="compareInfo" v-show="false">
                      <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="rightCompareForm.compareInfo" placeholder="请输入对比内容"></el-input>
                    </el-form-item>
                  </div>
                  <div v-if="resourcecomparetype === 'dna'">
                    <el-form-item label="哈希值" prop="dna">
                      <el-input class="form-item-content" v-model="rightCompareForm.dna" placeholder="请输入哈希值"></el-input>
                    </el-form-item>
                  </div>
                </el-form>
              </div>
            </el-card>
          </el-col>
        </el-row>
        <div style="margin-top: 20px; text-align: center;">
          <el-button class="form-item-content" style="width: 600px;" type="primary" @click="submitResourceCompare" :loading="resloading">对比</el-button>
        </div>
      </el-tab-pane>
      <el-tab-pane label="文本对比" name="textCompare">
        <div class="page-card">
          <div class="w-form">
            <el-form :model="textCompareForm" label-width="90px" label-position="left" :rules="textCompareFormRules" ref="textCompareForm">
              <el-form-item label="发布平台" prop="company">
                <el-select class="form-item-content" v-model="textCompareForm.company" placeholder="请选择发布平台">
                  <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
                  </el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="查询类型" prop="type">
                <el-select class="form-item-content" v-model="textcomparetype" placeholder="请选择查询类型">
                  <el-option v-for="restype in comparetypes" :key="restype.value" :label="restype.label" :value="restype.value">
                  </el-option>
                </el-select>
              </el-form-item>
              <div v-if="textcomparetype === 'user'">
                <el-form-item label="内容URL" prop="url">
                  <el-input class="form-item-content" v-model="textCompareForm.url" placeholder="请输入URL" @blur="handleURL(textCompareForm.url)"></el-input>
                </el-form-item>
                <el-form-item label="用户ID" prop="uid">
                  <el-input class="form-item-content" v-model="textCompareForm.uid" placeholder="请输入用户ID"></el-input>
                </el-form-item>
                <el-form-item label="消息ID" prop="mid">
                  <el-input class="form-item-content" v-model="textCompareForm.mid" placeholder="请输入消息ID"></el-input>
                </el-form-item>
              </div>
              <div v-if="textcomparetype === 'dna'">
                <el-form-item label="哈希值" prop="dna">
                  <el-input class="form-item-content" v-model="textCompareForm.dna" placeholder="请输入哈希值"></el-input>
                </el-form-item>
              </div>
              <el-form-item label="内容" prop="compareInfo">
                <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="textCompareForm.compareInfo" placeholder="请输入对比内容"></el-input>
              </el-form-item>
              <el-form-item size="medium">
                <el-button class="form-item-content" type="primary" @click="submitTextCompare" :loading="textloading">对比</el-button>
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
                <el-input class="form-item-content" v-model="leftResultForm.company" :readonly="true"></el-input>
              </el-form-item>
              <div v-if="resourcecomparetype === 'user'">
                <el-form-item label="URL" prop="url">
                  <el-input class="form-item-content" v-model="leftResultForm.url" :readonly="true"></el-input>
                </el-form-item>
                <el-form-item label="用户ID" prop="uid">
                  <el-input class="form-item-content" v-model="leftResultForm.uid" :readonly="true"></el-input>
                </el-form-item>
                <el-form-item label="消息ID" prop="mid">
                  <el-input class="form-item-content" v-model="leftResultForm.mid" :readonly="true"></el-input>
                </el-form-item>
              </div>
              <div v-if="resourcecomparetype === 'dna'">
                <el-form-item label="哈希值" prop="dna">
                  <el-input class="form-item-content" v-model="leftResultForm.dna" :readonly="true"></el-input>
                </el-form-item>
              </div>
              <el-form-item label="消息内容" prop="content">
                <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="leftResultForm.content" :readonly="true"></el-input>
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
                <el-input class="form-item-content" v-model="rightResultForm.company" :readonly="true"></el-input>
              </el-form-item>
              <div v-if="resourcecomparetype === 'user'">
                <el-form-item label="URL" prop="url">
                  <el-input class="form-item-content" v-model="rightResultForm.url" :readonly="true"></el-input>
                </el-form-item>
                <el-form-item label="用户ID" prop="uid">
                  <el-input class="form-item-content" v-model="rightResultForm.uid" :readonly="true"></el-input>
                </el-form-item>
                <el-form-item label="消息ID" prop="mid">
                  <el-input class="form-item-content" v-model="rightResultForm.mid" :readonly="true"></el-input>
                </el-form-item>
              </div>
              <div v-if="resourcecomparetype === 'dna'">
                <el-form-item label="哈希值" prop="dna">
                  <el-input class="form-item-content" v-model="rightResultForm.dna" :readonly="true"></el-input>
                </el-form-item>
              </div>
              <el-form-item label="消息内容" prop="content">
                <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="rightResultForm.content" :readonly="true"></el-input>
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
        <el-form-item label="URL" prop="url" v-if="textcomparetype === 'user'">
          <el-input class="form-item-content" v-model="resultForm.url" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="哈希值" prop="dna" v-if="textcomparetype === 'dna'">
          <el-input class="form-item-content" v-model="resultForm.dna" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="相似度" prop="similarity">
          <el-input class="form-item-content" v-model="resultForm.similarity" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="对比信息" prop="compareinfo">
          <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="resultForm.compareinfo" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="原文信息" prop="content">
          <el-input :autosize="{ minRows: 6 }" class="form-item-content" type="textarea" v-model="resultForm.content" :readonly="true"></el-input>
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
import { STORAGE_KEY, COMPANYS } from '@/utils/constants'
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
      resloading: false,
      textloading: false,
      companys: COMPANYS,
      comparetypes: [
        { label: 'URL查询', value: 'user' },
        { label: '哈希值查询', value: 'dna' }
      ],
      textCompareForm: {
        dna: '',
        url: '',
        uid: '',
        mid: '',
        company: '',
        compareInfo: ''
      },
      textCompareFormRules: {
        dna: [{ required: true, message: '请输入哈希值', trigger: 'blur' }],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ],
        compareInfo: [
          { required: true, message: '请输入对比内容', trigger: 'blur' }
        ]
      },
      resourcecomparetype: 'user',
      textcomparetype: 'user',
      leftCompareForm: {
        dna: '',
        url: '',
        uid: '',
        mid: '',
        company: ''
      },
      leftCompareFormRules: {
        dna: [{ required: true, message: '请输入哈希值', trigger: 'blur' }],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ]
      },
      rightCompareForm: {
        dna: '',
        url: '',
        uid: '',
        mid: '',
        company: ''
      },
      rightCompareFormRules: {
        dna: [{ required: true, message: '请输入哈希值', trigger: 'blur' }],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ]
      },
      twoResultDialogVisible: false,
      leftResultForm: {
        company: '',
        dna: '',
        url: '',
        uid: '',
        mid: '',
        content: ''
      },
      rightResultForm: {
        company: '',
        dna: '',
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
      this.textCompareForm.company = this.companys[0].value
      this.leftCompareForm.company = this.companys[0].value
      this.rightCompareForm.company = this.companys[0].value
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
                compareType: this.resourcecomparetype
              }

              switch (params.compareType) {
                case 'user':
                  params.src_uid = leftFormData.uid
                  params.src_mid = leftFormData.mid
                  params.src_company = leftFormData.company
                  params.dst_uid = rightFormData.uid
                  params.dst_mid = rightFormData.mid
                  params.dst_company = rightFormData.company
                  break
                case 'dna':
                  params.src_dna = leftFormData.dna
                  params.dst_dna = rightFormData.dna
                  break
              }
              console.log(params)

              this.resloading = true
              API.compare(params).then(
                res => {
                  if (
                    res.data.msg.toUpperCase() === 'OK' &&
                    res.data.data !== undefined
                  ) {
                    let rdd = res.data.data
                    // similarity
                    this.towResultSimilarity = tool.checkData(rdd.similarity)
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
                    this.leftResultForm.dna = tool.checkData(
                      this.leftCompareForm.dna
                    )
                    this.leftResultForm.content = tool.checkData(rdd.src)
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
                    this.rightResultForm.dna = tool.checkData(
                      this.rightCompareForm.dna
                    )
                    this.rightResultForm.content = tool.checkData(
                      rdd.src
                    )
                    // show
                    this.twoResultDialogVisible = true
                  } else {
                    this.$error(res.data.msg)
                  }
                  this.resloading = false
                },
                _ => {
                  this.resloading = false
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
            compareType: this.textcomparetype
          }

          switch (this.textcomparetype) {
            case 'user':
              params.compareType = 'user'
              params.src_uid = formData.uid
              params.src_mid = formData.mid
              params.src_company = formData.company
              params.dst_content = formData.compareInfo
              break
            case 'dna':
              params.compareType = 'dna'
              params.src_dna = formData.dna
              params.dst_content = formData.compareInfo
              break
          }
          console.log(params)

          this.textloading = true
          API.compareText(params).then(
            res => {
              console.log(res)
              if (
                res.data.msg.toUpperCase() === 'OK' &&
                res.data.data !== undefined
              ) {
                let rdd = res.data.data
                this.resultForm.similarity = tool.checkData(rdd.similarity)
                this.resultForm.url = tool.checkData(formData.url)
                this.resultForm.dna = tool.checkData(formData.dna)
                this.resultForm.compareinfo = tool.checkData(
                  formData.compareInfo
                )
                this.resultForm.content = tool.checkData(rdd.src)
                this.resultDialogVisible = true
              } else {
                this.$error(res.data.msg)
              }
              this.textloading = false
            },
            _ => {
              this.textloading = false
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
</style>
