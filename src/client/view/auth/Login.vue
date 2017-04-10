<template>
  <div class="login-wrapper" v-show="!loggedIn">
    <div class="bg"></div>
    <h1>{{$t('title')}}</h1>
    <el-form ref="form" :model="form" :rules="rules"
      @submit.native.prevent="onSubmit">
      <el-form-item>
        <el-select :value="globalConfig.lang" @input="changeLang(arguments[0])" style="width:100%">
          <el-option v-for="lang in globalConfig.langs" :key="lang.value"
            :label="lang.label" :value="lang.value"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="username">
        <el-input v-model="form.username" :placeholder="$t('login.username')"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input v-model="form.password" type="password" :placeholder="$t('login.password')"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button class="login-button" :class="{error: loginError}" type="primary"
          native-type="submit" :loading="loading">{{isInit?$t('initialize'):$t('login.button')}}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { mapGetters, mapActions } from 'vuex'
import locales from 'locales/login'
export default {
  locales,
  data () {
    return {
      form: {
        username: '',
        password: ''
      },
      rules: {
        username: [{
          required: true, message: this.$t('login.username'), trigger: 'blur'
        }, {
          min: 3, message: this.$t('login.msg.usernameMinlength', { min: 3 }), trigger: 'blur'
        }],
        password: [{
          required: true, message: this.$t('login.password'), trigger: 'blur'
        }, {
          min: 3, message: this.$t('login.msg.passwordMinlength', { min: 3 }), trigger: 'blur'
        }]
      },
      loading: false,
      loginError: false
    }
  },
  computed: {
    ...mapGetters(['loggedIn', 'globalConfig']),
    isInit () {
      return !!this.$route.meta.initialize
    }
  },
  methods: {
    ...mapActions(['login', 'changeLang', 'initialize', 'setUserInfo']),
    onSubmit () {
      this.$refs.form.validate(valid => {
        if (valid) {
          this.loading = true
          if (this.isInit) {
            // initialize app
            this.initialize(this.form).then(data => {
              this.setUserInfo({
                id: '1',
                role: 'admin',
                username: this.form.username,
                access_token: data.token // eslint-disable-line
              }).then(() => {
                this.$notify(this.$t('tips.initializeCompleted'))
                this.$router.push('/')
              })
            }).catch(res => {
              this.$notify(res.message || this.$t('tips.initializeFailed'))
            })
          } else {
            this.login(this.form).then((data) => {
              this.loading = false
              this.$router.push(this.$route.query.redirect || '/')
            }).catch((err) => {
              this.$notify({
                title: this.$t('message.error'),
                message: err.message || this.$t('login.authFail'),
                type: 'error',
                duration: 1500
              })
              this.loading = false
              this.loginError = true
              setTimeout(() => {
                this.loginError = false
              }, 500)
            })
          }
        }
      })
    }
  }
}
</script>
<style lang="stylus">
@import "../../assets/css/variable"
$input-width = 15rem
.login-wrapper
  align-self center
  .bg
    position absolute
    left 0
    right 0
    top 0
    bottom 0
    width 100%
    height 100%
    background-size cover
    background-position 100%
    background-image url('../../assets/images/login-bg.jpg')
  > h1
    position relative
    margin 0 0 1rem
    text-align center
    z-index 1
  > form
    width $input-width
    margin 0 auto
    .el-input__inner
      color $color-black-exact-light
      border-color $color-silver-light
      background-color transparent
      &:focus
        color $color-black
        border-color $color-black
  .login-button
    width 100%
    &.error
      animation shake .5s
</style>
