<template>
    <div class="login_container">
        <div class="title_container">
            <h3 class="title">药 品 保 障 系 统</h3>
        </div>
        <div class="login_box">
            <el-form ref="loginForm" label-width="0px" class="login_form">
                <el-form-item >
                    <el-input v-model="accountid" prefix-icon="el-icon-user-solid" placeholder="请输入账号" ></el-input>
                </el-form-item>
                <el-form-item >
                    <el-input prefix-icon="el-icon-lock" placeholder="请输入密码" v-model="passwd" show-password></el-input>
                </el-form-item>
                <el-form-item class="btns">
                    <el-button :loading="loading" type="primary" @click="handleLogin()">登录</el-button>
                    <el-button type="info" @click="resetForm()">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
import { queryAccountList } from '@/api/account'
import {queryAccess} from '@/api/account'
import md5 from 'js-md5'

export default {
  name: 'Login',
  data() {
    return {
      loading: false,
      redirect: undefined,
      accountList: [],
      value: '',
      accountid:'',
      passwd:''
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        this.accountList = response
      }
    })
  },
  methods: {
    resetForm(){
      this.accountid ='',
      this.passwd ='',
      this.value =''
    },
    handleLogin() {
      this.value=""
      let passwd = md5(this.passwd)
      queryAccess({
        accountID:this.accountid,
        passWd:passwd
      }).then(response =>{
        if (response.Access == true) {
            this.value = this.accountid
        }
        if (this.value) {
          this.loading = true
          this.$store.dispatch('account/login', this.value).then(() => {
            this.$router.push({ path: this.redirect || '/' })
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else {
          this.$message('帐号或密码错误')
        }
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.title_container{
    position: relative;

    .title{
        font-size: 40px;
        margin: 0px auto 50px auto;
        padding: 40px 5px 10px 15px;
        text-align: center;
        color: #fff;
        font-weight: bold;
    }
}

.login_container{
    background-color: #2b4b6b;
    height: 100%;
}

.login_box{
    width: 450px;
    height: 300px;
    background: #fff;
    border-radius: 5px;
    position: absolute;
    left: 50%;
    top: 60%;
    transform: translate(-50%,-50%);
}

.avatar_box{
    height: 130px;
    width: 130px;
    border: 1px solid #eee;
    border-radius: 50%;
    padding: 10px;
    box-shadow: 0 0 10px;
    position: absolute;
    left: 50%;
    transform: translate(-50%,-50%);
    background-color: #fff;
    img{
        height: 100%;
        width: 100%;
        border-radius: 50%;
        background-color: #eee;
    }
}

.btns{
    display: flex;
    justify-content: flex-end;
}

.login_form{
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
}
</style>
