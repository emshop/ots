import vue from "vue";

//所有保存的枚举数据
window._EnumList_ = {};
//根据type保存的回调函数
window._EnumCallbackFunc_ = {
  "*" : function(tp){return [];}
}

/*
* 枚举对象使用时须通过引用并进行初始化
* import enums from './enums'
* Vue.use(enums);
*/
export function Enum() {}


//当参数为function时回调 回调函数须为同步请求
/*
*Vue.prototype.$enum.callback(function(type){
*  return = Vue.prototype.$http.xpost("/dds/dictionary/get", { dic_type: "product_status" }, "", true) 
*})
*/
Enum.prototype.callback = function (callback, tp) {
  if (typeof callback == "function"){
    var type = tp || "*"
    window._EnumCallbackFunc_[type] = callback
    return
  }
  throw new Error("无效参数，类型:function(type){return [];}返回一个数组");
};

//数据保存
Enum.prototype.set = function (data, type) {
  if (typeof data == "function"){
    if (!type){
        return
    }
    window._EnumCallbackFunc_[type] = data
    return
  }
  
  if (!Array.isArray(data) || data.length == 0) { 
    return
  }

  if(!checkData(data)){
    throw new Error("数据格式错误，格式:[{type:'zxc',name:'上架',value:'1', pid: 1002},...]其中type与pid参数为非必要参数，name与value为必要参数");
  }

  data.forEach(function (item) {
    var tp = item.type || type
    if(tp){
      if (!window._EnumList_[tp]) {
        window._EnumList_[tp] = [];
      }
      window._EnumList_[tp].push(item);
    }
  });
  return
};

//数据获取
Enum.prototype.get = function (type, pid, group) {
  if (!type) return [];

  var result = getEnumList(type, pid, group) //根据存储查询
  if (needCallback(result)){
    result = getEnumListByCallback(type, pid, group)//根据回调获取
  }
  return result
};

//根据value值获取name
Enum.prototype.getName = function (type, value, group) {
  if (!value) {
    return "-"
  }

  var enumMap = Enum.prototype.get(type, null, group)
  var data = value.split(",")
  var result = []

  for (var i = 0; i < data.length; i++){
    for (var j = 0; j < enumMap.length; j++){
      if (enumMap[j].value == data[i]) {
        result.push(enumMap[j].name)
        break
      }
    }
  }
  return result.join(",") || value 
}

//对应type数据刷新
Enum.prototype.clear = function (type) {
  window._EnumList_[type] = null;
};

//filter
vue.filter('fltrEnum', (value, enumType) => {
  return Enum.prototype.getName(enumType, value, "fltr")
})

//数据格式检查
function checkData(data){
  for (var i = 0; i < data.length; i++){
    if(!Object.prototype.hasOwnProperty.call(data[i], "name") || !Object.prototype.hasOwnProperty.call(data[i], "value")){
      return false
    }
  }
  return true
}

//根据回调函数获取数据
function getEnumListByCallback(type, pid, group){
  var handle = window._EnumCallbackFunc_[type] || window._EnumCallbackFunc_["*"]
  
  var data = handle.apply(this, [type])
  Enum.prototype.set(data, type)
  return getEnumList(type, pid, group)
}

//根据type从window._EnumList_中获取相应的数据
function getEnumList(type, pid, group){
  var list = []
  var location = window.location.pathname
  var result = window._EnumList_[type] || []
  result.forEach((item)=>{
    if ((!item.group || item.group == group) && (!pid ||item.pid == pid) && (!item.location || item.location == location || item.location.indexOf(location) > -1)) {
      list.push(item)
    }
  }) 
  return list
}

//检查是否需要执行回调
function needCallback(list){
  if(!list || list.length == 0){
     return true
  }
  if (list.length == 1){
      return list[0].value == "*" || list[0].value == "0"
  }
  return false
}
