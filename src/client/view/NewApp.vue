<template>
  <div class="m-create-app flex flex-1 flex-main-center flex-cross-center">
    <el-form ref="form" :rules="rules" :model="form" @submit.native.prevent="onSubmit">
      <el-form-item prop="name">
        <el-input v-model="form.name" :placeholder="$t('app.create.name')"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" native-type="submit" :loading="pending" style="width:100%">{{$t('app.create.btn')}}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { app as res } from 'resources'
import locales from 'locales/apps'
export default {
  locales,
  data () {
    return {
      pending: false,
      form: {
        name: ''
      },
      rules: {
        name: [
          { required: true, message: this.$t('app.create.name'), trigger: 'blur' },
          { min: 3, message: this.$t('app.create.minlength'), trigger: 'blur' },
          { validator (rule, val, cb) {
            if (/[a-zA-Z-_\d]+/.test(val)) {
              cb()
            } else {
              cb(rule.message)
            }
          }, message: this.$t('app.create.character'), trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    onSubmit () {
      this.$refs.form.validate(valid => {
        if (valid) {
          this.loading = true
          res.save(null, this.form).then(() => {
            this.$router.push(`/apps/${this.form.name}/deploy`)
          }).catch(() => {}).then(() => {
            this.loading = false
          })
        }
      })
    }
  }
}
</script>
<style lang="stylus">
.m-create-app
  .el-form
    max-width 16rem
    width 100%
    margin 0 auto
</style>
