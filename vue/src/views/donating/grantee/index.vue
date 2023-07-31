<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ accountId }}</p>
      <p>用户名: {{ userName }}</p>
      <p>余额: ￥{{ balance }} 元</p>
    </el-alert>
    <div v-if="donatingList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in donatingList" :key="index" :span="6" :offset="1">
        <el-card class="d-all-card">
          <div slot="header" class="clearfix">
            <span>{{ val.donatingStatus }}</span>
            <el-button v-if="roles[0] !== 'admin'&&val.grantee===accountId&&val.donatingStatus==='捐赠中'" style="float: right; padding: 3px 6px" type="text" @click="updateDonating(val,'done')">确认接收</el-button>
            <el-button v-if="roles[0] !== 'admin'&&(val.donor===accountId||val.grantee===accountId)&&val.donatingStatus==='捐赠中'" style="float: right; padding: 3px 0" type="text" @click="updateDonating(val,'cancelled')">取消</el-button>
          </div>
          <div class="item">
            <el-tag>房产ID: </el-tag>
            <span>{{ val.objectOfDonating }}</span>
          </div>
          <div class="item">
            <el-tag type="success">捐赠者ID: </el-tag>
            <span>{{ val.donor }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">受赠人ID: </el-tag>
            <span>{{ val.grantee }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">创建时间: </el-tag>
            <span>{{ val.createTime }}</span>
          </div>
          <el-popover placement="right" width="500" trigger="click">
            <!--<baidu-map :center="location" :zoom="zoom" @ready="getRoute" :scroll-wheel-zoom="true"> 
              <bm-view style="width: 270px; height:200px; flex: 1"></bm-view>
              <bm-local-search :keyword="addressKeyword" :auto-viewport="true" style="display: none"></bm-local-search>
            </baidu-map>-->
            <baidu-map class="map" :scroll-wheel-zoom="true" :center="center" :zoom="zoom" @ready="ready">
              <new-polyline
               v-if="points && points.length >= 2 && checkPoints({ points })"
               :path="points"
               :icons="icons"
               stroke-color="#0091ea"
               :stroke-opacity="0.8"
               :stroke-weight="10"
              ></new-polyline>
            </baidu-map>
            <h4>药品ID{{ val.objectOfDonating }}</h4>
              <el-timeline>
                <el-timeline-item
                  v-for="(activity, index) in activities"
                  :key="index"
                  :icon="activity.icon"
                  :type="activity.type"
                  :color="activity.color"
                  :size="activity.size"
                  :timestamp="activity.timestamp">
                  {{activity.content}}
                </el-timeline-item>
              </el-timeline>
            <el-button slot="reference">点击查看物流信息</el-button>
          </el-popover>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import newPolyline from "@/components/common/new_polyline"
import { mapGetters } from 'vuex'
import { queryDonatingList, updateDonating } from '@/api/donating'

export default {
  components: {
      newPolyline
    },
  name: 'AllDonating',
  data() {
    return {
      /*location: {
        lng: 116.404,
        lat: 39.915
      },
      zoom: 12.8,*/
      center: {
        lng: 116.404,
        lat: 39.915
      },
      zoom: 5,
      points: [],
      icons: [],
      loading: true,
      donatingList: [],
      logistics:false,
      activities: [{
          content: '已发货:南京市',
          timestamp: '2020-11-10 08：30',
          size: 'large',
          type: 'primary',
          icon: 'el-icon-more'
        }, {
          content: '运输中：到达合肥市',
          timestamp: '2020-11-12 12：00',
          color: '#0bbd87',
          size: 'large',
          icon:'el-icon-s-promotion'
        }, {
          content: '运输中：到达武汉市',
          timestamp: '2020-11-14 20：00',
          color: '#409EFF',
          size: 'large',
          icon:'el-icon-s-promotion'
        }, {
          content: '已签收 华中科技大学同济医学院附属同济医院',
          timestamp: '2020-11-15 10：00',
          color: '#fb7293',
          size: 'large',
          icon:'el-icon-s-flag'
        }],
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ])
  },
  created() {
    queryDonatingList().then(response => {
      if (response !== null) {
        this.donatingList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
  },
  methods: {
    ready({ BMap, map }) {
      this.points = this.getPointsSomehow();
      var sy = new BMap.Symbol(BMap_Symbol_SHAPE_FORWARD_OPEN_ARROW, {
        scale: 0.5, // 图标缩放大小
        strokeColor: "#fff", // 设置矢量图标的线填充颜色
        strokeWeight: "3" // 设置线宽
      });
      var icons = new BMap.IconSequence(sy, "5%", "15%");
      this.icons.push(icons)
      console.log(this.points)
      map.addOverlay(this.points)
    },
    getPointsSomehow() {
      // 116.324356,39.980648
      // 118.532031,32.010158
      // 121.475599,31.380939
      var arr = [
        { lng: 118.780133, lat: 32.069235 },
        { lng: 117.291117, lat: 31.798837 },
        { lng: 114.316551, lat: 30.528813 },
        { lng: 114.274665, lat: 30.584394 }
      ];
      return arr;
    },
    // 查重 如果数组中只有一组有意义的坐标,绘制折线时可能会报错,因为绘制一条折线需要两组不同的坐标点
    checkPoints({ points }) {
      // 拿到第一组点
      var originPoint = points[0];
      // 将第一组点与后续点进行对比 如果坐标集合中只有一组实际坐标点则return false
      // 只要坐标集合中有两组不同的实际坐标点,则return true
      for (var i = 1; i < points.length; i++) {
        if (
          points[i].lat !== originPoint.lat ||
          points[i].lng !== originPoint.lng
        ) {
          return true;
        }
      }
      return false;
    },
    /*
    showlogistics(){
      this.logistics =!this.logistics
      console.logistics
    },
    getRoute({BMap, map}) {
      map.centerAndZoom(new BMap.Point(118.780133, 32.069235), 15)
      let myP1 = new BMap.Point(118.780133, 32.069235)
      let myP2 = new BMap.Point(117.291117, 31.798837)
      let driving = new BMap.DrivingRoute(map, {renderOptions: {map: map, autoViewport: true}})
      driving.search(myP1, myP2)
    },
    getLocationPoint(e) {
      this.lng = e.point.lng;
      this.lat = e.point.lat;
      //创建地址解析器的实例 
      let geoCoder = new BMap.Geocoder();
      //获取位置对应的坐标
      geoCoder.getPoint(this.addressKeyword, point => {this.selectedLng = point.lng;this.selectedLat = point.lat;});
      //利用坐标获取地址的详细信息
      geocoder.getLocation(e.point, res=> {console.log(res)})
    },*/
    
    updateDonating(item, type) {
      let tip = ''
      if (type === 'done') {
        tip = '确认接受捐赠'
      } else {
        tip = '取消捐赠操作'
      }
      this.$confirm('是否要' + tip + '?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).then(() => {
        this.loading = true
        updateDonating({
          donor: item.donor,
          grantee: item.grantee,
          objectOfDonating: item.objectOfDonating,
          status: type
        }).then(response => {
          this.loading = false
          if (response !== null) {
            this.$message({
              type: 'success',
              message: tip + '操作成功!'
            })
          } else {
            this.$message({
              type: 'error',
              message: tip + '操作失败!'
            })
          }
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        }).catch(_ => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
        this.$message({
          type: 'info',
          message: '已取消' + tip
        })
      })
    }
  }
}

</script>

<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }


  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .d-all-card {
    width: 280px;
    height: 320px;
    margin: 18px;
  }

  .map {
    width: 100%;
    height: 300px;}

</style>
