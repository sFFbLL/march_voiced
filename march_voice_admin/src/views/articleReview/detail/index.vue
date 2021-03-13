<template>
  <!-- eladmin自带的布局方式 -->
  <div v-loading="loading" class="app-container">
    <!-- 头部盒子 -->
    <!-- <el-row> -->
    <!-- <el-col :span="10"> -->
    <div class="style-a4">
      <div class="position-box">
        <div class="body-box">
          <div class="header-but">
            <!-- 状态显示 -->
            <div
              v-if="this.$route.query.status == 1"
              class="status-box success"
            >
              <span>已通过</span>
            </div>
            <div v-if="this.$route.query.status == 2" class="status-box wait">
              <span>待审核</span>
            </div>
            <div v-if="this.$route.query.status == 3" class="status-box wait">
              <span>已驳回</span>
            </div>
            <div class="status-box reject">
              <Reject v-if="this.showReject" :content="content" :data="data" />
            </div>
          </div>
          <div v-if="nois" class="ee">
            <el-alert
              title="该文章含有敏感词汇，请自行选择是否推荐！"
              type="warning"
              :closable="false"
              show-icon
            />
          </div>
          <div class="tittttt">
            <div class="tittle">
              <span>{{ title }}</span>
            </div>
            <!-- 获奖成果标题 -->
            <span class="ql-editor" v-html="showKeyWord(contents)" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Reject from '@/components/Affirmation/DetailAffirmation.vue'
import { getInfoById, sensitive } from '@/api/review/articleReview.js'

export default {
  name: 'Detail',
  components: { Reject },
  data() {
    return {
      nois: false,
      recommend: null,
      content: {
        status: this.$route.query.status,
        recommend: this.$route.query.recommend,
        id: this.$route.query.id
      },
      title: null,
      data: {
        status: null
      },
      loading: false,
      showReject: true,
      contents: null,
      Sensitive1: [],
      arrnew: []
    }
  },
  created() {
    getInfoById('/api/article/detail', this.$route.query.id).then(
      data => {
        this.title = data.data.title
        this.contents = data.data.content
      }
    )
    sensitive('/api/article/word', this.$route.query.id).then(
      data => {
        if (data.data.Sensitive != null) {
          this.arrnew = data.data.Sensitive.map((obj, index) => {
            return obj.word
          }).join(',').split(',')
          if (this.arrnew && this.arrnew != null) {
            this.nois = true
          }
        }
      }
    )
  },

  methods: {
    showStatus: function(data) {
      console.log(data)
      this.$route.query.status = data
      console.log(this.$route.query.status)
    },
    showKeyWord(val) {
      if (val) {
        val = val + '' // 转化为字符串类型
      } else {
        return
      }
      // 1.传入关键词数组keyWordArr
      const keyWordArr = this.arrnew
      let str = val

      if (keyWordArr && keyWordArr != '') {
        // 2.定制关键词对应正则
        keyWordArr.forEach(e => {
          const regStr = '' + `(${e})`
          const reg = new RegExp(regStr, 'g')
          // 3.正则替换，关键词飘红
          str = str.replace(reg, '<span style="color:red;">' + e + '</span>')
        })
      }
      return str
    },
    sensitive,
    getInfoById,
    showHtml(contents) {
      return contents
        .replace(contents ? /&(?!#?\w+;)/g : /&/g, '&amp;')
        .replace(/&lt;/g, '<')
        .replace(/&gt;/g, '>')
        .replace(/&quot;/g, '"')
        .replace(/&#39;/g, "'")
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@import "@/assets/styles/detail.scss";
</style>
