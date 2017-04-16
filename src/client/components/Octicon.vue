<template>
  <svg version="1.1" :class="clazz" :role="label ? 'img' : 'presentation'" :aria-label="label" :x="x" :y="y" :width="width" :height="height" :view-box.camel="box" :style="style">
    <slot><path v-if="icon" :d="icon.d" /></slot>
  </svg>
</template>

<style>
.octicon {
  display: inline-block;
  fill: currentColor;
}

.octicon > path {
  transform-origin: 50% 50%;
}

.octicon-flip-horizontal > path {
  transform: scale(-1, 1);
}

.octicon-flip-vertical > path {
  transform: scale(1, -1);
}

.octicon-spin > path {
  animation: octicon-spin 1s 0s infinite linear;
}

.octicon-inverse {
  color: #fff;
}

@keyframes octicon-spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>

<script>
const icons = {}

export default {
  name: 'octicon',
  props: {
    name: {
      type: String,
      required: true
    },
    scale: [Number, String],
    spin: Boolean,
    inverse: Boolean,
    flip: {
      validator (val) {
        return val === 'horizontal' || val === 'vertical'
      }
    },
    label: String
  },
  data () {
    return {
      x: false,
      y: false,
      childrenWidth: 0,
      childrenHeight: 0,
      outerScale: 1
    }
  },
  computed: {
    normalizedScale () {
      let scale = this.scale
      scale = typeof scale === 'undefined' ? 1 : Number(scale)
      if (isNaN(scale) || scale <= 0) {
        console.warn(`Invalid prop: prop "scale" should be a number over 0.`, this)
        return this.outerScale
      }
      return scale * this.outerScale
    },
    clazz () {
      return {
        'octicon': true,
        'octicon-spin': this.spin,
        'octicon-flip-horizontal': this.flip === 'horizontal',
        'octicon-flip-vertical': this.flip === 'vertical',
        'octicon-inverse': this.inverse
      }
    },
    icon () {
      if (this.name) {
        return icons[this.name]
      }
      return null
    },
    box () {
      if (this.icon) {
        if (this.icon.viewBox) {
          return this.icon.viewBox
        }
        return `0 0 ${this.icon.width} ${this.icon.height}`
      }
      return `0 0 ${this.width} ${this.height}`
    },
    width () {
      return this.childrenWidth || this.icon && this.icon.width * this.normalizedScale || 0
    },
    height () {
      return this.childrenHeight || this.icon && this.icon.height * this.normalizedScale || 0
    },
    style () {
      if (this.normalizedScale === 1) {
        return false
      }
      return {
        fontSize: this.normalizedScale + 'em'
      }
    }
  },
  mounted () {
    if (this.icon) {
      return
    }
    this.$children.forEach(child => {
      child.outerScale = this.normalizedScale
    })
    let width = 0
    let height = 0
    this.$children.forEach(child => {
      width = Math.max(width, child.width)
      height = Math.max(height, child.height)
    })
    this.childrenWidth = width
    this.childrenHeight = height
    this.$children.forEach(child => {
      child.x = (width - child.width) / 2
      child.y = (height - child.height) / 2
    })
  },
  register (data) {
    for (const name in data) {
      icons[name] = data[name]
    }
  },
  icons
}
</script>
