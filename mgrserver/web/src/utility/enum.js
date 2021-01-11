import vue from "vue";
import $ from 'jquery';

window.EnumList = {};
export function EnumUtility() { }

EnumUtility.prototype.Set = function (data, type, isAddition) {
  if (!Array.isArray(data) || data.length == 0) {
    return false
  }
  if (isAddition) {
    return window.EnumList["enums_" + type] = JSON.stringify(data);
  }
  var __enumData = {};
  data.forEach(function (item) {
    if (!__enumData[item.type]) {
      __enumData[item.type] = [];
    }
    __enumData[item.type].push(item);
  });
  for (var key in __enumData) {
    window.EnumList["enums_" + key] = JSON.stringify(__enumData[key]);
  }
  return
};

EnumUtility.prototype.Get = function (type, param, url) {
  if (!type) return [];
  if (!window.EnumList["enums_" + type]) {
    var postData = { dic_type: type }
    if (JSON.stringify(param) != JSON.stringify({})) {
      for (var key in param) {
        if (param.hasOwnProperty(key) === true) {
          postData[key] = param[key];
        }
      }
    }
    var isAddition = false
    var postUrl = "/dds/dictionary/get"
    if (url != "") {
      postUrl = url
      isAddition = true
    }
    $.ajax({
      url: process.env.VUE_APP_API_URL + postUrl,
      data: postData,
      async: false,
      type: "GET",
      success: function (data) {
        EnumUtility.prototype.Set(data, type, isAddition)
      }
    })
  }
  return window.EnumList["enums_" + type] ? JSON.parse(window.EnumList["enums_" + type]) : [];
};

export const EnumFilter = vue.filter('EnumFilter', (value, enumType) => {
  if (value !== '') {
    let enumUtility = new EnumUtility()
    let enumMap = enumUtility.Get(enumType)
    let result = value
    if (enumMap != undefined && enumMap.length > 0) {
      enumMap.forEach(item => {
        if (item.value === value) {
          result = item.name
        }
      })
      return result
    }
  } else {
    return '-'
  }
})
