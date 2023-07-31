<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="用户" prop="proprietor">
        <el-select v-model="ruleForm.proprietor" placeholder="请选择用户" @change="selectGet">
          <el-option
            :key="accountId"
            :label="userName"
            :value="accountId"
          >
            <span style="float: left; color: #8492a6; font-size: 13px">{{ userName }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ accountId }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <!-- <el-form-item label="总空间 ㎡" prop="totalArea">
        <el-input-number v-model="ruleForm.totalArea" :precision="2" :step="0.1" :min="0" />
      </el-form-item> -->

      <el-form-item label="药品名" prop="drugName">
        <el-input v-model="ruleForm.drugName" placeholder="请输入药品名" style="width:200px;"></el-input>
      </el-form-item>

      <el-form-item label="商品名" prop="tradeName">
        <el-input v-model="ruleForm.tradeName" placeholder="请输入商品名" style="width:200px;"></el-input>
      </el-form-item>
      
      <el-form-item label="数量" prop="drugNumber">
        <el-input-number v-model="ruleForm.drugNumber" :precision="2" :step="1" :min="0" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即发布</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createDrugRequest } from '@/api/drugRequest'

export default {
  name: 'AddDrugRequest',
  data() {
    var checkNumber = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        proprietor: '',
        drugName: '',
        tradeName:'',
        drugNumber: 0
      },
      accountList: [],
      rules: {
        proprietor: [
          { required: true, message: '请选择用户', trigger: 'change' }
        ],
        drugName: [
          { required: true, message: '请输入药品名', trigger: 'change' }
        ],
        tradeName: [
          { required: true, message: '请输入商品名', trigger: 'change' }
        ],
        drugNumber: [
          { validator: checkNumber, trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'userName'
    ])
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        // 过滤掉管理员
        this.accountList = response.filter(item =>
          item.userName !== '管理员'
        )
      }
    })
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即新增?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loading = true
            createDrugRequest({
              accountId: this.accountId,
              proprietor: this.ruleForm.proprietor,
              drugName: this.ruleForm.drugName,
              tradeName: this.ruleForm.tradeName,
              drugNumber: this.ruleForm.drugNumber
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '新增成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '新增失败!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: '已取消新增'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.ruleForm.proprietor = accountId
    }
  }
}
</script>

<style scoped>
</style>
