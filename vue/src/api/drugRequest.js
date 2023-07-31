import request from '@/utils/request'

// 新建需求
export function createDrugRequest(data) {
  return request({
    url: '/createDrugRequest',
    method: 'post',
    data
  })
}

// 获取需求信息(空json{}可以查询所有，指定proprietor可以查询指定用户名下的需求信息)
export function queryDrugRequestList(data) {
  return request({
    url: '/queryDrugRequestList',
    method: 'post',
    data
  })
}
