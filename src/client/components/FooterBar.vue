<template>
  <div class="ui-footer-bar flex flex-cross-center flex-between">
    <div class="left">
      <a href="http://dokku.viewdocs.io/dokku/" target="_blank">Dokku</a>
      <a href="https://github.com/erguotou520/dokku-web" target="_blank">Dokku Web</a>
      <a class="btn" href="https://github.com/erguotou520/dokku-web/issues" target="_blank">Support</a>
    </div>
    <div class="right">
      <a href="" class="btn">Donate</a>
      <a href="https://github.com/erguotou520/dokku-web/blob/master/LICENSE" target="_blank">License</a>
    </div>
  </div>
</template>
<style lang="stylus">
@import "../assets/css/variable.styl"
.ui-footer-bar
  height 3.125rem
  padding 0 1rem
  background-color #fff
  border-top 1px solid $color-gray-light
  .left
  .right
    > *
      color $color-silver
      margin 0 .5rem
      &:first-child
        margin-left 0
      &:last-child
        margin-right 0
    .btn
      padding 4px 8px
      background linear-gradient(to right bottom, lighten($color-primary, 10%), $color-primary)
      color #fff
      border-radius 4px
</style>
